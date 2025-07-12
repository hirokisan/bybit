package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// SubscribeAllLiquidation :
func (s *V5WebsocketPublicService) SubscribeAllLiquidation(
	key V5WebsocketPublicAllLiquidationParamKey,
	f func(V5WebsocketPublicAllLiquidationResponse) error,
) (func() error, error) {
	if err := s.addParamAllLiquidationFunc(key, f); err != nil {
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
		s.removeParamAllLiquidationFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicAllLiquidationParamKey :
type V5WebsocketPublicAllLiquidationParamKey struct {
	Symbol SymbolV5
}

// Topic :
func (k *V5WebsocketPublicAllLiquidationParamKey) Topic() string {
	return fmt.Sprintf("%s.%s", V5WebsocketPublicTopicAllLiquidation, k.Symbol)
}

// V5WebsocketPublicAllLiquidationResponse :
type V5WebsocketPublicAllLiquidationResponse struct {
	Topic     string                                `json:"topic"`
	Type      string                                `json:"type"`
	TimeStamp int64                                 `json:"ts"`
	Data      []V5WebsocketPublicAllLiquidationData `json:"data"`
}

// V5WebsocketPublicAllLiquidationData :
type V5WebsocketPublicAllLiquidationData struct {
	UpdatedTime uint64   `json:"T"` // The updated timestamp (ms)
	Symbol      SymbolV5 `json:"s"` // Symbol name
	Side        Side     `json:"S"` // Position side. Buy,Sell. When you receive a Buy update, this means that a long position has been liquidated
	Size        string   `json:"v"` // Executed size
	Price       string   `json:"p"` // Bankruptcy price
}

// Key :
func (r *V5WebsocketPublicAllLiquidationResponse) Key() V5WebsocketPublicAllLiquidationParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicAllLiquidation.String() || len(arr) != 2 {
		return V5WebsocketPublicAllLiquidationParamKey{}
	}

	return V5WebsocketPublicAllLiquidationParamKey{
		Symbol: SymbolV5(arr[1]),
	}
}

// addParamAllLiquidationFunc :
func (s *V5WebsocketPublicService) addParamAllLiquidationFunc(key V5WebsocketPublicAllLiquidationParamKey, f func(V5WebsocketPublicAllLiquidationResponse) error) error {
	if _, exist := s.paramAllLiquidationMap[key]; exist {
		return errors.New("already registered for this key")
	}
	s.paramAllLiquidationMap[key] = f
	return nil
}

// removeParamAllLiquidationFunc :
func (s *V5WebsocketPublicService) removeParamAllLiquidationFunc(key V5WebsocketPublicAllLiquidationParamKey) {
	delete(s.paramAllLiquidationMap, key)
}

// retrieveAllLiquidationFunc :
func (s *V5WebsocketPublicService) retrieveAllLiquidationFunc(key V5WebsocketPublicAllLiquidationParamKey) (func(V5WebsocketPublicAllLiquidationResponse) error, error) {
	f, exist := s.paramAllLiquidationMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
