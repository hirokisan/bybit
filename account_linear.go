package bybit

import (
	"encoding/json"
	"fmt"
	"net/url"
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

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateLinearOrderParam: %w", err)
	}

	if err := s.Client.postJSON("/private/linear/order/create", body, &res); err != nil {
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

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.Client.getPrivately("/private/linear/position/list", query, &res); err != nil {
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

	if err := s.Client.getPrivately("/private/linear/position/list", nil, &res); err != nil {
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

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelLinearOrderParam: %w", err)
	}

	if err := s.Client.postJSON("/private/linear/order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SaveLinearLeverageResponse :
type SaveLinearLeverageResponse struct {
	CommonResponse `json:",inline"`
}

// SaveLinearLeverageParam :
type SaveLinearLeverageParam struct {
	Symbol       SymbolUSDT `json:"symbol"`
	BuyLeverage  float64    `json:"buy_leverage"`
	SellLeverage float64    `json:"sell_leverage"`
}

// SaveLinearLeverage :
func (s *AccountService) SaveLinearLeverage(param SaveLinearLeverageParam) (*SaveLinearLeverageResponse, error) {
	var res SaveLinearLeverageResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for SaveLinearLeverageParam: %w", err)
	}

	if err := s.Client.postJSON("/private/linear/position/set-leverage", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// LinearExecutionListResponse :
type LinearExecutionListResponse struct {
	CommonResponse `json:",inline"`
	Result         LinearExecutionListResult `json:"result"`
}

// LinearExecutionListResult :
type LinearExecutionListResult struct {
	CurrentPage          int                   `json:"current_page"`
	LinearExecutionLists []LinearExecutionList `json:"data"`
}

// LinearExecutionList :
type LinearExecutionList struct {
	OrderID          string     `json:"order_id"`
	OrderLinkID      string     `json:"order_link_id"`
	Side             Side       `json:"side"`
	Symbol           SymbolUSDT `json:"symbol"`
	OrderPrice       float64    `json:"order_price"`
	OrderQty         float64    `json:"order_qty"`
	OrderType        OrderType  `json:"order_type"`
	FeeRate          float64    `json:"fee_rate"`
	ExecPrice        float64    `json:"exec_price"`
	ExecType         ExecType   `json:"exec_type"`
	ExecQty          float64    `json:"exec_qty"`
	ExecFee          float64    `json:"exec_fee"`
	ExecValue        float64    `json:"exec_value"`
	LeavesQty        float64    `json:"leaves_qty"`
	ClosedSize       float64    `json:"closed_size"`
	LastLiquidityInd string     `json:"last_liquidity_ind"`
	TradeTimeMs      float64    `json:"trade_time_ms"`
}

// LinearExecutionListParam :
type LinearExecutionListParam struct {
	Symbol SymbolUSDT `json:"symbol"`

	StartTime *int      `json:"start_time"`
	EndTime   *int      `json:"end_time"`
	ExecType  *ExecType `json:"exec_type"`
	Page      *int      `json:"page"`
	Limit     *int      `json:"limit"`
}

// LinearExecutionList :
// NOTE(TODO) : somehow got EOF 404(path not found)
func (s *AccountService) LinearExecutionList(param LinearExecutionListParam) (*LinearExecutionListResponse, error) {
	var res LinearExecutionListResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for LinearExecutionListParam: %w", err)
	}

	if err := s.Client.postJSON("/private/linear/trade/execution/list", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
