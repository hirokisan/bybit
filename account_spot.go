package bybit

import (
	"errors"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// SpotPostOrderParam :
type SpotPostOrderParam struct {
	Symbol SymbolSpot    `url:"symbol"`
	Qty    float64       `url:"qty"`
	Side   Side          `url:"side"`
	Type   OrderTypeSpot `url:"type"`

	TimeInForce *TimeInForceSpot `url:"timeInForce,omitempty"`
	Price       *float64         `url:"price,omitempty"`
	OrderLinkID *string          `url:"orderLinkId,omitempty"`
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

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.Client.postForm("/spot/v1/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotGetOrderParam :
type SpotGetOrderParam struct {
	OrderID     *string `url:"orderId,omitempty"`
	OrderLinkID *string `url:"orderLinkId,omitempty"`
}

// SpotGetOrderResponse :
type SpotGetOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotGetOrderResult `json:"result"`
}

// SpotGetOrderResult :
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

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.Client.getPrivately("/spot/v1/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotDeleteOrderParam :
type SpotDeleteOrderParam struct {
	OrderID     *string `url:"orderId,omitempty"`
	OrderLinkID *string `url:"orderLinkId,omitempty"`
}

// SpotDeleteOrderResponse :
type SpotDeleteOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotDeleteOrderResult `json:"result"`
}

// SpotDeleteOrderResult :
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

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.Client.deletePrivately("/spot/v1/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotDeleteOrderFastParam struct {
	Symbol SymbolSpot `url:"symbolId"`

	OrderID     *string `url:"orderId,omitempty"`
	OrderLinkID *string `url:"orderLinkId,omitempty"`
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

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.Client.deletePrivately("/spot/v1/order/fast", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchCancelParam struct {
	Symbol SymbolSpot `url:"symbolId"`

	Side  *Side           `url:"side,omitempty"`
	Types []OrderTypeSpot `url:"orderTypes,omitempty" del:","`
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

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.Client.deletePrivately("/spot/order/batch-cancel", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchFastCancelParam struct {
	Symbol SymbolSpot `url:"symbolId"`

	Side  *Side           `url:"side,omitempty"`
	Types []OrderTypeSpot `url:"orderTypes,omitempty" del:","`
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

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.Client.deletePrivately("/spot/order/batch-fast-cancel", queryString, &res); err != nil {
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
