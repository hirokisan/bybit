package bybit

import (
	"encoding/json"
	"fmt"

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
