package bybit

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// CreateOrder :
func (s *V5WebsocketTradeService) CreateOrder(orders []*V5CreateOrderParam) error {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	headers := make(map[string]string)
	headers["X-BAPI-TIMESTAMP"] = timestamp
	headers["X-BAPI-RECV-WINDOW"] = "8000"

	param := struct {
		ReqId   string                `json:"reqId"`
		Headers map[string]string     `json:"header"`
		Op      string                `json:"op"`
		Args    []*V5CreateOrderParam `json:"args"`
	}{
		ReqId:   uuid.New().String(),
		Headers: headers,
		Op:      "order.create",
		Args:    orders,
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return err
	}

	if err := s.writeMessage(websocket.TextMessage, buf); err != nil {
		return err
	}
	return nil
}

func (s *V5WebsocketTradeService) CancelOrder(orders []*V5CancelOrderParam) error {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	headers := make(map[string]string)
	headers["X-BAPI-TIMESTAMP"] = timestamp
	headers["X-BAPI-RECV-WINDOW"] = "8000"

	param := struct {
		ReqId   string                `json:"reqId"`
		Headers map[string]string     `json:"header"`
		Op      string                `json:"op"`
		Args    []*V5CancelOrderParam `json:"args"`
	}{
		ReqId:   uuid.New().String(),
		Headers: headers,
		Op:      "order.cancel",
		Args:    orders,
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return err
	}

	if err := s.writeMessage(websocket.TextMessage, buf); err != nil {
		return err
	}
	return nil
}
