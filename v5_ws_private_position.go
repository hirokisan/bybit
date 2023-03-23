package bybit

import (
	"encoding/json"
	"errors"

	"github.com/gorilla/websocket"
)

// SubscribePosition :
func (s *V5WebsocketPrivateService) SubscribePosition(
	f func(V5WebsocketPrivatePositionResponse) error,
) (func() error, error) {
	key := V5WebsocketPrivateParamKey{
		Topic: V5WebsocketPrivateTopicPosition,
	}
	if err := s.addParamPositionFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{V5WebsocketPrivateTopicPosition},
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
			Args: []interface{}{V5WebsocketPrivateTopicPosition},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamPositionFunc(key)
		return nil
	}, nil
}

// V5WebsocketPrivatePositionResponse :
type V5WebsocketPrivatePositionResponse struct {
	ID           string                           `json:"id"`
	Topic        V5WebsocketPrivateTopic          `json:"topic"`
	CreationTime int64                            `json:"creationTime"`
	Data         []V5WebsocketPrivatePositionData `json:"data"`
}

// V5WebsocketPrivatePositionData :
type V5WebsocketPrivatePositionData struct {
	AutoAddMargin   int        `json:"autoAddMargin"`
	PositionIdx     int        `json:"positionIdx"`
	TradeMode       int        `json:"tradeMode"`
	RiskID          int        `json:"riskId"`
	RiskLimitValue  string     `json:"riskLimitValue"`
	Symbol          SymbolV5   `json:"symbol"`
	Side            Side       `json:"side"`
	Size            string     `json:"size"`
	EntryPrice      string     `json:"entryPrice"`
	Leverage        string     `json:"leverage"`
	PositionValue   string     `json:"positionValue"`
	MarkPrice       string     `json:"markPrice"`
	PositionBalance string     `json:"positionBalance"`
	PositionIM      string     `json:"positionIM"`
	PositionMM      string     `json:"positionMM"`
	TakeProfit      string     `json:"takeProfit"`
	StopLoss        string     `json:"stopLoss"`
	TrailingStop    string     `json:"trailingStop"`
	UnrealisedPnl   string     `json:"unrealisedPnl"`
	CumRealisedPnl  string     `json:"cumRealisedPnl"`
	CreatedTime     string     `json:"CreatedTime"`
	UpdatedTime     string     `json:"updatedTime"`
	TpslMode        TpSlMode   `json:"tpslMode"`
	LiqPrice        string     `json:"liqPrice"`
	BustPrice       string     `json:"bustPrice"`
	Category        CategoryV5 `json:"category"`
	PositionStatus  string     `json:"positionStatus"`
}

// Key :
func (r *V5WebsocketPrivatePositionResponse) Key() V5WebsocketPrivateParamKey {
	return V5WebsocketPrivateParamKey{
		Topic: r.Topic,
	}
}

// addParamPositionFunc :
func (s *V5WebsocketPrivateService) addParamPositionFunc(param V5WebsocketPrivateParamKey, f func(V5WebsocketPrivatePositionResponse) error) error {
	if _, exist := s.paramPositionMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramPositionMap[param] = f
	return nil
}

// removeParamPositionFunc :
func (s *V5WebsocketPrivateService) removeParamPositionFunc(key V5WebsocketPrivateParamKey) {
	delete(s.paramPositionMap, key)
}

// retrievePositionFunc :
func (s *V5WebsocketPrivateService) retrievePositionFunc(key V5WebsocketPrivateParamKey) (func(V5WebsocketPrivatePositionResponse) error, error) {
	f, exist := s.paramPositionMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
