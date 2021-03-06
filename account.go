package bybit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AccountService :
type AccountService struct {
	Client *Client
}

// CreateOrderResponse :
type CreateOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateOrderResult `json:"result"`
}

// CreateOrderResult :
type CreateOrderResult struct {
	CreateOrder `json:",inline"`
}

// CreateOrder :
type CreateOrder struct {
	UserID        int         `json:"user_id"`
	OrderID       string      `json:"order_id"`
	Symbol        Symbol      `json:"symbol"`
	Side          Side        `json:"side"`
	OrderType     OrderType   `json:"order_type"`
	Price         float64     `json:"price"`
	Qty           float64     `json:"qty"`
	TimeInForce   TimeInForce `json:"time_in_force"`
	OrderStatus   OrderStatus `json:"order_status"`
	LastExecTime  int         `json:"last_exec_time"`
	LastExecPrice float64     `json:"last_exec_price"`
	LeavesQty     float64     `json:"leaves_qty"`
	CumExecQty    float64     `json:"cum_exec_qty"`
	CumExecValue  float64     `json:"cum_exec_value"`
	CumExecFee    float64     `json:"cum_exec_fee"`
	RejectReason  string      `json:"reject_reason"`
	OrderLinkID   string      `json:"order_link_id"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

// CreateOrderParam :
type CreateOrderParam struct {
	Side        Side        `json:"side"`
	Symbol      Symbol      `json:"symbol"`
	OrderType   OrderType   `json:"order_type"`
	Qty         int         `json:"qty"`
	TimeInForce TimeInForce `json:"time_in_force"`

	Price          *float64 `json:"price,omitempty"`
	TakeProfit     *float64 `json:"take_profit,omitempty"`
	StopLoss       *float64 `json:"stop_loss,omitempty"`
	ReduceOnly     *bool    `json:"reduce_only,omitempty"`
	CloseOnTrigger *bool    `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string  `json:"order_link_id,omitempty"`
}

// param -> json -> body -> post

// CreateOrder :
func (s *AccountService) CreateOrder(param CreateOrderParam) (*CreateOrderResponse, error) {
	var res CreateOrderResponse

	if !s.Client.HasAuth() {
		return nil, fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	url := s.Client.BuildURL("/v2/private/order/create", nil)

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateOrderParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Printf("%v", resp.Request)
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
