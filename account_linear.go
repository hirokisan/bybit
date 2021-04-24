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

// ListLinearPositionResponse :
type ListLinearPositionResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListLinearPositionResult `json:"result"`
}

// ListLinearPositionResult :
type ListLinearPositionResult struct {
	UserID              int           `json:"user_id"`
	Symbol              SymbolInverse `json:"symbol"`
	Side                Side          `json:"side"`
	Size                float64       `json:"size"`
	PositionValue       float64       `json:"position_value"`
	EntryPrice          float64       `json:"entry_price"`
	LiqPrice            float64       `json:"liq_price"`
	BustPrice           float64       `json:"bust_price"`
	Leverage            float64       `json:"leverage"`
	AutoAddMargin       float64       `json:"auto_add_margin"`
	IsIsolated          bool          `json:"is_isolated"`
	PositionMargin      float64       `json:"position_margin"`
	OccClosingFee       float64       `json:"occ_closing_fee"`
	RealisedPnl         float64       `json:"realised_pnl"`
	CumRealisedPnl      float64       `json:"cum_realised_pnl"`
	FreeQty             float64       `json:"free_qty"`
	TpSlMode            TpSlMode      `json:"tp_sl_mode"`
	DeleverageIndicator int           `json:"deleverage_indicator"`
	UnrealisedPnl       float64       `json:"unrealised_pnl"`
	RiskID              int           `json:"risk_id"`
}

// ListLinearPosition :
func (s *AccountService) ListLinearPosition(symbol SymbolUSDT) (*ListLinearPositionResponse, error) {
	var res ListLinearPositionResponse

	params := map[string]string{
		"symbol": string(symbol),
	}
	url, err := s.Client.BuildPrivateURL("/private/linear/position/list", params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListLinearPositionsResponse :
type ListLinearPositionsResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListLinearPositionsResult `json:"result"`
}

// ListLinearPositionsResult :
type ListLinearPositionsResult struct {
	IsValid                  bool `json:"is_valid"`
	ListLinearPositionResult `json:"data,inline"`
}

// ListLinearPositions :
func (s *AccountService) ListLinearPositions() (*ListLinearPositionsResponse, error) {
	var res ListLinearPositionsResponse

	url, err := s.Client.BuildPrivateURL("/private/linear/position/list", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CancelLinearOrderResponse :
type CancelLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelLinearOrderResult `json:"result"`
}

// CancelLinearOrderResult :
type CancelLinearOrderResult struct {
	CancelLinearOrder `json:",inline"`
}

// CancelLinearOrder :
type CancelLinearOrder struct {
	OrderID string `json:"order_id"`
}

// CancelLinearOrderParam :
type CancelLinearOrderParam struct {
	Symbol SymbolUSDT `json:"symbol"`

	OrderID     *string `json:"order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelLinearOrder :
func (s *AccountService) CancelLinearOrder(param CancelLinearOrderParam) (*CancelLinearOrderResponse, error) {
	var res CancelLinearOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	url, err := s.Client.BuildPrivateURL("/private/linear/order/cancel", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelLinearOrderParam: %w", err)
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
