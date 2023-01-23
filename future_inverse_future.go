package bybit

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// FutureInverseFutureServiceI :
type FutureInverseFutureServiceI interface {
	// Market Data Endpoints
	OrderBook(SymbolFuture) (*OrderBookResponse, error)
	ListKline(ListKlineParam) (*ListKlineResponse, error)
	Tickers(SymbolFuture) (*TickersResponse, error)
	TradingRecords(TradingRecordsParam) (*TradingRecordsResponse, error)
	Symbols() (*SymbolsResponse, error)
	MarkPriceKline(MarkPriceKlineParam) (*MarkPriceKlineResponse, error)
	IndexPriceKline(IndexPriceKlineParam) (*IndexPriceKlineResponse, error)
	OpenInterest(OpenInterestParam) (*OpenInterestResponse, error)
	BigDeal(BigDealParam) (*BigDealResponse, error)
	AccountRatio(AccountRatioParam) (*AccountRatioResponse, error)

	// Account Data Endpoints
	CreateFuturesOrder(CreateFuturesOrderParam) (*CreateFuturesOrderResponse, error)
	ListFuturesOrder(ListFuturesOrderParam) (*ListFuturesOrderResponse, error)
	CancelFuturesOrder(CancelFuturesOrderParam) (*CancelFuturesOrderResponse, error)
	CancelAllFuturesOrder(CancelAllFuturesOrderParam) (*CancelAllFuturesOrderResponse, error)
	QueryFuturesOrder(QueryFuturesOrderParam) (*QueryFuturesOrderResponse, error)
	CreateFuturesStopOrder(CreateFuturesStopOrderParam) (*CreateFuturesStopOrderResponse, error)
	ListFuturesStopOrder(ListFuturesStopOrderParam) (*ListFuturesStopOrderResponse, error)
	CancelFuturesStopOrder(CancelFuturesStopOrderParam) (*CancelFuturesStopOrderResponse, error)
	CancelAllFuturesStopOrder(CancelAllFuturesStopOrderParam) (*CancelAllFuturesStopOrderResponse, error)
	QueryFuturesStopOrder(QueryFuturesStopOrderParam) (*QueryFuturesStopOrderResponse, error)
	ListFuturesPositions(SymbolFuture) (*ListFuturesPositionsResponse, error)
	FuturesTradingStop(FuturesTradingStopParam) (*FuturesTradingStopResponse, error)
	FuturesSaveLeverage(FuturesSaveLeverageParam) (*FuturesSaveLeverageResponse, error)
	APIKeyInfo() (*APIKeyInfoResponse, error)

	// Wallet Data Endpoints
	Balance(Coin) (*BalanceResponse, error)
}

// FutureInverseFutureService :
type FutureInverseFutureService struct {
	client *Client

	*FutureCommonService
}

// CreateFuturesOrderResponse :
type CreateFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateFuturesOrderResult `json:"result"`
}

// CreateFuturesOrderResult :
type CreateFuturesOrderResult struct {
	UserID        int             `json:"user_id"`
	OrderID       string          `json:"order_id"`
	Symbol        SymbolFuture    `json:"symbol"`
	Side          Side            `json:"side"`
	OrderType     OrderType       `json:"order_type"`
	Price         float64         `json:"price"`
	Qty           float64         `json:"qty"`
	TimeInForce   TimeInForce     `json:"time_in_force"`
	OrderStatus   OrderStatus     `json:"order_status"`
	LastExecTime  float64         `json:"last_exec_time"`
	LastExecPrice float64         `json:"last_exec_price"`
	LeavesQty     float64         `json:"leaves_qty"`
	CumExecQty    float64         `json:"cum_exec_qty"`
	CumExecValue  float64         `json:"cum_exec_value"`
	CumExecFee    float64         `json:"cum_exec_fee"`
	RejectReason  string          `json:"reject_reason"`
	OrderLinkID   string          `json:"order_link_id"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	TakeProfit    string          `json:"take_profit"`
	StopLoss      string          `json:"stop_loss"`
	TpTriggerBy   TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy   TriggerByFuture `json:"sl_trigger_by"`
}

// CreateFuturesOrderParam :
type CreateFuturesOrderParam struct {
	Side        Side         `json:"side"`
	Symbol      SymbolFuture `json:"symbol"`
	OrderType   OrderType    `json:"order_type"`
	Qty         int          `json:"qty"`
	TimeInForce TimeInForce  `json:"time_in_force"`

	Price          *float64         `json:"price,omitempty"`
	PositionIdx    *int             `json:"position_idx,omitempty"`
	ReduceOnly     *bool            `json:"reduce_only,omitempty"`
	CloseOnTrigger *bool            `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string          `json:"order_link_id,omitempty"`
	TakeProfit     *float64         `json:"take_profit,omitempty"`
	StopLoss       *float64         `json:"stop_loss,omitempty"`
	TpTriggerBy    *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy    *TriggerByFuture `json:"sl_trigger_by,omitempty"`
}

// CreateFuturesOrder :
func (s *FutureInverseFutureService) CreateFuturesOrder(param CreateFuturesOrderParam) (*CreateFuturesOrderResponse, error) {
	var res CreateFuturesOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateFuturesOrderParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListFuturesOrderResponse :
type ListFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         ListFuturesOrderResult `json:"result"`
}

// ListFuturesOrderResult :
type ListFuturesOrderResult struct {
	ListFuturesOrders []ListFuturesOrder `json:"data"`
}

// ListFuturesOrder :
type ListFuturesOrder struct {
	UserID       int             `json:"user_id"`
	PositionIdx  int             `json:"position_idx"`
	Symbol       SymbolFuture    `json:"symbol"`
	Side         Side            `json:"side"`
	OrderType    OrderType       `json:"order_type"`
	Price        string          `json:"price"`
	Qty          string          `json:"qty"`
	TimeInForce  TimeInForce     `json:"time_in_force"`
	OrderLinkID  string          `json:"order_link_id"`
	OrderID      string          `json:"order_id"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
	OrderStatus  OrderStatus     `json:"order_status"`
	LeavesQty    string          `json:"leaves_qty"`
	LeavesValue  string          `json:"leaves_value"`
	CumExecQty   string          `json:"cum_exec_qty"`
	CumExecValue string          `json:"cum_exec_value"`
	CumExecFee   string          `json:"cum_exec_fee"`
	RejectReason string          `json:"reject_reason"`
	TakeProfit   string          `json:"take_profit"`
	StopLoss     string          `json:"stop_loss"`
	TpTriggerBy  TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy  TriggerByFuture `json:"sl_trigger_by"`
	Cursor       string          `json:"cursor"`
}

// ListFuturesOrderParam :
type ListFuturesOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	OrderStatus *OrderStatus `url:"order_status,omitempty"`
	Direction   *Direction   `url:"direction,omitempty"`
	Limit       *int         `url:"limit,omitempty"`
	Cursor      *string      `url:"cursor,omitempty"`
}

// ListFuturesOrder :
func (s *FutureInverseFutureService) ListFuturesOrder(param ListFuturesOrderParam) (*ListFuturesOrderResponse, error) {
	var res ListFuturesOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/futures/private/order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelFuturesOrderResponse :
type CancelFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelFuturesOrderResult `json:"result"`
}

// CancelFuturesOrderResult :
type CancelFuturesOrderResult struct {
	UserID        int          `json:"user_id"`
	OrderID       string       `json:"order_id"`
	Symbol        SymbolFuture `json:"symbol"`
	Side          Side         `json:"side"`
	OrderType     OrderType    `json:"order_type"`
	Price         float64      `json:"price"`
	Qty           float64      `json:"qty"`
	TimeInForce   TimeInForce  `json:"time_in_force"`
	OrderStatus   OrderStatus  `json:"order_status"`
	LastExecTime  float64      `json:"last_exec_time"`
	LastExecPrice float64      `json:"last_exec_price"`
	LeavesQty     float64      `json:"leaves_qty"`
	CumExecQty    float64      `json:"cum_exec_qty"`
	CumExecValue  float64      `json:"cum_exec_value"`
	CumExecFee    float64      `json:"cum_exec_fee"`
	RejectReason  string       `json:"reject_reason"`
	OrderLinkID   string       `json:"order_link_id"`
	CreatedAt     string       `json:"created_at"`
	UpdatedAt     string       `json:"updated_at"`
}

// CancelFuturesOrderParam :
type CancelFuturesOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`

	OrderID     *string `json:"order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelFuturesOrder :
func (s *FutureInverseFutureService) CancelFuturesOrder(param CancelFuturesOrderParam) (*CancelFuturesOrderResponse, error) {
	var res CancelFuturesOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelFuturesOrderParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllFuturesOrderResponse :
type CancelAllFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []CancelAllFuturesOrderResult `json:"result"`
}

// CancelAllFuturesOrderResult :
type CancelAllFuturesOrderResult struct {
	ClOrdID     string       `json:"clOrdID"`
	UserID      int          `json:"user_id"`
	Symbol      SymbolFuture `json:"symbol"`
	Side        Side         `json:"side"`
	OrderType   OrderType    `json:"order_type"`
	Price       string       `json:"price"`
	Qty         float64      `json:"qty"`
	TimeInForce TimeInForce  `json:"time_in_force"`
	CreateType  string       `json:"create_type"`
	CancelType  string       `json:"cancel_type"`
	OrderStatus OrderStatus  `json:"order_status"`
	LeavesQty   float64      `json:"leaves_qty"`
	LeavesValue string       `json:"leaves_value"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	CrossStatus string       `json:"cross_status"`
	CrossSeq    int          `json:"cross_seq"`
}

// CancelAllFuturesOrderParam :
type CancelAllFuturesOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllFuturesOrder :
func (s *FutureInverseFutureService) CancelAllFuturesOrder(param CancelAllFuturesOrderParam) (*CancelAllFuturesOrderResponse, error) {
	var res CancelAllFuturesOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllFuturesOrderParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryFuturesOrderResponse :
type QueryFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         QueryFuturesOrderResult `json:"result"`
}

// QueryFuturesOrderResult :
type QueryFuturesOrderResult struct {
	UserID       int                    `json:"user_id"`
	PositionIdx  int                    `json:"position_idx"`
	Symbol       SymbolFuture           `json:"symbol"`
	Side         Side                   `json:"side"`
	OrderType    OrderType              `json:"order_type"`
	Price        string                 `json:"price"`
	Qty          float64                `json:"qty"`
	TimeInForce  TimeInForce            `json:"time_in_force"`
	OrderStatus  OrderStatus            `json:"order_status"`
	ExtFields    map[string]interface{} `json:"ext_fields"`
	LastExecTime string                 `json:"last_exec_time"`
	LeavesQty    int                    `json:"leaves_qty"`
	LeavesValue  string                 `json:"leaves_value"`
	CumExecQty   int                    `json:"cum_exec_qty"`
	CumExecValue string                 `json:"cum_exec_value"`
	CumExecFee   string                 `json:"cum_exec_fee"`
	RejectReason string                 `json:"reject_reason"`
	CancelType   string                 `json:"cancel_type"`
	OrderLinkID  string                 `json:"order_link_id"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	OrderID      string                 `json:"order_id"`
	TakeProfit   string                 `json:"take_profit"`
	StopLoss     string                 `json:"stop_loss"`
	TpTriggerBy  TriggerByFuture        `json:"tp_trigger_by"`
	SlTriggerBy  TriggerByFuture        `json:"sl_trigger_by"`
}

// QueryFuturesOrderParam :
type QueryFuturesOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	OrderID     *string `url:"order_id,omitempty"`
	OrderLinkID *string `url:"order_link_id,omitempty"`
}

// QueryFuturesOrder :
func (s *FutureInverseFutureService) QueryFuturesOrder(param QueryFuturesOrderParam) (*QueryFuturesOrderResponse, error) {
	var res QueryFuturesOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/futures/private/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateFuturesStopOrderResponse :
type CreateFuturesStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateFuturesStopOrderResult `json:"result"`
}

// CreateFuturesStopOrderResult :
type CreateFuturesStopOrderResult struct {
	UserID       int             `json:"user_id"`
	Symbol       SymbolFuture    `json:"symbol"`
	Side         Side            `json:"side"`
	OrderType    OrderType       `json:"order_type"`
	Price        string          `json:"price"`
	Qty          string          `json:"qty"`
	TimeInForce  TimeInForce     `json:"time_in_force"`
	Remark       string          `json:"remark"`
	LeavesQty    string          `json:"leaves_qty"`
	LeavesValue  string          `json:"leaves_value"`
	StopPx       string          `json:"stop_px"`
	RejectReason string          `json:"reject_reason"`
	StopOrderID  string          `json:"stop_order_id"`
	OrderLinkID  string          `json:"order_link_id"`
	TriggerBy    TriggerByFuture `json:"trigger_by"`
	BasePrice    string          `json:"base_price"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
	TpTriggerBy  TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy  TriggerByFuture `json:"sl_trigger_by"`
	TakeProfit   string          `json:"take_profit"`
	StopLoss     string          `json:"stop_loss"`
}

// CreateFuturesStopOrderParam :
type CreateFuturesStopOrderParam struct {
	Side        Side         `json:"side"`
	Symbol      SymbolFuture `json:"symbol"`
	OrderType   OrderType    `json:"order_type"`
	Qty         float64      `json:"qty"`
	BasePrice   float64      `json:"base_price"`
	StopPx      float64      `json:"stop_px"`
	TimeInForce TimeInForce  `json:"time_in_force"`

	PositionIdx    *int             `json:"position_idx,omitempty"`
	Price          *float64         `json:"price,omitempty"`
	TriggerBy      *TriggerByFuture `json:"trigger_by,omitempty"`
	CloseOnTrigger *bool            `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string          `json:"order_link_id,omitempty"`
	TakeProfit     *float64         `json:"take_profit,omitempty"`
	StopLoss       *float64         `json:"stop_loss,omitempty"`
	TpTriggerBy    *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy    *TriggerByFuture `json:"sl_trigger_by,omitempty"`
}

// CreateFuturesStopOrder :
func (s *FutureInverseFutureService) CreateFuturesStopOrder(param CreateFuturesStopOrderParam) (*CreateFuturesStopOrderResponse, error) {
	var res CreateFuturesStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateFuturesStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/stop-order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListFuturesStopOrderResponse :
type ListFuturesStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         ListFuturesStopOrderResult `json:"result"`
}

// ListFuturesStopOrderResult :
type ListFuturesStopOrderResult struct {
	ListFuturesStopOrders []ListFuturesStopOrder `json:"data"`
}

// ListFuturesStopOrder :
type ListFuturesStopOrder struct {
	UserID          int                 `json:"user_id"`
	PositionIdx     int                 `json:"position_idx"`
	StopOrderStatus OrderStatus         `json:"stop_order_status"`
	Symbol          SymbolFuture        `json:"symbol"`
	Side            Side                `json:"side"`
	OrderType       OrderType           `json:"order_type"`
	StopOrderType   StopOrderTypeFuture `json:"stop_order_type"`
	Price           string              `json:"price"`
	Qty             string              `json:"qty"`
	TimeInForce     TimeInForce         `json:"time_in_force"`
	BasePrice       string              `json:"base_price"`
	OrderLinkID     string              `json:"order_link_id"`
	CreatedAt       string              `json:"created_at"`
	UpdatedAt       string              `json:"updated_at"`
	StopPx          string              `json:"stop_px"`
	StopOrderID     string              `json:"stop_order_id"`
	TriggerBy       TriggerByFuture     `json:"trigger_by"`
	TakeProfit      string              `json:"take_profit"`
	StopLoss        string              `json:"stop_loss"`
	TpTriggerBy     TriggerByFuture     `json:"tp_trigger_by"`
	SlTriggerBy     TriggerByFuture     `json:"sl_trigger_by"`
	Cursor          string              `json:"cursor"`
}

// ListFuturesStopOrderParam :
type ListFuturesStopOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	StopOrderStatus *OrderStatus `url:"stop_order_status,omitempty"`
	Direction       *Direction   `url:"direction,omitempty"`
	Limit           *int         `url:"limit,omitempty"`
	Cursor          *string      `url:"cursor,omitempty"`
}

// ListFuturesStopOrder :
func (s *FutureInverseFutureService) ListFuturesStopOrder(param ListFuturesStopOrderParam) (*ListFuturesStopOrderResponse, error) {
	var res ListFuturesStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/futures/private/stop-order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelFuturesStopOrderResponse :
type CancelFuturesStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelFuturesStopOrderResult `json:"result"`
}

// CancelFuturesStopOrderResult :
type CancelFuturesStopOrderResult struct {
	StopOrderID string `json:"stop_order_id"`
}

// CancelFuturesStopOrderParam :
type CancelFuturesStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`

	StopOrderID *string `json:"stop_order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelFuturesStopOrder :
func (s *FutureInverseFutureService) CancelFuturesStopOrder(param CancelFuturesStopOrderParam) (*CancelFuturesStopOrderResponse, error) {
	var res CancelFuturesStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelFuturesStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/stop-order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllFuturesStopOrderResponse :
type CancelAllFuturesStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []CancelAllFuturesStopOrderResult `json:"result"`
}

// CancelAllFuturesStopOrderResult :
type CancelAllFuturesStopOrderResult struct {
	ClOrdID           string              `json:"clOrdID"`
	UserID            int                 `json:"user_id"`
	Symbol            SymbolFuture        `json:"symbol"`
	Side              Side                `json:"side"`
	OrderType         OrderType           `json:"order_type"`
	Price             string              `json:"price"`
	Qty               float64             `json:"qty"`
	TimeInForce       TimeInForce         `json:"time_in_force"`
	CreateType        string              `json:"create_type"`
	CancelType        string              `json:"cancel_type"`
	OrderStatus       OrderStatus         `json:"order_status"`
	LeavesQty         float64             `json:"leaves_qty"`
	LeavesValue       string              `json:"leaves_value"`
	CreatedAt         string              `json:"created_at"`
	UpdatedAt         string              `json:"updated_at"`
	CrossStatus       string              `json:"cross_status"`
	CrossSeq          int                 `json:"cross_seq"`
	StopOrderType     StopOrderTypeFuture `json:"stop_order_type"`
	TriggerBy         TriggerByFuture     `json:"trigger_by"`
	BasePrice         string              `json:"base_price"`
	ExpectedDirection string              `json:"expected_direction"`
}

// CancelAllFuturesStopOrderParam :
type CancelAllFuturesStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllFuturesStopOrder :
func (s *FutureInverseFutureService) CancelAllFuturesStopOrder(param CancelAllFuturesStopOrderParam) (*CancelAllFuturesStopOrderResponse, error) {
	var res CancelAllFuturesStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllFuturesStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/stop-order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryFuturesStopOrderResponse :
type QueryFuturesStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         QueryFuturesStopOrderResult `json:"result"`
}

// QueryFuturesStopOrderResult :
type QueryFuturesStopOrderResult struct {
	UserID       int                    `json:"user_id"`
	PositionIdx  int                    `json:"position_idx"`
	Symbol       SymbolFuture           `json:"symbol"`
	Side         Side                   `json:"side"`
	OrderType    OrderType              `json:"order_type"`
	Price        string                 `json:"price"`
	Qty          float64                `json:"qty"`
	StopPx       string                 `json:"stop_px"`
	BasePrice    string                 `json:"base_price"`
	TimeInForce  TimeInForce            `json:"time_in_force"`
	OrderStatus  OrderStatus            `json:"order_status"`
	ExtFields    map[string]interface{} `json:"ext_fields"`
	LeavesQty    int                    `json:"leaves_qty"`
	LeavesValue  string                 `json:"leaves_value"`
	CumExecQty   int                    `json:"cum_exec_qty"`
	CumExecValue string                 `json:"cum_exec_value"`
	CumExecFee   string                 `json:"cum_exec_fee"`
	RejectReason string                 `json:"reject_reason"`
	OrderLinkID  string                 `json:"order_link_id"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	OrderID      string                 `json:"order_id"`
	TriggerBy    TriggerByFuture        `json:"trigger_by"`
	TakeProfit   string                 `json:"take_profit"`
	StopLoss     string                 `json:"stop_loss"`
	TpTriggerBy  TriggerByFuture        `json:"tp_trigger_by"`
	SlTriggerBy  TriggerByFuture        `json:"sl_trigger_by"`
}

// QueryFuturesStopOrderParam :
type QueryFuturesStopOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	StopOrderID *string `url:"stop_order_id,omitempty"`
	OrderLinkID *string `url:"order_link_id,omitempty"`
}

// QueryFuturesStopOrder :
func (s *FutureInverseFutureService) QueryFuturesStopOrder(param QueryFuturesStopOrderParam) (*QueryFuturesStopOrderResponse, error) {
	var res QueryFuturesStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/futures/private/stop-order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListFuturesPositionsResponse :
type ListFuturesPositionsResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListFuturesPositionsResult `json:"result"`
}

// ListFuturesPositionsResult :
type ListFuturesPositionsResult struct {
	Data ListFuturesPositionsResultData `json:"data"`
}

// ListFuturesPositionsResultData :
type ListFuturesPositionsResultData struct {
	ID                  int          `json:"id"`
	PositionIdx         int          `json:"position_idx"`
	Mode                int          `json:"mode"`
	UserID              int          `json:"user_id"`
	RiskID              int          `json:"risk_id"`
	Symbol              SymbolFuture `json:"symbol"`
	Side                Side         `json:"side"`
	Size                float64      `json:"size"`
	PositionValue       string       `json:"position_value"`
	EntryPrice          string       `json:"entry_price"`
	IsIsolated          bool         `json:"is_isolated"`
	AutoAddMargin       float64      `json:"auto_add_margin"`
	Leverage            string       `json:"leverage"`
	EffectiveLeverage   string       `json:"effective_leverage"`
	PositionMargin      string       `json:"position_margin"`
	LiqPrice            string       `json:"liq_price"`
	BustPrice           string       `json:"bust_price"`
	OccClosingFee       string       `json:"occ_closing_fee"`
	OccFundingFee       string       `json:"occ_funding_fee"`
	TakeProfit          string       `json:"take_profit"`
	StopLoss            string       `json:"stop_loss"`
	TrailingStop        string       `json:"trailing_stop"`
	PositionStatus      string       `json:"position_status"`
	DeleverageIndicator int          `json:"deleverage_indicator"`
	OcCalcData          string       `json:"oc_calc_data"`
	OrderMargin         string       `json:"order_margin"`
	WalletBalance       string       `json:"wallet_balance"`
	RealisedPnl         string       `json:"realised_pnl"`
	UnrealisedPnl       float64      `json:"unrealised_pnl"`
	CumRealisedPnl      string       `json:"cum_realised_pnl"`
	CrossSeq            float64      `json:"cross_seq"`
	PositionSeq         float64      `json:"position_seq"`
	CreatedAt           string       `json:"created_at"`
	UpdatedAt           string       `json:"updated_at"`
	TpSlMode            TpSlMode     `json:"tp_sl_mode"`
}

// ListFuturesPositions :
func (s *FutureInverseFutureService) ListFuturesPositions(symbol SymbolFuture) (*ListFuturesPositionsResponse, error) {
	var res ListFuturesPositionsResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPrivately("/futures/private/position/list", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FuturesTradingStopResponse :
type FuturesTradingStopResponse struct {
	CommonResponse `json:",inline"`
	Result         FuturesTradingStopResult `json:"result"`
}

// FuturesTradingStopResult :
type FuturesTradingStopResult struct {
	ID                  int                    `json:"id"`
	UserID              int                    `json:"user_id"`
	Symbol              SymbolFuture           `json:"symbol"`
	Side                Side                   `json:"side"`
	Size                float64                `json:"size"`
	PositionValue       float64                `json:"position_value"`
	EntryPrice          float64                `json:"entry_price"`
	RiskID              int                    `json:"risk_id"`
	AutoAddMargin       float64                `json:"auto_add_margin"`
	Leverage            float64                `json:"leverage"`
	PositionMargin      float64                `json:"position_margin"`
	LiqPrice            float64                `json:"liq_price"`
	BustPrice           float64                `json:"bust_price"`
	OccClosingFee       float64                `json:"occ_closing_fee"`
	OccFundingFee       float64                `json:"occ_funding_fee"`
	TakeProfit          float64                `json:"take_profit"`
	StopLoss            float64                `json:"stop_loss"`
	TrailingStop        float64                `json:"trailing_stop"`
	PositionStatus      string                 `json:"position_status"`
	DeleverageIndicator int                    `json:"deleverage_indicator"`
	OcCalcData          string                 `json:"oc_calc_data"`
	OrderMargin         float64                `json:"order_margin"`
	WalletBalance       float64                `json:"wallet_balance"`
	RealisedPnl         float64                `json:"realised_pnl"`
	CumRealisedPnl      float64                `json:"cum_realised_pnl"`
	CumCommission       float64                `json:"cum_commission"`
	CrossSeq            float64                `json:"cross_seq"`
	PositionSeq         float64                `json:"position_seq"`
	CreatedAt           string                 `json:"created_at"`
	UpdatedAt           string                 `json:"updated_at"`
	ExtFields           map[string]interface{} `json:"ext_fields"`
}

// FuturesTradingStopParam :
type FuturesTradingStopParam struct {
	Symbol SymbolFuture `json:"symbol"`

	PositionIdx       *int             `json:"position_idx,omitempty"`
	TakeProfit        *float64         `json:"take_profit,omitempty"`
	StopLoss          *float64         `json:"stop_loss,omitempty"`
	TrailingStop      *float64         `json:"trailing_stop,omitempty"`
	TpTriggerBy       *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy       *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	NewTrailingActive *float64         `json:"new_trailing_active,omitempty"`
	SlSize            *float64         `json:"sl_size,omitempty"`
	TpSize            *float64         `json:"tp_size,omitempty"`
}

// FuturesTradingStop :
func (s *FutureInverseFutureService) FuturesTradingStop(param FuturesTradingStopParam) (*FuturesTradingStopResponse, error) {
	var res FuturesTradingStopResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for FuturesTradingStopParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/position/trading-stop", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FuturesSaveLeverageResponse :
type FuturesSaveLeverageResponse struct {
	CommonResponse `json:",inline"`
	Result         float64 `json:"result"`
}

// FuturesSaveLeverageParam :
type FuturesSaveLeverageParam struct {
	Symbol       SymbolFuture `json:"symbol"`
	BuyLeverage  float64      `json:"buy_leverage"`
	SellLeverage float64      `json:"sell_leverage"`
}

// FuturesSaveLeverage :
func (s *FutureInverseFutureService) FuturesSaveLeverage(param FuturesSaveLeverageParam) (*FuturesSaveLeverageResponse, error) {
	var res FuturesSaveLeverageResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for FuturesSaveLeverageParam: %w", err)
	}

	if err := s.client.postJSON("/futures/private/position/leverage/save", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
