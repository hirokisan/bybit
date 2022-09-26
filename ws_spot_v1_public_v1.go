package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// SpotWebsocketV1PublicV1Service :
type SpotWebsocketV1PublicV1Service struct {
	connection *websocket.Conn

	paramTradeMap map[SpotWebsocketV1PublicV1TradeParamKey]func(SpotWebsocketV1PublicV1TradeResponse) error
}

const (
	// SpotWebsocketV1PublicV1Path :
	SpotWebsocketV1PublicV1Path = "/spot/quote/ws/v1"
)

// SpotWebsocketV1PublicV1Event :
type SpotWebsocketV1PublicV1Event string

const (
	// SpotWebsocketV1PublicV1EventSubscribe :
	SpotWebsocketV1PublicV1EventSubscribe = "sub"
	// SpotWebsocketV1PublicV1EventUnsubscribe :
	SpotWebsocketV1PublicV1EventUnsubscribe = "cancel"
)

// SpotWebsocketV1PublicV1Topic :
type SpotWebsocketV1PublicV1Topic string

const (
	// SpotWebsocketV1PublicV1TopicTrade :
	SpotWebsocketV1PublicV1TopicTrade = SpotWebsocketV1PublicV1Topic("trade")
)

// SpotWebsocketV1PublicV1TradeParamKey :
type SpotWebsocketV1PublicV1TradeParamKey struct {
	Symbol SymbolSpot
	Topic  SpotWebsocketV1PublicV1Topic
}

// SpotWebsocketV1PublicV1TradeResponse :
type SpotWebsocketV1PublicV1TradeResponse struct {
	Symbol         SymbolSpot                   `json:"symbol"`
	SymbolName     string                       `json:"symbolName"`
	Topic          SpotWebsocketV1PublicV1Topic `json:"topic"`
	SendTime       int                          `json:"sendTime"`
	IsFirstMessage bool                         `json:"f"`

	Params SpotWebsocketV1PublicV1TradeResponseParams `json:"params"`
	Data   []SpotWebsocketV1PublicV1TradeContent      `json:"data"`
}

// SpotWebsocketV1PublicV1TradeResponseParams :
type SpotWebsocketV1PublicV1TradeResponseParams struct {
	RealtimeInterval string `json:"realtimeInterval"`
	Binary           string `json:"binary"`
}

// SpotWebsocketV1PublicV1TradeContent :
type SpotWebsocketV1PublicV1TradeContent struct {
	TradeID        string `json:"v"`
	Timestamp      int    `json:"t"`
	Price          string `json:"p"`
	Quantity       string `json:"q"`
	IsBuySideTaker bool   `json:"m"`
}

// Key :
func (p *SpotWebsocketV1PublicV1TradeResponse) Key() SpotWebsocketV1PublicV1TradeParamKey {
	return SpotWebsocketV1PublicV1TradeParamKey{
		Symbol: p.Symbol,
		Topic:  p.Topic,
	}
}

// SpotWebsocketV1PublicV1TradeParamChild :
type SpotWebsocketV1PublicV1TradeParamChild struct {
	Binary bool `json:"binary"`
}

// SpotWebsocketV1PublicV1TradeParam :
type SpotWebsocketV1PublicV1TradeParam struct {
	Symbol SymbolSpot                             `json:"symbol"`
	Topic  SpotWebsocketV1PublicV1Topic           `json:"topic"`
	Event  SpotWebsocketV1PublicV1Event           `json:"event"`
	Params SpotWebsocketV1PublicV1TradeParamChild `json:"params"`
}

// Key :
func (p *SpotWebsocketV1PublicV1TradeParam) Key() SpotWebsocketV1PublicV1TradeParamKey {
	return SpotWebsocketV1PublicV1TradeParamKey{
		Symbol: p.Symbol,
		Topic:  p.Topic,
	}
}

// addParamTradeFunc :
func (s *SpotWebsocketV1PublicV1Service) addParamTradeFunc(param SpotWebsocketV1PublicV1TradeParamKey, f func(SpotWebsocketV1PublicV1TradeResponse) error) error {
	if _, exist := s.paramTradeMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramTradeMap[param] = f
	return nil
}

// removeParamTradeFunc :
func (s *SpotWebsocketV1PublicV1Service) removeParamTradeFunc(key SpotWebsocketV1PublicV1TradeParamKey) {
	delete(s.paramTradeMap, key)
}

// retrieveTradeFunc :
func (s *SpotWebsocketV1PublicV1Service) retrieveTradeFunc(key SpotWebsocketV1PublicV1TradeParamKey) (func(SpotWebsocketV1PublicV1TradeResponse) error, error) {
	f, exist := s.paramTradeMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}

// judgeTopic :
func (s *SpotWebsocketV1PublicV1Service) judgeTopic(respBody []byte) (SpotWebsocketV1PublicV1Topic, error) {
	result := struct {
		Topic SpotWebsocketV1PublicV1Topic `json:"topic"`
	}{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}
	return result.Topic, nil
}

// parseResponse :
func (s *SpotWebsocketV1PublicV1Service) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// SubscribeTrade :
func (s *SpotWebsocketV1PublicV1Service) SubscribeTrade(symbol SymbolSpot, f func(response SpotWebsocketV1PublicV1TradeResponse) error) (func() error, error) {
	param := SpotWebsocketV1PublicV1TradeParam{
		Symbol: symbol,
		Topic:  SpotWebsocketV1PublicV1TopicTrade,
		Event:  SpotWebsocketV1PublicV1EventSubscribe,
		Params: SpotWebsocketV1PublicV1TradeParamChild{
			Binary: false,
		},
	}
	if err := s.addParamTradeFunc(param.Key(), f); err != nil {
		return nil, err
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	if err := s.connection.WriteMessage(websocket.TextMessage, []byte(buf)); err != nil {
		return nil, err
	}

	return func() error {
		param.Event = SpotWebsocketV1PublicV1EventUnsubscribe
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.connection.WriteMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamTradeFunc(param.Key())
		return nil
	}, nil
}

// Start :
func (s *SpotWebsocketV1PublicV1Service) Start(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			if err := s.Run(); err != nil {
				if IsErrWebsocketClosed(err) {
					return
				}
				log.Println(err)
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
			return
		case <-ticker.C:
			if err := s.Ping(); err != nil {
				return
			}
		case <-ctx.Done():
			log.Println("interrupt")

			if err := s.Close(); err != nil {
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

// Run :
func (s *SpotWebsocketV1PublicV1Service) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeTopic(message)
	if err != nil {
		return err
	}
	switch topic {
	case SpotWebsocketV1PublicV1TopicTrade:
		var resp SpotWebsocketV1PublicV1TradeResponse
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
	}
	return nil
}

// Ping :
func (s *SpotWebsocketV1PublicV1Service) Ping() error {
	if err := s.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *SpotWebsocketV1PublicV1Service) Close() error {
	if err := s.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return nil
}
