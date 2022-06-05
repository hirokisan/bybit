package bybit

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// SpotPostOrderParam :
type SpotPostOrderParam struct {
	Symbol SymbolSpot    `json:"symbol"`
	Qty    float64       `json:"qty"`
	Side   Side          `json:"side"`
	Type   OrderTypeSpot `json:"type"`

	TimeInForce *TimeInForceSpot `json:"timeInForce"`
	Price       *float64         `json:"price"`
	OrderLinkID *string          `json:"orderLinkId"`
}

func (p SpotPostOrderParam) build() url.Values {
	ps := url.Values{}
	ps.Add("symbol", string(p.Symbol))
	ps.Add("qty", strconv.FormatFloat(p.Qty, 'f', 2, 64))
	ps.Add("side", string(p.Side))
	ps.Add("type", string(p.Type))

	if p.Price != nil {
		ps.Add("price", strconv.FormatFloat(*p.Price, 'f', 2, 64))
	}
	if p.TimeInForce != nil {
		ps.Add("timeInForce", string(*p.TimeInForce))
	}
	if p.OrderLinkID != nil {
		ps.Add("orderLinkId", string(*p.OrderLinkID))
	}
	return ps
}

// SpotPostOrderResponse :
type SpotPostOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotPostOrderResult `json:"result"`
}

// SpotPostOrderResult :
type SpotPostOrderResult struct {
	OrderID      string          `json:"orderId"`
	OrderLinkID  string          `json:"orderLinkId"`
	Symbol       string          `json:"symbol"`
	TransactTime string          `json:"transactTime"`
	Price        string          `json:"price"`
	OrigQty      string          `json:"origQty"`
	Type         OrderTypeSpot   `json:"type"`
	Side         string          `json:"side"`
	Status       OrderStatusSpot `json:"status"`
	TimeInForce  TimeInForceSpot `json:"timeInForce"`
	AccountID    string          `json:"accountId"`
	SymbolName   string          `json:"symbolName"`
	ExecutedQty  string          `json:"executedQty"`
}

// SpotPostOrder :
func (s *AccountService) SpotPostOrder(param SpotPostOrderParam) (*SpotPostOrderResponse, error) {
	var res SpotPostOrderResponse

	if err := s.Client.postForm("/spot/v1/order", param.build(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotGetOrderParam struct {
	OrderID     *string `json:"orderId"`
	OrderLinkID *string `json:"orderLinkId"`
}

func (p SpotGetOrderParam) build() url.Values {
	result := url.Values{}
	if p.OrderID != nil {
		result.Add("orderId", *p.OrderID)
	}
	if p.OrderLinkID != nil {
		result.Add("orderLinkId", *p.OrderLinkID)
	}
	return result
}

type SpotGetOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotGetOrderResult `json:"result"`
}

type SpotGetOrderResult struct {
	AccountId           string `json:"accountId"`
	ExchangeId          string `json:"exchangeId"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbolName"`
	OrderLinkId         string `json:"orderLinkId"`
	OrderId             string `json:"orderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
}

// SpotGetOrder :
func (s *AccountService) SpotGetOrder(param SpotGetOrderParam) (*SpotGetOrderResponse, error) {
	var res SpotGetOrderResponse

	if err := s.Client.getPrivately("/spot/v1/order", param.build(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotDeleteOrderParam struct {
	OrderID     *string `json:"orderId"`
	OrderLinkID *string `json:"orderLinkId"`
}

func (p SpotDeleteOrderParam) build() url.Values {
	result := url.Values{}
	if p.OrderID != nil {
		result.Add("orderId", *p.OrderID)
	}
	if p.OrderLinkID != nil {
		result.Add("orderLinkId", *p.OrderLinkID)
	}
	return result
}

type SpotDeleteOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotDeleteOrderResult `json:"result"`
}

type SpotDeleteOrderResult struct {
	OrderId      string `json:"orderId"`
	OrderLinkId  string `json:"orderLinkId"`
	Symbol       string `json:"symbol"`
	Status       string `json:"status"`
	AccountId    string `json:"accountId"`
	TransactTime string `json:"transactTime"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	ExecutedQty  string `json:"executedQty"`
	TimeInForce  string `json:"timeInForce"`
	Type         string `json:"type"`
	Side         string `json:"side"`
}

// SpotDeleteOrder :
func (s *AccountService) SpotDeleteOrder(param SpotDeleteOrderParam) (*SpotDeleteOrderResponse, error) {
	var res SpotDeleteOrderResponse

	if err := s.Client.deletePrivately("/spot/v1/order", param.build(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotDeleteOrderFastParam struct {
	Symbol SymbolSpot `json:"symbolId"`

	OrderID     *string `json:"orderId"`
	OrderLinkID *string `json:"orderLinkId"`
}

func (p SpotDeleteOrderFastParam) build() url.Values {
	result := url.Values{}
	result.Add("symbolId", string(p.Symbol))

	if p.OrderID != nil {
		result.Add("orderId", *p.OrderID)
	}
	if p.OrderLinkID != nil {
		result.Add("orderLinkId", *p.OrderLinkID)
	}
	return result
}

type SpotDeleteOrderFastResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotDeleteOrderFastResult `json:"result"`
}

type SpotDeleteOrderFastResult struct {
	IsCancelled bool `json:"isCancelled"`
}

// SpotDeleteOrderFast :
func (s *AccountService) SpotDeleteOrderFast(param SpotDeleteOrderFastParam) (*SpotDeleteOrderFastResponse, error) {
	var res SpotDeleteOrderFastResponse

	if err := s.Client.deletePrivately("/spot/v1/order/fast", param.build(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchCancelParam struct {
	Symbol SymbolSpot `json:"symbolId"`

	Side  *Side           `json:"side"`
	Types []OrderTypeSpot `json:"orderTypes"`
}

func (p SpotOrderBatchCancelParam) build() url.Values {
	result := url.Values{}
	result.Add("symbolId", string(p.Symbol))

	if p.Side != nil {
		result.Add("side", string(*p.Side))
	}
	if len(p.Types) != 0 {
		var types []string
		for _, t := range p.Types {
			types = append(types, string(t))
		}
		result.Add("orderTypes", strings.Join(types, ","))
	}
	return result
}

type SpotOrderBatchCancelResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotOrderBatchCancelResult `json:"result"`
}

type SpotOrderBatchCancelResult struct {
	Success bool `json:"success"`
}

func (s *AccountService) SpotOrderBatchCancel(param SpotOrderBatchCancelParam) (*SpotOrderBatchCancelResponse, error) {
	var res SpotOrderBatchCancelResponse

	if err := s.Client.deletePrivately("/spot/order/batch-cancel", param.build(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchFastCancelParam struct {
	Symbol SymbolSpot `json:"symbolId"`

	Side  *Side           `json:"side"`
	Types []OrderTypeSpot `json:"orderTypes"`
}

func (p SpotOrderBatchFastCancelParam) build() url.Values {
	result := url.Values{}
	result.Add("symbolId", string(p.Symbol))
	if p.Side != nil {
		result.Add("side", string(*p.Side))
	}
	if len(p.Types) != 0 {
		var types []string
		for _, t := range p.Types {
			types = append(types, string(t))
		}
		result.Add("orderTypes", strings.Join(types, ","))
	}
	return result
}

type SpotOrderBatchFastCancelResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotOrderBatchFastCancelResult `json:"result"`
}

type SpotOrderBatchFastCancelResult struct {
	Success bool `json:"success"`
}

func (s *AccountService) SpotOrderBatchFastCancel(param SpotOrderBatchFastCancelParam) (*SpotOrderBatchFastCancelResponse, error) {
	var res SpotOrderBatchFastCancelResponse

	if err := s.Client.deletePrivately("/spot/order/batch-fast-cancel", param.build(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchCancelByIDsResponse struct {
	CommonResponse `json:",inline"`
	Result         []SpotOrderBatchCancelByIDsResult `json:"result"`
}

type SpotOrderBatchCancelByIDsResult struct {
	OrderID string `json:"orderId"`
	Code    string `json:"code"`
}

// TODO : have bug multiple orderIds
func (s *AccountService) SpotOrderBatchCancelByIDs(orderIDs []string) (*SpotOrderBatchCancelByIDsResponse, error) {
	var res SpotOrderBatchCancelByIDsResponse

	if len(orderIDs) > 100 {
		return nil, errors.New("orderIDs length must be no more than 100")
	}

	query := url.Values{}
	query.Add("orderIds", strings.Join(orderIDs, ","))
	if err := s.Client.deletePrivately("/spot/order/batch-cancel-by-ids", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
