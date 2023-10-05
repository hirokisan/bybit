package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// SubscribeTrade :
func (s *V5WebsocketPublicService) SubscribeTrade(
	key V5WebsocketPublicTradeParamKey,
	f func(V5WebsocketPublicTradeResponse) error,
) (func() error, error) {
	if err := s.addParamTradeFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{key.Topic()},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	if err := s.writeMessage(websocket.TextMessage, buf); err != nil {
		return nil, err
	}
	return func() error {
		param := struct {
			Op   string        `json:"op"`
			Args []interface{} `json:"args"`
		}{
			Op:   "unsubscribe",
			Args: []interface{}{key.Topic()},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamTradeFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicTradeParamKey :
type V5WebsocketPublicTradeParamKey struct {
	Symbol SymbolV5
}

// Topic :
func (k *V5WebsocketPublicTradeParamKey) Topic() string {
	return fmt.Sprintf("%s.%s", V5WebsocketPublicTopicTrade, k.Symbol)
}

// V5WebsocketPublicTradeResponse :
type V5WebsocketPublicTradeResponse struct {
	Topic     string                       `json:"topic"`
	Type      string                       `json:"type"`
	TimeStamp int64                        `json:"ts"`
	Data      []V5WebsocketPublicTradeData `json:"data"`
}

// V5WebsocketPublicTradeData :
type V5WebsocketPublicTradeData struct {
	Timestamp  uint64   `json:"T"`  // The timestamp (ms) that the order is filled
	Symbol     SymbolV5 `json:"s"`  // Symbol name
	Side       Side     `json:"S"`  // Side of taker. Buy,Sell
	Value      string   `json:"v"`  // Trade size
	Trade      string   `json:"p"`  // Trade price
	Direction  string   `json:"L"`  // Direction of price change. Unique field for future
	ID         string   `json:"i"`  // Trade ID
	BlockTrade bool     `json:"BT"` // Whether it is a block trade order or not
}

// Key :
func (r *V5WebsocketPublicTradeResponse) Key() V5WebsocketPublicTradeParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicTrade.String() || len(arr) != 2 {
		return V5WebsocketPublicTradeParamKey{}
	}

	return V5WebsocketPublicTradeParamKey{
		Symbol: SymbolV5(arr[1]),
	}
}

// addParamTradeFunc :
func (s *V5WebsocketPublicService) addParamTradeFunc(key V5WebsocketPublicTradeParamKey, f func(V5WebsocketPublicTradeResponse) error) error {
	if _, exist := s.paramTradeMap[key]; exist {
		return errors.New("already registered for this key")
	}
	s.paramTradeMap[key] = f
	return nil
}

// removeParamTradeFunc :
func (s *V5WebsocketPublicService) removeParamTradeFunc(key V5WebsocketPublicTradeParamKey) {
	delete(s.paramTradeMap, key)
}

// retrievePositionFunc :
func (s *V5WebsocketPublicService) retrieveTradeFunc(key V5WebsocketPublicTradeParamKey) (func(V5WebsocketPublicTradeResponse) error, error) {
	f, exist := s.paramTradeMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
