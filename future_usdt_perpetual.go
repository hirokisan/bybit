package bybit

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// FutureUSDTPerpetualServiceI :
type FutureUSDTPerpetualServiceI interface {
	// Market Data Endpoints
	OrderBook(SymbolFuture) (*OrderBookResponse, error)
	ListLinearKline(ListLinearKlineParam) (*ListLinearKlineResponse, error)
	Tickers(SymbolFuture) (*TickersResponse, error)
	Symbols() (*SymbolsResponse, error)
	OpenInterest(OpenInterestParam) (*OpenInterestResponse, error)
	BigDeal(BigDealParam) (*BigDealResponse, error)
	AccountRatio(AccountRatioParam) (*AccountRatioResponse, error)

	// Account Data Endpoints
	CreateLinearOrder(CreateLinearOrderParam) (*CreateLinearOrderResponse, error)
	ListLinearOrder(ListLinearOrderParam) (*ListLinearOrderResponse, error)
	CancelLinearOrder(CancelLinearOrderParam) (*CancelLinearOrderResponse, error)
	LinearCancelAllOrder(LinearCancelAllParam) (*LinearCancelAllResponse, error)
	ReplaceLinearOrder(ReplaceLinearOrderParam) (*ReplaceLinearOrderResponse, error)
	QueryLinearOrder(QueryLinearOrderParam) (*QueryLinearOrderResponse, error)
	CreateLinearStopOrder(CreateLinearStopOrderParam) (*CreateLinearStopOrderResponse, error)
	ListLinearStopOrder(ListLinearStopOrderParam) (*ListLinearStopOrderResponse, error)
	CancelLinearStopOrder(CancelLinearStopOrderParam) (*CancelLinearStopOrderResponse, error)
	CancelAllLinearStopOrder(CancelAllLinearStopOrderParam) (*CancelAllLinearStopOrderResponse, error)
	QueryLinearStopOrder(QueryLinearStopOrderParam) (*QueryLinearStopOrderResponse, error)
	ListLinearPosition(SymbolFuture) (*ListLinearPositionResponse, error)
	ListLinearPositions() (*ListLinearPositionsResponse, error)
	SaveLinearLeverage(SaveLinearLeverageParam) (*SaveLinearLeverageResponse, error)
	LinearTradingStop(LinearTradingStopParam) (*LinearTradingStopResponse, error)
	LinearExecutionList(LinearExecutionListParam) (*LinearExecutionListResponse, error)
	APIKeyInfo() (*APIKeyInfoResponse, error)

	// Wallet Data Endpoints
	Balance(Coin) (*BalanceResponse, error)
}

// FutureUSDTPerpetualService :
type FutureUSDTPerpetualService struct {
	client *Client

	*FutureCommonService
}

// ListLinearKlineParam :
type ListLinearKlineParam struct {
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`

	Limit *int `url:"limit,omitempty"`
}

// ListLinearKlineResponse :
type ListLinearKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListLinearKlineResult `json:"result"`
}

// ListLinearKlineResult :
type ListLinearKlineResult struct {
	Symbol   SymbolFuture `json:"symbol"`
	Period   Period       `json:"period"`
	Interval string       `json:"interval"`
	StartAt  int          `json:"start_at"`
	OpenTime int          `json:"open_time"`
	Volume   float64      `json:"volume"`
	Open     float64      `json:"open"`
	High     float64      `json:"high"`
	Low      float64      `json:"low"`
	Close    float64      `json:"close"`
	Turnover float64      `json:"turnover"`
}

// ListLinearKline :
func (s *FutureCommonService) ListLinearKline(param ListLinearKlineParam) (*ListLinearKlineResponse, error) {
	var res ListLinearKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/public/linear/kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

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
	OrderID        string          `json:"order_id"`
	UserID         int             `json:"user_id"`
	Symbol         SymbolFuture    `json:"symbol"`
	Side           Side            `json:"side"`
	OrderType      OrderType       `json:"order_type"`
	Price          float64         `json:"price"`
	Qty            float64         `json:"qty"`
	TimeInForce    TimeInForce     `json:"time_in_force"`
	OrderStatus    OrderStatus     `json:"order_status"`
	LastExecPrice  float64         `json:"last_exec_price"`
	CumExecQty     float64         `json:"cum_exec_qty"`
	CumExecValue   float64         `json:"cum_exec_value"`
	CumExecFee     float64         `json:"cum_exec_fee"`
	ReduceOnly     bool            `json:"reduce_only"`
	CloseOnTrigger bool            `json:"close_on_trigger"`
	OrderLinkID    string          `json:"order_link_id"`
	CreatedTime    string          `json:"created_time"`
	UpdatedTime    string          `json:"updated_time"`
	TakeProfit     float64         `json:"take_profit"`
	StopLoss       float64         `json:"stop_loss"`
	TpTriggerBy    TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy    TriggerByFuture `json:"sl_trigger_by"`
}

// CreateLinearOrderParam :
type CreateLinearOrderParam struct {
	Side           Side         `json:"side"`
	Symbol         SymbolFuture `json:"symbol"`
	OrderType      OrderType    `json:"order_type"`
	Qty            float64      `json:"qty"`
	TimeInForce    TimeInForce  `json:"time_in_force"`
	ReduceOnly     bool         `json:"reduce_only"`
	CloseOnTrigger bool         `json:"close_on_trigger"`

	Price       *float64         `json:"price,omitempty"`
	TakeProfit  *float64         `json:"take_profit,omitempty"`
	StopLoss    *float64         `json:"stop_loss,omitempty"`
	TpTriggerBy *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	OrderLinkID *string          `json:"order_link_id,omitempty"`
	PositionIdx *int             `json:"position_idx,omitempty"`
}

// CreateLinearOrder :
func (s *FutureUSDTPerpetualService) CreateLinearOrder(param CreateLinearOrderParam) (*CreateLinearOrderResponse, error) {
	var res CreateLinearOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateLinearOrderParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListLinearOrderResponse :
type ListLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         ListLinearOrderResult `json:"result"`
}

// ListLinearOrderResult :
type ListLinearOrderResult struct {
	CurrentPage int                            `json:"current_page"`
	Content     []ListLinearOrderResultContent `json:"data"`
}

// ListLinearOrderResultContent :
type ListLinearOrderResultContent struct {
	OrderID        string          `json:"order_id"`
	UserID         int             `json:"user_id"`
	Symbol         SymbolFuture    `json:"symbol"`
	Side           Side            `json:"side"`
	OrderType      OrderType       `json:"order_type"`
	Price          float64         `json:"price"`
	Qty            float64         `json:"qty"`
	TimeInForce    TimeInForce     `json:"time_in_force"`
	OrderStatus    OrderStatus     `json:"order_status"`
	LastExecPrice  float64         `json:"last_exec_price"`
	CumExecQty     float64         `json:"cum_exec_qty"`
	CumExecValue   float64         `json:"cum_exec_value"`
	CumExecFee     float64         `json:"cum_exec_fee"`
	ReduceOnly     bool            `json:"reduce_only"`
	CloseOnTrigger bool            `json:"close_on_trigger"`
	OrderLinkID    string          `json:"order_link_id"`
	CreatedTime    string          `json:"created_time"`
	UpdatedTime    string          `json:"updated_time"`
	TakeProfit     float64         `json:"take_profit"`
	StopLoss       float64         `json:"stop_loss"`
	TpTriggerBy    TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy    TriggerByFuture `json:"sl_trigger_by"`
}

// ListLinearOrderParam :
type ListLinearOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	OrderID     *string      `url:"order_id,omitempty"`
	OrderLinkID *string      `url:"order_link_id,omitempty"`
	Order       *Order       `url:"order,omitempty"`
	Page        *int         `url:"page,omitempty"`
	Limit       *int         `url:"limit,omitempty"`
	OrderStatus *OrderStatus `url:"order_status,omitempty"`
}

// ListLinearOrder :
func (s *FutureUSDTPerpetualService) ListLinearOrder(param ListLinearOrderParam) (*ListLinearOrderResponse, error) {
	var res ListLinearOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/private/linear/order/list", queryString, &res); err != nil {
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
	UserID              int          `json:"user_id"`
	Symbol              SymbolFuture `json:"symbol"`
	Side                Side         `json:"side"`
	Size                float64      `json:"size"`
	PositionValue       float64      `json:"position_value"`
	EntryPrice          float64      `json:"entry_price"`
	LiqPrice            float64      `json:"liq_price"`
	BustPrice           float64      `json:"bust_price"`
	Leverage            float64      `json:"leverage"`
	AutoAddMargin       float64      `json:"auto_add_margin"`
	IsIsolated          bool         `json:"is_isolated"`
	PositionMargin      float64      `json:"position_margin"`
	OccClosingFee       float64      `json:"occ_closing_fee"`
	RealisedPnl         float64      `json:"realised_pnl"`
	CumRealisedPnl      float64      `json:"cum_realised_pnl"`
	FreeQty             float64      `json:"free_qty"`
	TpSlMode            TpSlMode     `json:"tp_sl_mode"`
	DeleverageIndicator int          `json:"deleverage_indicator"`
	UnrealisedPnl       float64      `json:"unrealised_pnl"`
	RiskID              int          `json:"risk_id"`
}

// ListLinearPosition :
func (s *FutureUSDTPerpetualService) ListLinearPosition(symbol SymbolFuture) (*ListLinearPositionResponse, error) {
	var res ListLinearPositionResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPrivately("/private/linear/position/list", query, &res); err != nil {
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
func (s *FutureUSDTPerpetualService) ListLinearPositions() (*ListLinearPositionsResponse, error) {
	var res ListLinearPositionsResponse

	if err := s.client.getPrivately("/private/linear/position/list", nil, &res); err != nil {
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
	Symbol SymbolFuture `json:"symbol"`

	OrderID     *string `json:"order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelLinearOrder :
func (s *FutureUSDTPerpetualService) CancelLinearOrder(param CancelLinearOrderParam) (*CancelLinearOrderResponse, error) {
	var res CancelLinearOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelLinearOrderParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/order/cancel", body, &res); err != nil {
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
	Symbol       SymbolFuture `json:"symbol"`
	BuyLeverage  float64      `json:"buy_leverage"`
	SellLeverage float64      `json:"sell_leverage"`
}

// SaveLinearLeverage :
func (s *FutureUSDTPerpetualService) SaveLinearLeverage(param SaveLinearLeverageParam) (*SaveLinearLeverageResponse, error) {
	var res SaveLinearLeverageResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for SaveLinearLeverageParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/position/set-leverage", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// LinearTradingStopResponse :
type LinearTradingStopResponse struct {
	CommonResponse `json:",inline"`
}

// LinearTradingStopParam :
type LinearTradingStopParam struct {
	Symbol SymbolFuture `json:"symbol"`
	Side   Side         `json:"side"`

	TakeProfit   *float64         `json:"take_profit,omitempty"`
	StopLoss     *float64         `json:"stop_loss,omitempty"`
	TrailingStop *float64         `json:"trailing_stop,omitempty"`
	TpTriggerBy  *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy  *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	SlSize       *float64         `json:"sl_size,omitempty"`
	TpSize       *float64         `json:"tp_size,omitempty"`
	PositionIdx  *int             `json:"position_idx,omitempty"`
}

// LinearTradingStop :
func (s *FutureUSDTPerpetualService) LinearTradingStop(param LinearTradingStopParam) (*LinearTradingStopResponse, error) {
	var res LinearTradingStopResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for LinearTradingStopParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/position/trading-stop", body, &res); err != nil {
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
	OrderID          string       `json:"order_id"`
	OrderLinkID      string       `json:"order_link_id"`
	Side             Side         `json:"side"`
	Symbol           SymbolFuture `json:"symbol"`
	OrderPrice       float64      `json:"order_price"`
	OrderQty         float64      `json:"order_qty"`
	OrderType        OrderType    `json:"order_type"`
	FeeRate          float64      `json:"fee_rate"`
	ExecPrice        float64      `json:"exec_price"`
	ExecType         ExecType     `json:"exec_type"`
	ExecQty          float64      `json:"exec_qty"`
	ExecFee          float64      `json:"exec_fee"`
	ExecValue        float64      `json:"exec_value"`
	LeavesQty        float64      `json:"leaves_qty"`
	ClosedSize       float64      `json:"closed_size"`
	LastLiquidityInd string       `json:"last_liquidity_ind"`
	TradeTimeMs      float64      `json:"trade_time_ms"`
}

// LinearExecutionListParam :
type LinearExecutionListParam struct {
	Symbol SymbolFuture `url:"symbol"`

	StartTime *int      `url:"start_time,omitempty"`
	EndTime   *int      `url:"end_time,omitempty"`
	ExecType  *ExecType `url:"exec_type,omitempty"`
	Page      *int      `url:"page,omitempty"`
	Limit     *int      `url:"limit,omitempty"`
}

// LinearExecutionList :
func (s *FutureUSDTPerpetualService) LinearExecutionList(param LinearExecutionListParam) (*LinearExecutionListResponse, error) {
	var res LinearExecutionListResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/private/linear/trade/execution/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// LinearCancelAllParam : Parameters to be supplied to cancel all endpoint
type LinearCancelAllParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// LinearCancelAllResponse : Response from cancel all endpoint
type LinearCancelAllResponse struct {
	CommonResponse `json:",inline"`
	Result         LinearCancelAllResult `json:"result"`
}

type LinearCancelAllResult []string

// LinearCancelAllOrder : Cancel all active orders that are unfilled or partially filled. Fully filled orders cannot be cancelled.
func (s *FutureUSDTPerpetualService) LinearCancelAllOrder(param LinearCancelAllParam) (*LinearCancelAllResponse, error) {
	var res LinearCancelAllResponse

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal for LinearCancelAllParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/order/cancel-all", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// ReplaceLinearOrderResponse :
type ReplaceLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         ReplaceLinearOrderResult `json:"result"`
}

// ReplaceLinearOrderResult :
type ReplaceLinearOrderResult struct {
	OrderID string `json:"order_id"`
}

// ReplaceLinearOrderParam :
type ReplaceLinearOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`

	OrderID     *string          `json:"order_id,omitempty"`
	OrderLinkID *string          `json:"order_link_id,omitempty"`
	NewQuantity *float64         `json:"p_r_qty,omitempty"`
	NewPrice    *float64         `json:"p_r_price,omitempty"`
	TakeProfit  *float64         `json:"take_profit,omitempty"`
	StopLoss    *float64         `json:"stop_loss,omitempty"`
	TpTriggerBy *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy *TriggerByFuture `json:"sl_trigger_by,omitempty"`
}

// ReplaceLinearOrder :
func (s *FutureUSDTPerpetualService) ReplaceLinearOrder(param ReplaceLinearOrderParam) (*ReplaceLinearOrderResponse, error) {
	var res ReplaceLinearOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for ReplaceLinearOrderResult: %w", err)
	}

	if err := s.client.postJSON("/private/linear/order/replace", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryLinearOrderResponse :
type QueryLinearOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []QueryLinearOrderResult `json:"result"`
}

// QueryLinearOrderResult :
type QueryLinearOrderResult struct {
	OrderID        string          `json:"order_id"`
	UserID         int             `json:"user_id"`
	Symbol         SymbolFuture    `json:"symbol"`
	Side           Side            `json:"side"`
	OrderType      OrderType       `json:"order_type"`
	Price          float64         `json:"price"`
	Qty            float64         `json:"qty"`
	TimeInForce    TimeInForce     `json:"time_in_force"`
	OrderStatus    OrderStatus     `json:"order_status"`
	LastExecPrice  float64         `json:"last_exec_price"`
	CumExecQty     float64         `json:"cum_exec_qty"`
	CumExecValue   float64         `json:"cum_exec_value"`
	CumExecFee     float64         `json:"cum_exec_fee"`
	ReduceOnly     bool            `json:"reduce_only"`
	CloseOnTrigger bool            `json:"close_on_trigger"`
	OrderLinkID    string          `json:"order_link_id"`
	CreatedTime    string          `json:"created_time"`
	UpdatedTime    string          `json:"updated_time"`
	TakeProfit     float64         `json:"take_profit"`
	StopLoss       float64         `json:"stop_loss"`
	TpTriggerBy    TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy    TriggerByFuture `json:"sl_trigger_by"`
}

// QueryLinearOrderParam :
type QueryLinearOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	OrderID     *string `url:"order_id,omitempty"`
	OrderLinkID *string `url:"order_link_id,omitempty"`
}

// QueryLinearOrder :
func (s *FutureUSDTPerpetualService) QueryLinearOrder(param QueryLinearOrderParam) (*QueryLinearOrderResponse, error) {
	var res QueryLinearOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/private/linear/order/search", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateLinearStopOrderResponse :
type CreateLinearStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateLinearStopOrderResult `json:"result"`
}

// CreateLinearStopOrderResult :
type CreateLinearStopOrderResult struct {
	StopOrderID    string          `json:"stop_order_id"`
	UserID         int             `json:"user_id"`
	Symbol         SymbolFuture    `json:"symbol"`
	Side           Side            `json:"side"`
	OrderType      OrderType       `json:"order_type"`
	Price          float64         `json:"price"`
	Qty            float64         `json:"qty"`
	TimeInForce    TimeInForce     `json:"time_in_force"`
	OrderStatus    OrderStatus     `json:"order_status"`
	TriggerPrice   float64         `json:"trigger_price"`
	OrderLinkID    string          `json:"order_link_id"`
	CreatedTime    string          `json:"created_time"`
	UpdatedTime    string          `json:"updated_time"`
	BasePrice      string          `json:"base_price"`
	TriggerBy      TriggerByFuture `json:"trigger_by"`
	TpTriggerBy    TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy    TriggerByFuture `json:"sl_trigger_by"`
	TakeProfit     float64         `json:"take_profit"`
	StopLoss       float64         `json:"stop_loss"`
	ReduceOnly     bool            `json:"reduce_only"`
	CloseOnTrigger bool            `json:"close_on_trigger"`
	PositionIdx    int             `json:"position_idx"`
}

// CreateLinearStopOrderParam :
type CreateLinearStopOrderParam struct {
	Side           Side            `json:"side"`
	Symbol         SymbolFuture    `json:"symbol"`
	OrderType      OrderType       `json:"order_type"`
	Qty            float64         `json:"qty"`
	BasePrice      float64         `json:"base_price"`
	StopPx         float64         `json:"stop_px"`
	TimeInForce    TimeInForce     `json:"time_in_force"`
	TriggerBy      TriggerByFuture `json:"trigger_by"`
	ReduceOnly     bool            `json:"reduce_only"`
	CloseOnTrigger bool            `json:"close_on_trigger"`

	Price       *float64         `json:"price,omitempty"`
	OrderLinkID *string          `json:"order_link_id,omitempty"`
	TakeProfit  *float64         `json:"take_profit,omitempty"`
	StopLoss    *float64         `json:"stop_loss,omitempty"`
	TpTriggerBy *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	PositionIdx *int             `json:"position_idx,omitempty"`
}

// CreateLinearStopOrder :
func (s *FutureUSDTPerpetualService) CreateLinearStopOrder(param CreateLinearStopOrderParam) (*CreateLinearStopOrderResponse, error) {
	var res CreateLinearStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateLinearStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/stop-order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListLinearStopOrderResponse :
type ListLinearStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         ListLinearStopOrderResult `json:"result"`
}

// ListLinearStopOrderResult :
type ListLinearStopOrderResult struct {
	CurrentPage int                                `json:"current_page"`
	LastPage    int                                `json:"last_page"`
	Content     []ListLinearStopOrderResultContent `json:"data"`
}

// ListLinearStopOrderResultContent :
type ListLinearStopOrderResultContent struct {
	StopOrderID    string          `json:"stop_order_id"`
	UserID         int             `json:"user_id"`
	Symbol         SymbolFuture    `json:"symbol"`
	Side           Side            `json:"side"`
	OrderType      OrderType       `json:"order_type"`
	Price          float64         `json:"price"`
	Qty            float64         `json:"qty"`
	TimeInForce    TimeInForce     `json:"time_in_force"`
	OrderStatus    OrderStatus     `json:"order_status"`
	TriggerPrice   float64         `json:"trigger_price"`
	OrderLinkID    string          `json:"order_link_id"`
	CreatedTime    string          `json:"created_time"`
	UpdatedTime    string          `json:"updated_time"`
	TakeProfit     float64         `json:"take_profit"`
	StopLoss       float64         `json:"stop_loss"`
	TriggerBy      TriggerByFuture `json:"trigger_by"`
	BasePrice      string          `json:"base_price"`
	TpTriggerBy    TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy    TriggerByFuture `json:"sl_trigger_by"`
	ReduceOnly     bool            `json:"reduce_only"`
	CloseOnTrigger bool            `json:"close_on_trigger"`
}

// ListLinearStopOrderParam :
type ListLinearStopOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	StopOrderID     *string      `url:"stop_order_id,omitempty"`
	OrderLinkID     *string      `url:"order_link_id,omitempty"`
	StopOrderStatus *OrderStatus `url:"stop_order_status,omitempty"`
	Order           *Order       `url:"order,omitempty"`
	Page            *int         `url:"page,omitempty"`
	Limit           *int         `url:"limit,omitempty"`
}

// ListLinearStopOrder :
func (s *FutureUSDTPerpetualService) ListLinearStopOrder(param ListLinearStopOrderParam) (*ListLinearStopOrderResponse, error) {
	var res ListLinearStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/private/linear/stop-order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelLinearStopOrderResponse :
type CancelLinearStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelLinearStopOrderResult `json:"result"`
}

// CancelLinearStopOrderResult :
type CancelLinearStopOrderResult struct {
	StopOrderID string `json:"stop_order_id"`
}

// CancelLinearStopOrderParam :
type CancelLinearStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`

	StopOrderID *string `json:"stop_order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelLinearStopOrder :
func (s *FutureUSDTPerpetualService) CancelLinearStopOrder(param CancelLinearStopOrderParam) (*CancelLinearStopOrderResponse, error) {
	var res CancelLinearStopOrderResponse

	if param.StopOrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either StopOrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelLinearStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/stop-order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllLinearStopOrderResponse :
type CancelAllLinearStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelAllLinearStopOrderResult `json:"result"`
}

// CancelAllLinearStopOrderResult :
type CancelAllLinearStopOrderResult []string

// CancelAllLinearStopOrderParam :
type CancelAllLinearStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllLinearStopOrder :
func (s *FutureUSDTPerpetualService) CancelAllLinearStopOrder(param CancelAllLinearStopOrderParam) (*CancelAllLinearStopOrderResponse, error) {
	var res CancelAllLinearStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllLinearStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/private/linear/stop-order/cancel-all", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryLinearStopOrderResponse :
type QueryLinearStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []QueryLinearStopOrderResult `json:"result"`
}

// QueryLinearStopOrderResult :
type QueryLinearStopOrderResult struct {
	StopOrderID    string          `json:"stop_order_id"`
	UserID         int             `json:"user_id"`
	Symbol         SymbolFuture    `json:"symbol"`
	Side           Side            `json:"side"`
	OrderType      OrderType       `json:"order_type"`
	Price          float64         `json:"price"`
	Qty            float64         `json:"qty"`
	TimeInForce    TimeInForce     `json:"time_in_force"`
	OrderStatus    OrderStatus     `json:"order_status"`
	TriggerPrice   float64         `json:"trigger_price"`
	BasePrice      string          `json:"base_price"`
	OrderLinkID    string          `json:"order_link_id"`
	CreatedTime    string          `json:"created_time"`
	UpdatedTime    string          `json:"updated_time"`
	TakeProfit     float64         `json:"take_profit"`
	StopLoss       float64         `json:"stop_loss"`
	TpTriggerBy    TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy    TriggerByFuture `json:"sl_trigger_by"`
	TriggerBy      TriggerByFuture `json:"trigger_by"`
	ReduceOnly     bool            `json:"reduce_only"`
	CloseOnTrigger bool            `json:"close_on_trigger"`
}

// QueryLinearStopOrderParam :
type QueryLinearStopOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	StopOrderID *string `url:"stop_order_id,omitempty"`
	OrderLinkID *string `url:"order_link_id,omitempty"`
}

// QueryLinearStopOrder :
func (s *FutureUSDTPerpetualService) QueryLinearStopOrder(param QueryLinearStopOrderParam) (*QueryLinearStopOrderResponse, error) {
	var res QueryLinearStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/private/linear/stop-order/search", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
