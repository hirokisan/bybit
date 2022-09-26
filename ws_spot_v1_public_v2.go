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

// SpotWebsocketV1PublicV2Service :
type SpotWebsocketV1PublicV2Service struct {
	connection *websocket.Conn

	paramTradeMap map[SpotWebsocketV1PublicV2TradeParamKey]func(SpotWebsocketV1PublicV2TradeResponse) error
}

const (
	// SpotWebsocketV1PublicV2Path :
	SpotWebsocketV1PublicV2Path = "/spot/quote/ws/v2"
)

// SpotWebsocketV1PublicV2Event :
type SpotWebsocketV1PublicV2Event string

const (
	// SpotWebsocketV1PublicV2EventSubscribe :
	SpotWebsocketV1PublicV2EventSubscribe = "sub"
	// SpotWebsocketV1PublicV2EventUnsubscribe :
	SpotWebsocketV1PublicV2EventUnsubscribe = "cancel"
)

// SpotWebsocketV1PublicV2Topic :
type SpotWebsocketV1PublicV2Topic string

const (
	// SpotWebsocketV1PublicV2TopicTrade :
	SpotWebsocketV1PublicV2TopicTrade = SpotWebsocketV1PublicV2Topic("trade")
)

// SpotWebsocketV1PublicV2TradeParamKey :
type SpotWebsocketV1PublicV2TradeParamKey struct {
	Symbol SymbolSpot
	Topic  SpotWebsocketV1PublicV2Topic
}

// SpotWebsocketV1PublicV2TradeResponse :
type SpotWebsocketV1PublicV2TradeResponse struct {
	Topic  SpotWebsocketV1PublicV2Topic               `json:"topic"`
	Params SpotWebsocketV1PublicV2TradeResponseParams `json:"params"`
	Data   SpotWebsocketV1PublicV2TradeContent        `json:"data"`
}

// SpotWebsocketV1PublicV2TradeResponseParams :
type SpotWebsocketV1PublicV2TradeResponseParams struct {
	Symbol     SymbolSpot `json:"symbol"`
	SymbolName string     `json:"symbolName"`
	Binary     string     `json:"binary"`
}

// SpotWebsocketV1PublicV2TradeContent :
type SpotWebsocketV1PublicV2TradeContent struct {
	TradeID        string `json:"v"`
	Timestamp      int    `json:"t"`
	Price          string `json:"p"`
	Quantity       string `json:"q"`
	IsBuySideTaker bool   `json:"m"`
}

// Key :
func (p *SpotWebsocketV1PublicV2TradeResponse) Key() SpotWebsocketV1PublicV2TradeParamKey {
	return SpotWebsocketV1PublicV2TradeParamKey{
		Symbol: p.Params.Symbol,
		Topic:  p.Topic,
	}
}

// SpotWebsocketV1PublicV2TradeParamChild :
type SpotWebsocketV1PublicV2TradeParamChild struct {
	Symbol SymbolSpot `json:"symbol"`
	Binary bool       `json:"binary"`
}

// SpotWebsocketV1PublicV2TradeParam :
type SpotWebsocketV1PublicV2TradeParam struct {
	Topic  SpotWebsocketV1PublicV2Topic           `json:"topic"`
	Event  SpotWebsocketV1PublicV2Event           `json:"event"`
	Params SpotWebsocketV1PublicV2TradeParamChild `json:"params"`
}

// Key :
func (p *SpotWebsocketV1PublicV2TradeParam) Key() SpotWebsocketV1PublicV2TradeParamKey {
	return SpotWebsocketV1PublicV2TradeParamKey{
		Symbol: p.Params.Symbol,
		Topic:  p.Topic,
	}
}

// addParamTradeFunc :
func (s *SpotWebsocketV1PublicV2Service) addParamTradeFunc(param SpotWebsocketV1PublicV2TradeParamKey, f func(SpotWebsocketV1PublicV2TradeResponse) error) error {
	if _, exist := s.paramTradeMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramTradeMap[param] = f
	return nil
}

// removeParamTradeFunc :
func (s *SpotWebsocketV1PublicV2Service) removeParamTradeFunc(key SpotWebsocketV1PublicV2TradeParamKey) {
	delete(s.paramTradeMap, key)
}

// retrieveTradeFunc :
func (s *SpotWebsocketV1PublicV2Service) retrieveTradeFunc(key SpotWebsocketV1PublicV2TradeParamKey) (func(SpotWebsocketV1PublicV2TradeResponse) error, error) {
	f, exist := s.paramTradeMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}

// judgeTopic :
func (s *SpotWebsocketV1PublicV2Service) judgeTopic(respBody []byte) (SpotWebsocketV1PublicV2Topic, error) {
	result := struct {
		Topic SpotWebsocketV1PublicV2Topic `json:"topic"`
		Event SpotWebsocketV1PublicV2Event `json:"event"`
	}{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}
	if result.Event == SpotWebsocketV1PublicV2EventSubscribe {
		return "", nil
	}
	return result.Topic, nil
}

// parseResponse :
func (s *SpotWebsocketV1PublicV2Service) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// SubscribeTrade :
func (s *SpotWebsocketV1PublicV2Service) SubscribeTrade(symbol SymbolSpot, f func(response SpotWebsocketV1PublicV2TradeResponse) error) (func() error, error) {
	param := SpotWebsocketV1PublicV2TradeParam{
		Topic: SpotWebsocketV1PublicV2TopicTrade,
		Event: SpotWebsocketV1PublicV2EventSubscribe,
		Params: SpotWebsocketV1PublicV2TradeParamChild{
			Binary: false,
			Symbol: symbol,
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
		param.Event = SpotWebsocketV1PublicV2EventUnsubscribe
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
func (s *SpotWebsocketV1PublicV2Service) Start(ctx context.Context) {
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
func (s *SpotWebsocketV1PublicV2Service) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeTopic(message)
	if err != nil {
		return err
	}
	switch topic {
	case SpotWebsocketV1PublicV2TopicTrade:
		var resp SpotWebsocketV1PublicV2TradeResponse
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
func (s *SpotWebsocketV1PublicV2Service) Ping() error {
	if err := s.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *SpotWebsocketV1PublicV2Service) Close() error {
	if err := s.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return nil
}
