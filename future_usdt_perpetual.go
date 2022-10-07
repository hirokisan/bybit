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
	OrderBook(SymbolInverse) (*OrderBookResponse, error)
	Tickers(SymbolInverse) (*TickersResponse, error)
	Symbols() (*SymbolsResponse, error)
	OpenInterest(OpenInterestParam) (*OpenInterestResponse, error)
	BigDeal(BigDealParam) (*BigDealResponse, error)
	AccountRatio(AccountRatioParam) (*AccountRatioResponse, error)

	// Account Data Endpoints
	CreateLinearOrder(CreateLinearOrderParam) (*CreateLinearOrderResponse, error)
	ListLinearOrder(ListLinearOrderParam) (*ListLinearOrderResponse, error)
	CancelLinearOrder(CancelLinearOrderParam) (*CancelLinearOrderResponse, error)
	LinearCancelAllOrder(LinearCancelAllParam) (*LinearCancelAllResponse, error)
	ListLinearPosition(SymbolUSDT) (*ListLinearPositionResponse, error)
	ListLinearPositions() (*ListLinearPositionsResponse, error)
	SaveLinearLeverage(SaveLinearLeverageParam) (*SaveLinearLeverageResponse, error)
	LinearExecutionList(LinearExecutionListParam) (*LinearExecutionListResponse, error)

	// Wallet Data Endpoints
	Balance(Coin) (*BalanceResponse, error)
}

// FutureUSDTPerpetualService :
type FutureUSDTPerpetualService struct {
	client *Client

	*FutureCommonService
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

// ListLinearOrderParam :
type ListLinearOrderParam struct {
	Symbol SymbolUSDT `url:"symbol"`

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
func (s *FutureUSDTPerpetualService) ListLinearPosition(symbol SymbolUSDT) (*ListLinearPositionResponse, error) {
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
	Symbol SymbolUSDT `json:"symbol"`

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
	Symbol       SymbolUSDT `json:"symbol"`
	BuyLeverage  float64    `json:"buy_leverage"`
	SellLeverage float64    `json:"sell_leverage"`
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
	Symbol SymbolUSDT `url:"symbol"`

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
	Symbol SymbolUSDT `json:"symbol"`
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
