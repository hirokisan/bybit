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
	UserID        int           `json:"user_id"`
	OrderID       string        `json:"order_id"`
	Symbol        SymbolInverse `json:"symbol"`
	Side          Side          `json:"side"`
	OrderType     OrderType     `json:"order_type"`
	Price         float64       `json:"price"`
	Qty           float64       `json:"qty"`
	TimeInForce   TimeInForce   `json:"time_in_force"`
	OrderStatus   OrderStatus   `json:"order_status"`
	LastExecTime  float64       `json:"last_exec_time"`
	LastExecPrice float64       `json:"last_exec_price"`
	LeavesQty     float64       `json:"leaves_qty"`
	CumExecQty    float64       `json:"cum_exec_qty"`
	CumExecValue  float64       `json:"cum_exec_value"`
	CumExecFee    float64       `json:"cum_exec_fee"`
	RejectReason  string        `json:"reject_reason"`
	OrderLinkID   string        `json:"order_link_id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

// CreateOrderParam :
type CreateOrderParam struct {
	Side        Side          `json:"side"`
	Symbol      SymbolInverse `json:"symbol"`
	OrderType   OrderType     `json:"order_type"`
	Qty         int           `json:"qty"`
	TimeInForce TimeInForce   `json:"time_in_force"`

	Price          *float64 `json:"price,omitempty"`
	TakeProfit     *float64 `json:"take_profit,omitempty"`
	StopLoss       *float64 `json:"stop_loss,omitempty"`
	ReduceOnly     *bool    `json:"reduce_only,omitempty"`
	CloseOnTrigger *bool    `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string  `json:"order_link_id,omitempty"`
}

// CreateOrder :
func (s *AccountService) CreateOrder(param CreateOrderParam) (*CreateOrderResponse, error) {
	var res CreateOrderResponse

	url, err := s.Client.BuildPrivateURL("/v2/private/order/create", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateOrderParam: %w", err)
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

// ListPositionResponse :
type ListPositionResponse struct {
	CommonResponse `json:",inline"`
	Result         ListPositionResult `json:"result"`
}

// ListPositionResult :
type ListPositionResult struct {
	ID                  int           `json:"id"`
	UserID              int           `json:"user_id"`
	RiskID              int           `json:"risk_id"`
	Symbol              SymbolInverse `json:"symbol"`
	Side                Side          `json:"side"`
	Size                float64       `json:"size"`
	PositionValue       string        `json:"position_value"`
	EntryPrice          string        `json:"entry_price"`
	IsIsolated          bool          `json:"is_isolated"`
	AutoAddMargin       float64       `json:"auto_add_margin"`
	Leverage            string        `json:"leverage"`
	EffectiveLeverage   string        `json:"effective_leverage"`
	PositionMargin      string        `json:"position_margin"`
	LiqPrice            string        `json:"liq_price"`
	BustPrice           string        `json:"bust_price"`
	OccClosingFee       string        `json:"occ_closing_fee"`
	OccFundingFee       string        `json:"occ_funding_fee"`
	TakeProfit          string        `json:"take_profit"`
	StopLoss            string        `json:"stop_loss"`
	TrailingStop        string        `json:"trailing_stop"`
	PositionStatus      string        `json:"position_status"`
	DeleverageIndicator int           `json:"deleverage_indicator"`
	OcCalcData          string        `json:"oc_calc_data"`
	OrderMargin         string        `json:"order_margin"`
	WalletBalance       string        `json:"wallet_balance"`
	RealisedPnl         string        `json:"realised_pnl"`
	UnrealisedPnl       float64       `json:"unrealised_pnl"`
	CumRealisedPnl      string        `json:"cum_realised_pnl"`
	CrossSeq            float64       `json:"cross_seq"`
	PositionSeq         float64       `json:"position_seq"`
	CreatedAt           string        `json:"created_at"`
	UpdatedAt           string        `json:"updated_at"`
}

// ListPosition :
func (s *AccountService) ListPosition(symbol SymbolInverse) (*ListPositionResponse, error) {
	var res ListPositionResponse

	params := map[string]string{
		"symbol": string(symbol),
	}
	url, err := s.Client.BuildPrivateURL("/v2/private/position/list", params)
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

// ListPositionsResponse :
type ListPositionsResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListPositionsResult `json:"result"`
}

// ListPositionsResult :
type ListPositionsResult struct {
	IsValid            bool `json:"is_valid"`
	ListPositionResult `json:"data,inline"`
}

// ListPositions :
func (s *AccountService) ListPositions() (*ListPositionsResponse, error) {
	var res ListPositionsResponse

	url, err := s.Client.BuildPrivateURL("/v2/private/position/list", nil)
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

// CancelOrderResponse :
type CancelOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelOrderResult `json:"result"`
}

// CancelOrderResult :
type CancelOrderResult struct {
	CancelOrder `json:",inline"`
}

// CancelOrder :
// so far, same as CreateOrder
type CancelOrder struct {
	UserID        int           `json:"user_id"`
	OrderID       string        `json:"order_id"`
	Symbol        SymbolInverse `json:"symbol"`
	Side          Side          `json:"side"`
	OrderType     OrderType     `json:"order_type"`
	Price         float64       `json:"price"`
	Qty           float64       `json:"qty"`
	TimeInForce   TimeInForce   `json:"time_in_force"`
	OrderStatus   OrderStatus   `json:"order_status"`
	LastExecTime  float64       `json:"last_exec_time"`
	LastExecPrice float64       `json:"last_exec_price"`
	LeavesQty     float64       `json:"leaves_qty"`
	CumExecQty    float64       `json:"cum_exec_qty"`
	CumExecValue  float64       `json:"cum_exec_value"`
	CumExecFee    float64       `json:"cum_exec_fee"`
	RejectReason  string        `json:"reject_reason"`
	OrderLinkID   string        `json:"order_link_id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

// CancelOrderParam :
type CancelOrderParam struct {
	Symbol SymbolInverse `json:"symbol"`

	OrderID     *string `json:"order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelOrder :
func (s *AccountService) CancelOrder(param CancelOrderParam) (*CancelOrderResponse, error) {
	var res CancelOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	url, err := s.Client.BuildPrivateURL("/v2/private/order/cancel", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelOrderParam: %w", err)
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

// SaveLeverageResponse :
type SaveLeverageResponse struct {
	CommonResponse `json:",inline"`
	Result         float64 `json:"result"`
}

// SaveLeverageParam :
type SaveLeverageParam struct {
	Symbol   SymbolInverse `json:"symbol"`
	Leverage float64       `json:"leverage"`
}

// SaveLeverage :
func (s *AccountService) SaveLeverage(param SaveLeverageParam) (*SaveLeverageResponse, error) {
	var res SaveLeverageResponse

	url, err := s.Client.BuildPrivateURL("/v2/private/position/leverage/save", nil)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelOrderParam: %w", err)
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
