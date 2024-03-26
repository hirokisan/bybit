package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/signal"
	"strings"
	"sync"
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

	SubscribeKline(
		V5WebsocketPublicKlineParamKey,
		func(V5WebsocketPublicKlineResponse) error,
	) (func() error, error)

	SubscribeTicker(
		V5WebsocketPublicTickerParamKey,
		func(V5WebsocketPublicTickerResponse) error,
	) (func() error, error)

	SubscribeTrade(
		V5WebsocketPublicTradeParamKey,
		func(V5WebsocketPublicTradeResponse) error,
	) (func() error, error)

	SubscribeLiquidation(
		V5WebsocketPublicLiquidationParamKey,
		func(V5WebsocketPublicLiquidationResponse) error,
	) (func() error, error)
}

// V5WebsocketPublicService :
type V5WebsocketPublicService struct {
	client     *WebSocketClient
	connection *websocket.Conn
	category   CategoryV5

	mu sync.Mutex

	paramOrderBookMap   map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error
	paramKlineMap       map[V5WebsocketPublicKlineParamKey]func(V5WebsocketPublicKlineResponse) error
	paramTickerMap      map[V5WebsocketPublicTickerParamKey]func(V5WebsocketPublicTickerResponse) error
	paramTradeMap       map[V5WebsocketPublicTradeParamKey]func(V5WebsocketPublicTradeResponse) error
	paramLiquidationMap map[V5WebsocketPublicLiquidationParamKey]func(V5WebsocketPublicLiquidationResponse) error
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
	V5WebsocketPublicTopicOrderBook = V5WebsocketPublicTopic("orderbook")

	// V5WebsocketPublicTopicKline :
	V5WebsocketPublicTopicKline = V5WebsocketPublicTopic("kline")

	// V5WebsocketPublicTopicTicker :
	V5WebsocketPublicTopicTicker = V5WebsocketPublicTopic("tickers")

	// V5WebsocketPublicTopicTrade :
	V5WebsocketPublicTopicTrade = V5WebsocketPublicTopic("publicTrade")

	// V5WebsocketPublicTopicLiquidation :
	V5WebsocketPublicTopicLiquidation = V5WebsocketPublicTopic("liquidation")
)

func (t V5WebsocketPublicTopic) String() string {
	return string(t)
}

// judgeTopic :
func (s *V5WebsocketPublicService) judgeTopic(respBody []byte) (V5WebsocketPublicTopic, error) {
	parsedData := map[string]interface{}{}
	if err := json.Unmarshal(respBody, &parsedData); err != nil {
		return "", err
	}
	if topic, ok := parsedData["topic"].(string); ok {
		switch {
		case strings.Contains(topic, V5WebsocketPublicTopicOrderBook.String()):
			return V5WebsocketPublicTopicOrderBook, nil
		case strings.Contains(topic, V5WebsocketPublicTopicKline.String()):
			return V5WebsocketPublicTopicKline, nil
		case strings.Contains(topic, V5WebsocketPublicTopicTicker.String()):
			return V5WebsocketPublicTopicTicker, nil
		case strings.Contains(topic, V5WebsocketPublicTopicTrade.String()):
			return V5WebsocketPublicTopicTrade, nil
		case strings.Contains(topic, V5WebsocketPublicTopicLiquidation.String()):
			return V5WebsocketPublicTopicLiquidation, nil
		}
	}
	return "", nil
}

// UnmarshalJSON :
func (r *V5WebsocketPublicTickerData) UnmarshalJSON(data []byte) error {
	switch r.category {
	case CategoryV5Linear, CategoryV5Inverse:
		return json.Unmarshal(data, &r.LinearInverse)
	case CategoryV5Option:
		return json.Unmarshal(data, &r.Option)
	case CategoryV5Spot:
		return json.Unmarshal(data, &r.Spot)
	}
	return errors.New("unsupported format")
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
			s.client.debugf("caught websocket public service interrupt signal")

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
	case V5WebsocketPublicTopicKline:
		var resp V5WebsocketPublicKlineResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveKlineFunc(resp.Key())
		if err != nil {
			return err
		}

		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPublicTopicTicker:
		var resp V5WebsocketPublicTickerResponse
		resp.Data.category = s.category
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveTickerFunc(resp.Key())
		if err != nil {
			return err
		}

		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPublicTopicTrade:
		var resp V5WebsocketPublicTradeResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveTradeFunc(resp.Key())
		if err != nil {
			return err
		}

		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPublicTopicLiquidation:
		var resp V5WebsocketPublicLiquidationResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveLiquidationFunc(resp.Key())
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
	// NOTE: It appears that two messages need to be sent.
	// REF: https://github.com/hirokisan/bybit/pull/127#issuecomment-1537479346
	if err := s.writeMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	if err := s.writeMessage(websocket.TextMessage, []byte(`{"op":"ping"}`)); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *V5WebsocketPublicService) Close() error {
	if err := s.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil && !errors.Is(err, websocket.ErrCloseSent) {
		return err
	}
	return nil
}

func (s *V5WebsocketPublicService) writeMessage(messageType int, body []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.connection.WriteMessage(messageType, body); err != nil {
		return err
	}
	return nil
}
