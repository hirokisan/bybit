package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// SubscribeLiquidation :
func (s *V5WebsocketPublicService) SubscribeLiquidation(
	key V5WebsocketPublicLiquidationParamKey,
	f func(V5WebsocketPublicLiquidationResponse) error,
) (func() error, error) {
	if err := s.addParamLiquidationFunc(key, f); err != nil {
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
		s.removeParamLiquidationFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicLiquidationParamKey :
type V5WebsocketPublicLiquidationParamKey struct {
	Symbol SymbolV5
}

// Topic :
func (k *V5WebsocketPublicLiquidationParamKey) Topic() string {
	return fmt.Sprintf("%s.%s", V5WebsocketPublicTopicLiquidation, k.Symbol)
}

// V5WebsocketPublicLiquidationResponse :
type V5WebsocketPublicLiquidationResponse struct {
	Topic     string                             `json:"topic"`
	Type      string                             `json:"type"`
	TimeStamp int64                              `json:"ts"`
	Data      []V5WebsocketPublicLiquidationData `json:"data"`
}

// V5WebsocketPublicLiquidationData :
type V5WebsocketPublicLiquidationData struct {
	UpdatedTime uint64   `json:"updatedTime"` // The updated timestamp (ms)
	Symbol      SymbolV5 `json:"symbol"`      // Symbol name
	Side        Side     `json:"side"`        // Position side. Buy,Sell. When you receive a Buy update, this means that a long position has been liquidated
	Size        string   `json:"size"`        // Executed size
	Price       string   `json:"price"`       // Bankruptcy price
}

// Key :
func (r *V5WebsocketPublicLiquidationResponse) Key() V5WebsocketPublicLiquidationParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicLiquidation.String() || len(arr) != 2 {
		return V5WebsocketPublicLiquidationParamKey{}
	}

	return V5WebsocketPublicLiquidationParamKey{
		Symbol: SymbolV5(arr[1]),
	}
}

// addParamLiquidationFunc :
func (s *V5WebsocketPublicService) addParamLiquidationFunc(key V5WebsocketPublicLiquidationParamKey, f func(V5WebsocketPublicLiquidationResponse) error) error {
	if _, exist := s.paramLiquidationMap[key]; exist {
		return errors.New("already registered for this key")
	}
	s.paramLiquidationMap[key] = f
	return nil
}

// removeParamLiquidationFunc :
func (s *V5WebsocketPublicService) removeParamLiquidationFunc(key V5WebsocketPublicLiquidationParamKey) {
	delete(s.paramLiquidationMap, key)
}

// retrievePositionFunc :
func (s *V5WebsocketPublicService) retrieveLiquidationFunc(key V5WebsocketPublicLiquidationParamKey) (func(V5WebsocketPublicLiquidationResponse) error, error) {
	f, exist := s.paramLiquidationMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
