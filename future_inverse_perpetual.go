package bybit

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// FutureInversePerpetualServiceI :
type FutureInversePerpetualServiceI interface {
	// Market Data Endpoints
	OrderBook(SymbolFuture) (*OrderBookResponse, error)
	ListKline(ListKlineParam) (*ListKlineResponse, error)
	Tickers(SymbolFuture) (*TickersResponse, error)
	TradingRecords(TradingRecordsParam) (*TradingRecordsResponse, error)
	Symbols() (*SymbolsResponse, error)
	MarkPriceKline(MarkPriceKlineParam) (*MarkPriceKlineResponse, error)
	IndexPriceKline(IndexPriceKlineParam) (*IndexPriceKlineResponse, error)
	PremiumIndexKline(PremiumIndexKlineParam) (*PremiumIndexKlineResponse, error)
	OpenInterest(OpenInterestParam) (*OpenInterestResponse, error)
	BigDeal(BigDealParam) (*BigDealResponse, error)
	AccountRatio(AccountRatioParam) (*AccountRatioResponse, error)

	// Account Data Endpoints
	CreateOrder(CreateOrderParam) (*CreateOrderResponse, error)
	ListOrder(ListOrderParam) (*ListOrderResponse, error)
	CancelOrder(CancelOrderParam) (*CancelOrderResponse, error)
	CancelAllOrder(CancelAllOrderParam) (*CancelAllOrderResponse, error)
	QueryOrder(QueryOrderParam) (*QueryOrderResponse, error)
	CreateStopOrder(CreateStopOrderParam) (*CreateStopOrderResponse, error)
	ListStopOrder(ListStopOrderParam) (*ListStopOrderResponse, error)
	CancelStopOrder(CancelStopOrderParam) (*CancelStopOrderResponse, error)
	CancelAllStopOrder(CancelAllStopOrderParam) (*CancelAllStopOrderResponse, error)
	QueryStopOrder(QueryStopOrderParam) (*QueryStopOrderResponse, error)
	ListPosition(SymbolFuture) (*ListPositionResponse, error)
	ListPositions() (*ListPositionsResponse, error)
	TradingStop(TradingStopParam) (*TradingStopResponse, error)
	SaveLeverage(SaveLeverageParam) (*SaveLeverageResponse, error)
	APIKeyInfo() (*APIKeyInfoResponse, error)

	// Wallet Data Endpoints
	Balance(Coin) (*BalanceResponse, error)
}

// FutureInversePerpetualService :
type FutureInversePerpetualService struct {
	client *Client

	*FutureCommonService
}

// PremiumIndexKlineResponse :
type PremiumIndexKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []PremiumIndexKlineResult `json:"result"`
}

// PremiumIndexKlineResult :
type PremiumIndexKlineResult struct {
	Symbol   SymbolFuture `json:"symbol"`
	Period   Period       `json:"period"`
	OpenTime int          `json:"open_time"`
	Open     string       `json:"open"`
	High     string       `json:"high"`
	Low      string       `json:"low"`
	Close    string       `json:"close"`
}

// PremiumIndexKlineParam :
type PremiumIndexKlineParam struct {
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`

	Limit *int `url:"limit,omitempty"`
}

// PremiumIndexKline :
func (s *FutureInversePerpetualService) PremiumIndexKline(param PremiumIndexKlineParam) (*PremiumIndexKlineResponse, error) {
	var res PremiumIndexKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/premium-index-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
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

// CreateOrderParam :
type CreateOrderParam struct {
	Side        Side         `json:"side"`
	Symbol      SymbolFuture `json:"symbol"`
	OrderType   OrderType    `json:"order_type"`
	Qty         int          `json:"qty"`
	TimeInForce TimeInForce  `json:"time_in_force"`

	Price          *float64 `json:"price,omitempty"`
	TakeProfit     *float64 `json:"take_profit,omitempty"`
	StopLoss       *float64 `json:"stop_loss,omitempty"`
	ReduceOnly     *bool    `json:"reduce_only,omitempty"`
	CloseOnTrigger *bool    `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string  `json:"order_link_id,omitempty"`
}

// CreateOrder :
func (s *FutureInversePerpetualService) CreateOrder(param CreateOrderParam) (*CreateOrderResponse, error) {
	var res CreateOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateOrderParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListOrderResponse :
type ListOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         ListOrderResult `json:"result"`
}

// ListOrderResult :
type ListOrderResult struct {
	ListOrders []ListOrder `json:"data"`
}

// ListOrder :
type ListOrder struct {
	UserID       int             `json:"user_id"`
	Symbol       SymbolFuture    `json:"symbol"`
	Side         Side            `json:"side"`
	OrderType    OrderType       `json:"order_type"`
	Price        string          `json:"price"`
	Qty          string          `json:"qty"`
	TimeInForce  TimeInForce     `json:"time_in_force"`
	OrderStatus  OrderStatus     `json:"order_status"`
	LeavesQty    string          `json:"leaves_qty"`
	LeavesValue  string          `json:"leaves_value"`
	CumExecQty   string          `json:"cum_exec_qty"`
	CumExecValue string          `json:"cum_exec_value"`
	CumExecFee   string          `json:"cum_exec_fee"`
	RejectReason string          `json:"reject_reason"`
	OrderLinkID  string          `json:"order_link_id"`
	CreatedAt    string          `json:"created_at"`
	OrderID      string          `json:"order_id"`
	TakeProfit   string          `json:"take_profit"`
	StopLoss     string          `json:"stop_loss"`
	TpTriggerBy  TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy  TriggerByFuture `json:"sl_trigger_by"`
}

// ListOrderParam :
type ListOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	OrderStatus *OrderStatus `url:"order_status,omitempty"`
	Direction   *Direction   `url:"direction,omitempty"`
	Size        *int         `url:"size,omitempty"`
	Cursor      *string      `url:"cursor,omitempty"`
}

// ListOrder :
func (s *FutureInversePerpetualService) ListOrder(param ListOrderParam) (*ListOrderResponse, error) {
	var res ListOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/v2/private/order/list", queryString, &res); err != nil {
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
	ID                  int          `json:"id"`
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
}

// ListPosition :
func (s *FutureInversePerpetualService) ListPosition(symbol SymbolFuture) (*ListPositionResponse, error) {
	var res ListPositionResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPrivately("/v2/private/position/list", query, &res); err != nil {
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
func (s *FutureInversePerpetualService) ListPositions() (*ListPositionsResponse, error) {
	var res ListPositionsResponse

	if err := s.client.getPrivately("/v2/private/position/list", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// TradingStopResponse :
type TradingStopResponse struct {
	CommonResponse `json:",inline"`
	Result         TradingStopResult `json:"result"`
}

// TradingStopResult :
type TradingStopResult struct {
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

// TradingStopParam :
type TradingStopParam struct {
	Symbol SymbolFuture `json:"symbol"`

	TakeProfit        *float64         `json:"take_profit,omitempty"`
	StopLoss          *float64         `json:"stop_loss,omitempty"`
	TrailingStop      *float64         `json:"trailing_stop,omitempty"`
	TpTriggerBy       *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy       *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	NewTrailingActive *float64         `json:"new_trailing_active,omitempty"`
	SlSize            *float64         `json:"sl_size,omitempty"`
	TpSize            *float64         `json:"tp_size,omitempty"`
}

// TradingStop :
func (s *FutureInversePerpetualService) TradingStop(param TradingStopParam) (*TradingStopResponse, error) {
	var res TradingStopResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for TradingStopParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/position/trading-stop", body, &res); err != nil {
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

// CancelOrderParam :
type CancelOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`

	OrderID     *string `json:"order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelOrder :
func (s *FutureInversePerpetualService) CancelOrder(param CancelOrderParam) (*CancelOrderResponse, error) {
	var res CancelOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelOrderParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllOrderResponse :
type CancelAllOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []CancelAllOrderResult `json:"result"`
}

// CancelAllOrderResult :
type CancelAllOrderResult struct {
	ClOrdID     string       `json:"clOrdID"`
	OrderLinkID string       `json:"order_link_id"`
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

// CancelAllOrderParam :
type CancelAllOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllOrder :
func (s *FutureInversePerpetualService) CancelAllOrder(param CancelAllOrderParam) (*CancelAllOrderResponse, error) {
	var res CancelAllOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllOrderParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryOrderResponse :
type QueryOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []QueryOrderResult `json:"result"`
}

// QueryOrderResult :
type QueryOrderResult struct {
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

// QueryOrderParam :
type QueryOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	OrderID     *string `url:"order_id,omitempty"`
	OrderLinkID *string `url:"order_link_id,omitempty"`
}

// QueryOrder :
func (s *FutureInversePerpetualService) QueryOrder(param QueryOrderParam) (*QueryOrderResponse, error) {
	var res QueryOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/v2/private/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateStopOrderResponse :
type CreateStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateStopOrderResult `json:"result"`
}

// CreateStopOrderResult :
type CreateStopOrderResult struct {
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

// CreateStopOrderParam :
type CreateStopOrderParam struct {
	Side        Side         `json:"side"`
	Symbol      SymbolFuture `json:"symbol"`
	OrderType   OrderType    `json:"order_type"`
	Qty         int          `json:"qty"`
	BasePrice   float64      `json:"base_price"`
	StopPx      float64      `json:"stop_px"`
	TimeInForce TimeInForce  `json:"time_in_force"`

	Price          *float64         `json:"price,omitempty"`
	TriggerBy      *TriggerByFuture `json:"trigger_by,omitempty"`
	CloseOnTrigger *bool            `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string          `json:"order_link_id,omitempty"`
	TakeProfit     *float64         `json:"take_profit,omitempty"`
	StopLoss       *float64         `json:"stop_loss,omitempty"`
	TpTriggerBy    *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy    *TriggerByFuture `json:"sl_trigger_by,omitempty"`
}

// CreateStopOrder :
func (s *FutureInversePerpetualService) CreateStopOrder(param CreateStopOrderParam) (*CreateStopOrderResponse, error) {
	var res CreateStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/stop-order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListStopOrderResponse :
type ListStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         ListStopOrderResult `json:"result"`
}

// ListStopOrderResult :
type ListStopOrderResult struct {
	ListStopOrders []ListStopOrder `json:"data"`
}

// ListStopOrder :
type ListStopOrder struct {
	UserID          int                 `json:"user_id"`
	PositionIdx     int                 `json:"position_idx"`
	StopOrderStatus OrderStatus         `json:"stop_order_status"`
	Symbol          SymbolFuture        `json:"symbol"`
	Side            Side                `json:"side"`
	OrderType       OrderType           `json:"order_type"`
	Price           string              `json:"price"`
	Qty             string              `json:"qty"`
	TimeInForce     TimeInForce         `json:"time_in_force"`
	StopOrderType   StopOrderTypeFuture `json:"stop_order_type"`
	TriggerBy       TriggerByFuture     `json:"trigger_by"`
	BasePrice       string              `json:"base_price"`
	OrderLinkID     string              `json:"order_link_id"`
	CreatedAt       string              `json:"created_at"`
	UpdatedAt       string              `json:"updated_at"`
	StopPx          string              `json:"stop_px"`
	StopOrderID     string              `json:"stop_order_id"`
	TakeProfit      string              `json:"take_profit"`
	StopLoss        string              `json:"stop_loss"`
	TpTriggerBy     TriggerByFuture     `json:"tp_trigger_by"`
	SlTriggerBy     TriggerByFuture     `json:"sl_trigger_by"`
	Cursor          string              `json:"cursor"`
}

// ListStopOrderParam :
type ListStopOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	StopOrderStatus *OrderStatus `url:"stop_order_status,omitempty"`
	Direction       *Direction   `url:"direction,omitempty"`
	Limit           *int         `url:"limit,omitempty"`
	Cursor          *string      `url:"cursor,omitempty"`
}

// ListStopOrder :
func (s *FutureInversePerpetualService) ListStopOrder(param ListStopOrderParam) (*ListStopOrderResponse, error) {
	var res ListStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/v2/private/stop-order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelStopOrderResponse :
type CancelStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelStopOrderResult `json:"result"`
}

// CancelStopOrderResult :
type CancelStopOrderResult struct {
	StopOrderID string `json:"stop_order_id"`
}

// CancelStopOrderParam :
type CancelStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`

	StopOrderID *string `json:"stop_order_id,omitempty"`
	OrderLinkID *string `json:"order_link_id,omitempty"`
}

// CancelStopOrder :
func (s *FutureInversePerpetualService) CancelStopOrder(param CancelStopOrderParam) (*CancelStopOrderResponse, error) {
	var res CancelStopOrderResponse

	if param.StopOrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either StopOrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/stop-order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllStopOrderResponse :
type CancelAllStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []CancelAllStopOrderResult `json:"result"`
}

// CancelAllStopOrderResult :
type CancelAllStopOrderResult struct {
	ClOrdID           string              `json:"clOrdID"`
	OrderLinkID       string              `json:"order_link_id"`
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

// CancelAllStopOrderParam :
type CancelAllStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllStopOrder :
func (s *FutureInversePerpetualService) CancelAllStopOrder(param CancelAllStopOrderParam) (*CancelAllStopOrderResponse, error) {
	var res CancelAllStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllStopOrderParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/stop-order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryStopOrderResponse :
type QueryStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         []QueryStopOrderResult `json:"result"`
}

// QueryStopOrderResult :
type QueryStopOrderResult struct {
	UserID          int                    `json:"user_id"`
	PositionIdx     int                    `json:"position_idx"`
	Symbol          SymbolFuture           `json:"symbol"`
	Side            Side                   `json:"side"`
	OrderType       OrderType              `json:"order_type"`
	Price           string                 `json:"price"`
	Qty             float64                `json:"qty"`
	StopPx          string                 `json:"stop_px"`
	BasePrice       string                 `json:"base_price"`
	TimeInForce     TimeInForce            `json:"time_in_force"`
	StopOrderStatus OrderStatus            `json:"stop_order_status"`
	ExtFields       map[string]interface{} `json:"ext_fields"`
	LeavesQty       int                    `json:"leaves_qty"`
	LeavesValue     string                 `json:"leaves_value"`
	CumExecQty      int                    `json:"cum_exec_qty"`
	CumExecValue    string                 `json:"cum_exec_value"`
	CumExecFee      string                 `json:"cum_exec_fee"`
	RejectReason    string                 `json:"reject_reason"`
	OrderLinkID     string                 `json:"order_link_id"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
	OrderID         string                 `json:"order_id"`
	TriggerBy       TriggerByFuture        `json:"trigger_by"`
	TakeProfit      string                 `json:"take_profit"`
	StopLoss        string                 `json:"stop_loss"`
	TpTriggerBy     TriggerByFuture        `json:"tp_trigger_by"`
	SlTriggerBy     TriggerByFuture        `json:"sl_trigger_by"`
}

// QueryStopOrderParam :
type QueryStopOrderParam struct {
	Symbol SymbolFuture `url:"symbol"`

	StopOrderID *string `url:"stop_order_id,omitempty"`
	OrderLinkID *string `url:"order_link_id,omitempty"`
}

// QueryStopOrder :
func (s *FutureInversePerpetualService) QueryStopOrder(param QueryStopOrderParam) (*QueryStopOrderResponse, error) {
	var res QueryStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/v2/private/stop-order", queryString, &res); err != nil {
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
	Symbol   SymbolFuture `json:"symbol"`
	Leverage float64      `json:"leverage"`
}

// SaveLeverage :
func (s *FutureInversePerpetualService) SaveLeverage(param SaveLeverageParam) (*SaveLeverageResponse, error) {
	var res SaveLeverageResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelOrderParam: %w", err)
	}

	if err := s.client.postJSON("/v2/private/position/leverage/save", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
