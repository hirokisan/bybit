package bybit

import (
	"encoding/json"
	"errors"

	"github.com/gorilla/websocket"
)

// SubscribeOrder :
func (s *V5WebsocketPrivateService) SubscribeOrder(
	f func(V5WebsocketPrivateOrderResponse) error,
) (func() error, error) {
	key := V5WebsocketPrivateParamKey{
		Topic: V5WebsocketPrivateTopicOrder,
	}
	if err := s.addParamOrderFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{V5WebsocketPrivateTopicOrder},
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
			Args: []interface{}{V5WebsocketPrivateTopicOrder},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamOrderFunc(key)
		return nil
	}, nil
}

// V5WebsocketPrivateOrderResponse :
type V5WebsocketPrivateOrderResponse struct {
	ID           string                        `json:"id"`
	Topic        V5WebsocketPrivateTopic       `json:"topic"`
	CreationTime int64                         `json:"creationTime"`
	Data         []V5WebsocketPrivateOrderData `json:"data"`
}

// V5WebsocketPrivateOrderData :
type V5WebsocketPrivateOrderData struct {
	AvgPrice           string           `json:"avgPrice"`
	BlockTradeID       string           `json:"blockTradeId"`
	CancelType         string           `json:"cancelType"`
	Category           string           `json:"category"`
	CloseOnTrigger     bool             `json:"closeOnTrigger"`
	CreatedTime        string           `json:"createdTime"`
	CumExecFee         string           `json:"cumExecFee"`
	CumExecQty         string           `json:"cumExecQty"`
	CumExecValue       string           `json:"cumExecValue"`
	LeavesQty          string           `json:"leavesQty"`
	LeavesValue        string           `json:"leavesValue"`
	OrderID            string           `json:"orderId"`
	OrderIv            string           `json:"orderIv"`
	IsLeverage         string           `json:"isLeverage"`
	LastPriceOnCreated string           `json:"lastPriceOnCreated"`
	OrderStatus        OrderStatus      `json:"orderStatus"`
	OrderLinkID        string           `json:"orderLinkId"`
	OrderType          OrderType        `json:"orderType"`
	PositionIdx        int              `json:"positionIdx"`
	Price              string           `json:"price"`
	Qty                string           `json:"qty"`
	ReduceOnly         bool             `json:"reduceOnly"`
	RejectReason       string           `json:"rejectReason"`
	Side               Side             `json:"side"`
	SlTriggerBy        TriggerBy        `json:"slTriggerBy"`
	StopLoss           string           `json:"stopLoss"`
	StopOrderType      string           `json:"stopOrderType"`
	Symbol             SymbolV5         `json:"symbol"`
	TakeProfit         string           `json:"takeProfit"`
	TimeInForce        TimeInForce      `json:"timeInForce"`
	TpTriggerBy        TriggerBy        `json:"tpTriggerBy"`
	TriggerBy          TriggerBy        `json:"triggerBy"`
	TriggerDirection   TriggerDirection `json:"triggerDirection"`
	TriggerPrice       string           `json:"triggerPrice"`
	UpdatedTime        string           `json:"updatedTime"`
}

// Key :
func (r *V5WebsocketPrivateOrderResponse) Key() V5WebsocketPrivateParamKey {
	return V5WebsocketPrivateParamKey{
		Topic: r.Topic,
	}
}

// addParamOrderFunc :
func (s *V5WebsocketPrivateService) addParamOrderFunc(param V5WebsocketPrivateParamKey, f func(V5WebsocketPrivateOrderResponse) error) error {
	if _, exist := s.paramOrderMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramOrderMap[param] = f
	return nil
}

// removeParamOrderFunc :
func (s *V5WebsocketPrivateService) removeParamOrderFunc(key V5WebsocketPrivateParamKey) {
	delete(s.paramOrderMap, key)
}

// retrieveOrderFunc :
func (s *V5WebsocketPrivateService) retrieveOrderFunc(key V5WebsocketPrivateParamKey) (func(V5WebsocketPrivateOrderResponse) error, error) {
	f, exist := s.paramOrderMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
