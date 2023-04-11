package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// SubscribeKline :
func (s *V5WebsocketPublicService) SubscribeKline(
	key V5WebsocketPublicKlineParamKey,
	f func(V5WebsocketPublicKlineResponse) error,
) (func() error, error) {
	if err := s.addParamKlineFunc(key, f); err != nil {
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
		s.removeParamKlineFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicKlineParamKey :
type V5WebsocketPublicKlineParamKey struct {
	Interval Interval
	Symbol   SymbolV5
}

// Topic :
func (k *V5WebsocketPublicKlineParamKey) Topic() string {
	return fmt.Sprintf("%s.%s.%s", V5WebsocketPublicTopicKline, k.Interval, k.Symbol)
}

// V5WebsocketPublicKlineResponse :
type V5WebsocketPublicKlineResponse struct {
	Topic     string                       `json:"topic"`
	Type      string                       `json:"type"`
	TimeStamp int64                        `json:"ts"`
	Data      []V5WebsocketPublicKlineData `json:"data"`
}

// V5WebsocketPublicKlineData :
type V5WebsocketPublicKlineData struct {
	Start     int      `json:"start"`
	End       int      `json:"end"`
	Interval  Interval `json:"interval"`
	Open      string   `json:"open"`
	Close     string   `json:"close"`
	High      string   `json:"high"`
	Low       string   `json:"low"`
	Volume    string   `json:"volume"`
	Turnover  string   `json:"turnover"`
	Confirm   bool     `json:"confirm"`
	Timestamp int      `json:"timestamp"`
}

// Key :
func (r *V5WebsocketPublicKlineResponse) Key() V5WebsocketPublicKlineParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicKline.String() || len(arr) != 3 {
		return V5WebsocketPublicKlineParamKey{}
	}

	return V5WebsocketPublicKlineParamKey{
		Interval: Interval(arr[1]),
		Symbol:   SymbolV5(arr[2]),
	}
}

// addParamKlineFunc :
func (s *V5WebsocketPublicService) addParamKlineFunc(key V5WebsocketPublicKlineParamKey, f func(V5WebsocketPublicKlineResponse) error) error {
	if _, exist := s.paramKlineMap[key]; exist {
		return errors.New("already registered for this key")
	}
	s.paramKlineMap[key] = f
	return nil
}

// removeParamTradeFunc :
func (s *V5WebsocketPublicService) removeParamKlineFunc(key V5WebsocketPublicKlineParamKey) {
	delete(s.paramKlineMap, key)
}

// retrievePositionFunc :
func (s *V5WebsocketPublicService) retrieveKlineFunc(key V5WebsocketPublicKlineParamKey) (func(V5WebsocketPublicKlineResponse) error, error) {
	f, exist := s.paramKlineMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
