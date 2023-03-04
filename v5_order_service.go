package bybit

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
)

// V5OrderServiceI :
type V5OrderServiceI interface {
	CreateOrder(V5CreateOrderParam) (*V5CreateOrderResponse, error)
	CancelOrder(V5CancelOrderParam) (*V5CancelOrderResponse, error)
	GetOpenOrders(V5GetOpenOrdersParam) (*V5GetOpenOrdersResponse, error)
}

// V5OrderService :
type V5OrderService struct {
	client *Client
}

// V5CreateOrderParam :
type V5CreateOrderParam struct {
	Category  CategoryV5 `json:"category"`
	Symbol    SymbolV5   `json:"symbol"`
	Side      Side       `json:"side"`
	OrderType OrderType  `json:"orderType"`
	Qty       string     `json:"qty"`

	IsLeverage            *IsLeverage       `json:"isLeverage,omitempty"`
	Price                 *string           `json:"price,omitempty"`
	TriggerDirection      *TriggerDirection `json:"triggerDirection,omitempty"`
	OrderFilter           *OrderFilter      `json:"orderFilter,omitempty"` // If not passed, Order by default
	TriggerPrice          *string           `json:"triggerPrice,omitempty"`
	TriggerBy             *TriggerBy        `json:"triggerBy,omitempty"`
	OrderIv               *string           `json:"orderIv,omitempty"`     // option only.
	TimeInForce           *TimeInForce      `json:"timeInForce,omitempty"` // If not passed, GTC is used by default
	PositionIdx           *PositionIdx      `json:"positionIdx,omitempty"` // Under hedge-mode, this param is required
	OrderLinkID           *string           `json:"orderLinkId,omitempty"`
	TakeProfit            *string           `json:"takeProfit,omitempty"`
	StopLoss              *string           `json:"stopLoss,omitempty"`
	TpTriggerBy           *TriggerBy        `json:"tpTriggerBy,omitempty"`
	SlTriggerBy           *TriggerBy        `json:"slTriggerBy,omitempty"`
	ReduceOnly            *bool             `json:"reduce_only,omitempty"`
	CloseOnTrigger        *bool             `json:"closeOnTrigger,omitempty"`
	MarketMakerProtection *bool             `json:"mmp,omitempty"` // option only
}

// V5CreateOrderResponse :
type V5CreateOrderResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5CreateOrderResult `json:"result"`
}

// V5CreateOrderResult :
type V5CreateOrderResult struct {
	OrderID     string `json:"orderId"`
	OrderLinkID string `json:"orderLinkId"`
}

// CreateOrder :
func (s *V5OrderService) CreateOrder(param V5CreateOrderParam) (*V5CreateOrderResponse, error) {
	var res V5CreateOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/order/create", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5CancelOrderParam :
type V5CancelOrderParam struct {
	Category CategoryV5 `json:"category"`
	Symbol   SymbolV5   `json:"symbol"`

	OrderID     *string      `json:"orderId,omitempty"`
	OrderLinkID *string      `json:"orderLinkId,omitempty"`
	OrderFilter *OrderFilter `json:"orderFilter,omitempty"` // If not passed, Order by default
}

// V5CancelOrderResponse :
type V5CancelOrderResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5CancelOrderResult `json:"result"`
}

// V5CancelOrderResult :
type V5CancelOrderResult struct {
	OrderID     string `json:"orderId"`
	OrderLinkID string `json:"orderLinkId"`
}

// CancelOrder :
func (s *V5OrderService) CancelOrder(param V5CancelOrderParam) (*V5CancelOrderResponse, error) {
	var res V5CancelOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/order/cancel", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5GetOpenOrdersParam :
type V5GetOpenOrdersParam struct {
	Category CategoryV5 `url:"category"`

	Symbol      SymbolV5     `url:"symbol,omitempty"`
	BaseCoin    *Coin        `url:"baseCoin,omitempty"`
	SettleCoin  *Coin        `url:"settleCoin,omitempty"`
	OrderID     *string      `url:"orderId,omitempty"`
	OrderLinkID *string      `url:"orderLinkId,omitempty"`
	OpenOnly    *int         `url:"openOnly,omitempty"`
	OrderFilter *OrderFilter `url:"orderFilter,omitempty"` // If not passed, Order by default
	Limit       *int         `url:"limit,omitempty"`
	Cursor      *string      `url:"cursor,omitempty"`
}

// V5GetOpenOrdersResponse :
type V5GetOpenOrdersResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetOpenOrdersResult `json:"result"`
}

// V5GetOpenOrdersResult :
type V5GetOpenOrdersResult struct {
	Category       CategoryV5 `json:"category"`
	NextPageCursor string     `json:"nextPageCursor"`
	List           []struct {
		Symbol             SymbolV5    `json:"symbol"`
		OrderType          OrderType   `json:"orderType"`
		OrderLinkID        string      `json:"orderLinkId"`
		OrderID            string      `json:"orderId"`
		CancelType         string      `json:"cancelType"`
		AvgPrice           string      `json:"avgPrice"`
		StopOrderType      string      `json:"stopOrderType"`
		LastPriceOnCreated string      `json:"lastPriceOnCreated"`
		OrderStatus        OrderStatus `json:"orderStatus"`
		TakeProfit         string      `json:"takeProfit"`
		CumExecValue       string      `json:"cumExecValue"`
		TriggerDirection   int         `json:"triggerDirection"`
		IsLeverage         string      `json:"isLeverage"`
		RejectReason       string      `json:"rejectReason"`
		Price              string      `json:"price"`
		OrderIv            string      `json:"orderIv"`
		CreatedTime        string      `json:"createdTime"`
		TpTriggerBy        string      `json:"tpTriggerBy"`
		PositionIdx        int         `json:"positionIdx"`
		TimeInForce        TimeInForce `json:"timeInForce"`
		LeavesValue        string      `json:"leavesValue"`
		UpdatedTime        string      `json:"updatedTime"`
		Side               Side        `json:"side"`
		TriggerPrice       string      `json:"triggerPrice"`
		CumExecFee         string      `json:"cumExecFee"`
		LeavesQty          string      `json:"leavesQty"`
		SlTriggerBy        string      `json:"slTriggerBy"`
		CloseOnTrigger     bool        `json:"closeOnTrigger"`
		CumExecQty         string      `json:"cumExecQty"`
		ReduceOnly         bool        `json:"reduceOnly"`
		Qty                string      `json:"qty"`
		StopLoss           string      `json:"stopLoss"`
		TriggerBy          TriggerBy   `json:"triggerBy"`
	} `json:"list"`
}

// GetOpenOrders :
func (s *V5OrderService) GetOpenOrders(param V5GetOpenOrdersParam) (*V5GetOpenOrdersResponse, error) {
	var res V5GetOpenOrdersResponse

	if param.Category == "" {
		return nil, fmt.Errorf("Category needed")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/order/realtime", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
