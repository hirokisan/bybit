package bybit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateLinearOrderResponse :
type CreateLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateLinearOrderResult `json:"result"`
}

// CreateLinearOrderResult :
type CreateLinearOrderResult struct {
	CreateLinearOrder `json:",inline"`
}

// CreateLinearOrder :
type CreateLinearOrder struct {
	OrderID        string      `json:"order_id"`
	UserID         int         `json:"user_id"`
	Symbol         SymbolUSDT  `json:"symbol"`
	Side           Side        `json:"side"`
	OrderType      OrderType   `json:"order_type"`
	Price          float64     `json:"price"`
	Qty            float64     `json:"qty"`
	TimeInForce    TimeInForce `json:"time_in_force"`
	OrderStatus    OrderStatus `json:"order_status"`
	LastExecPrice  float64     `json:"last_exec_price"`
	CumExecQty     float64     `json:"cum_exec_qty"`
	CumExecValue   float64     `json:"cum_exec_value"`
	CumExecFee     float64     `json:"cum_exec_fee"`
	ReduceOnly     bool        `json:"reduce_only"`
	CloseOnTrigger bool        `json:"close_on_trigger"`
	OrderLinkID    string      `json:"order_link_id"`
	CreatedTime    string      `json:"created_time"`
	UpdatedTime    string      `json:"updated_time"`
	TakeProfit     float64     `json:"take_profit"`
	StopLoss       float64     `json:"stop_loss"`
	TpTriggerBy    string      `json:"tp_trigger_by"`
	SlTriggerBy    string      `json:"sl_trigger_by"`
}

// CreateLinearOrderParam :
type CreateLinearOrderParam struct {
	Side           Side        `json:"side"`
	Symbol         SymbolUSDT  `json:"symbol"`
	OrderType      OrderType   `json:"order_type"`
	Qty            float64     `json:"qty"`
	TimeInForce    TimeInForce `json:"time_in_force"`
	ReduceOnly     bool        `json:"reduce_only"`
	CloseOnTrigger bool        `json:"close_on_trigger"`

	Price       *float64 `json:"price,omitempty"`
	TakeProfit  *float64 `json:"take_profit,omitempty"`
	StopLoss    *float64 `json:"stop_loss,omitempty"`
	TpTriggerBy *string  `json:"tp_trigger_by"`
	SlTriggerBy *string  `json:"sl_trigger_by"`
	OrderLinkID *string  `json:"order_link_id,omitempty"`
}

// CreateLinearOrder :
func (s *AccountService) CreateLinearOrder(param CreateLinearOrderParam) (*CreateLinearOrderResponse, error) {
	var res CreateLinearOrderResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/order/create", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateLinearOrderParam: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
