package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// V5WebsocketPrivateServiceI :
type V5WebsocketPrivateServiceI interface {
	Start(context.Context, ErrHandler) error
	Subscribe() error
	Run() error
	Ping() error
	Close() error

	SubscribeOrder(
		func(V5WebsocketPrivateOrderResponse) error,
	) (func() error, error)

	SubscribePosition(
		func(V5WebsocketPrivatePositionResponse) error,
	) (func() error, error)

	SubscribeWallet(
		func(V5WebsocketPrivateWalletResponse) error,
	) (func() error, error)
}

// V5WebsocketPrivateService :
type V5WebsocketPrivateService struct {
	client     *WebSocketClient
	connection *websocket.Conn

	mu sync.Mutex

	paramOrderMap    map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateOrderResponse) error
	paramPositionMap map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponse) error
	paramWalletMap   map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateWalletResponse) error
}

const (
	// V5WebsocketPrivatePath :
	V5WebsocketPrivatePath = "/v5/private"
)

// V5WebsocketPrivateTopic :
type V5WebsocketPrivateTopic string

const (
	// V5WebsocketPrivateTopicPong :
	V5WebsocketPrivateTopicPong V5WebsocketPrivateTopic = "pong"

	// V5WebsocketPrivateTopicOrder :
	V5WebsocketPrivateTopicOrder V5WebsocketPrivateTopic = "order"

	// V5WebsocketPrivateTopicPosition :
	V5WebsocketPrivateTopicPosition V5WebsocketPrivateTopic = "position"

	// V5WebsocketPrivateTopicWallet :
	V5WebsocketPrivateTopicWallet V5WebsocketPrivateTopic = "wallet"
)

// V5WebsocketPrivateParamKey :
type V5WebsocketPrivateParamKey struct {
	Topic V5WebsocketPrivateTopic
}

// judgeTopic :
func (s *V5WebsocketPrivateService) judgeTopic(respBody []byte) (V5WebsocketPrivateTopic, error) {
	parsedData := map[string]interface{}{}
	if err := json.Unmarshal(respBody, &parsedData); err != nil {
		return "", err
	}
	if retMsg, ok := parsedData["op"].(string); ok && retMsg == "pong" {
		return V5WebsocketPrivateTopicPong, nil
	}
	if topic, ok := parsedData["topic"].(string); ok {
		return V5WebsocketPrivateTopic(topic), nil
	}
	if authStatus, ok := parsedData["success"].(bool); ok {
		if !authStatus {
			return "", errors.New("auth failed: " + parsedData["ret_msg"].(string))
		}
	}
	return "", nil
}

// parseResponse :
func (s *V5WebsocketPrivateService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// Subscribe : Apply for authentication when establishing a connection.
func (s *V5WebsocketPrivateService) Subscribe() error {
	param, err := s.client.buildAuthParam()
	if err != nil {
		return err
	}
	if err := s.writeMessage(websocket.TextMessage, param); err != nil {
		return err
	}
	return nil
}

// ErrHandler :
type ErrHandler func(isWebsocketClosed bool, err error)

// Start :
func (s *V5WebsocketPrivateService) Start(ctx context.Context, errHandler ErrHandler) error {
	done := make(chan struct{})

	go func() {
		defer close(done)
		defer s.connection.Close()
		_ = s.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
		s.connection.SetPongHandler(func(string) error {
			_ = s.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
			return nil
		})

		for {
			if err := s.Run(); err != nil {
				if errHandler == nil {
					return
				}
				errHandler(IsErrWebsocketClosed(err), err)
				return
			}
		}
	}()

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	for {
		select {
		case <-done:
			return nil
		case <-ticker.C:
			if err := s.Ping(); err != nil {
				return err
			}
		case <-ctx.Done():
			log.Println("interrupt")

			if err := s.Close(); err != nil {
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

// Run :
func (s *V5WebsocketPrivateService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeTopic(message)
	if err != nil {
		return err
	}
	switch topic {
	case V5WebsocketPrivateTopicPong:
		if err := s.connection.PongHandler()("pong"); err != nil {
			return fmt.Errorf("pong: %w", err)
		}
	case V5WebsocketPrivateTopicOrder:
		var resp V5WebsocketPrivateOrderResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveOrderFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPrivateTopicPosition:
		var resp V5WebsocketPrivatePositionResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrievePositionFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPrivateTopicWallet:
		var resp V5WebsocketPrivateWalletResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveWalletFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	}

	return nil
}

// Ping :
func (s *V5WebsocketPrivateService) Ping() error {
	if err := s.writeMessage(websocket.TextMessage, []byte(`{"op":"ping"}`)); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *V5WebsocketPrivateService) Close() error {
	if err := s.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil && !errors.Is(err, websocket.ErrCloseSent) {
		return err
	}
	return nil
}

func (s *V5WebsocketPrivateService) writeMessage(messageType int, body []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.connection.WriteMessage(messageType, body); err != nil {
		return err
	}
	return nil
}
