package bybit

import (
	"encoding/json"
	"net/http"
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

func (p SpotPostOrderParam) build() map[string]string {
	ps := map[string]string{}
	ps["symbol"] = string(p.Symbol)
	ps["qty"] = strconv.FormatFloat(p.Qty, 'f', 2, 64)
	ps["side"] = string(p.Side)
	ps["type"] = string(p.Type)
	if p.Price != nil {
		ps["price"] = strconv.FormatFloat(*p.Price, 'f', 2, 64)
	}
	if p.TimeInForce != nil {
		ps["timeInForce"] = string(*p.TimeInForce)
	}
	if p.OrderLinkID != nil {
		ps["orderLinkId"] = string(*p.OrderLinkID)
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

	url, err := s.Client.BuildPublicURL("/spot/v1/order", nil)
	if err != nil {
		return nil, err
	}

	ps := param.build()
	s.Client.populateSignature(ps)

	body := strings.NewReader(encodeURLParamsFrom(ps))
	resp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

type SpotGetOrderParam struct {
	OrderID     *string `json:"orderId"`
	OrderLinkID *string `json:"orderLinkId"`
}

func (p SpotGetOrderParam) build() map[string]string {
	result := map[string]string{}
	if p.OrderID != nil {
		result["orderId"] = *p.OrderID
	}
	if p.OrderLinkID != nil {
		result["orderLinkId"] = *p.OrderLinkID
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

	url, err := s.Client.BuildPrivateURL("/spot/v1/order", param.build())
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

type SpotDeleteOrderParam struct {
	OrderID     *string `json:"orderId"`
	OrderLinkID *string `json:"orderLinkId"`
}

func (p SpotDeleteOrderParam) build() map[string]string {
	result := map[string]string{}
	if p.OrderID != nil {
		result["orderId"] = *p.OrderID
	}
	if p.OrderLinkID != nil {
		result["orderLinkId"] = *p.OrderLinkID
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

	url, err := s.Client.BuildPrivateURL("/spot/v1/order", param.build())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

type SpotDeleteOrderFastParam struct {
	Symbol SymbolSpot `json:"symbolId"`

	OrderID     *string `json:"orderId"`
	OrderLinkID *string `json:"orderLinkId"`
}

func (p SpotDeleteOrderFastParam) build() map[string]string {
	result := map[string]string{
		"symbolId": string(p.Symbol),
	}
	if p.OrderID != nil {
		result["orderId"] = *p.OrderID
	}
	if p.OrderLinkID != nil {
		result["orderLinkId"] = *p.OrderLinkID
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

	url, err := s.Client.BuildPrivateURL("/spot/v1/order/fast", param.build())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

type SpotOrderBatchCancelParam struct {
	Symbol SymbolSpot `json:"symbolId"`

	Side  *Side           `json:"side"`
	Types []OrderTypeSpot `json:"orderTypes"`
}

func (p SpotOrderBatchCancelParam) build() map[string]string {
	result := map[string]string{
		"symbolId": string(p.Symbol),
	}
	if p.Side != nil {
		result["side"] = string(*p.Side)
	}
	if len(p.Types) != 0 {
		var types []string
		for _, t := range p.Types {
			types = append(types, string(t))
		}
		result["orderTypes"] = strings.Join(types, ",")
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

	url, err := s.Client.BuildPrivateURL("/spot/v1/order/batch-cancel", param.build())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
