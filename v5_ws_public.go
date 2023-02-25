package bybit

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// V5WebsocketPublicServiceI :
type V5WebsocketPublicServiceI interface {
	Start(context.Context, ErrHandler) error
	Run() error
	Ping() error
	Close() error

	SubscribeOrderBook(
		V5WebsocketPublicOrderBookParamKey,
		func(V5WebsocketPublicOrderBookResponse) error,
	) (func() error, error)
}

// V5WebsocketPublicService :
type V5WebsocketPublicService struct {
	client     *WebSocketClient
	connection *websocket.Conn

	paramOrderBookMap map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error
}

const (
	// V5WebsocketPublicPath :
	V5WebsocketPublicPath = "/v5/public"
)

// V5WebsocketPublicPathFor :
func V5WebsocketPublicPathFor(category CategoryV5) string {
	return V5WebsocketPublicPath + "/" + string(category)
}

// V5WebsocketPublicTopic :
type V5WebsocketPublicTopic string

const (
	// V5WebsocketPublicTopicOrderBook :
	V5WebsocketPublicTopicOrderBook = "orderbook"
)

// judgeTopic :
func (s *V5WebsocketPublicService) judgeTopic(respBody []byte) (V5WebsocketPublicTopic, error) {
	parsedData := map[string]interface{}{}
	if err := json.Unmarshal(respBody, &parsedData); err != nil {
		return "", err
	}
	if topic, ok := parsedData["topic"].(string); ok {
		switch {
		case strings.Contains(topic, "orderbook"):
			return V5WebsocketPublicTopicOrderBook, nil
		}
	}
	return "", nil
}

// parseResponse :
func (s *V5WebsocketPublicService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// Start :
func (s *V5WebsocketPublicService) Start(ctx context.Context, errHandler ErrHandler) error {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			if err := s.Run(); err != nil {
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
func (s *V5WebsocketPublicService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeTopic(message)
	if err != nil {
		return err
	}
	switch topic {
	case V5WebsocketPublicTopicOrderBook:
		var resp V5WebsocketPublicOrderBookResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveOrderBookFunc(resp.Key())
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
func (s *V5WebsocketPublicService) Ping() error {
	if err := s.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *V5WebsocketPublicService) Close() error {
	if err := s.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return nil
}
