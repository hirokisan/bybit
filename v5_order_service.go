package bybit

import (
	"encoding/json"
	"fmt"
)

// V5OrderServiceI :
type V5OrderServiceI interface {
	CreateOrder(V5CreateOrderParam) (*V5CreateOrderResponse, error)
	CancelOrder(V5CancelOrderParam) (*V5CancelOrderResponse, error)
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

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/order/cancel", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}
