package bybit

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// SpotV1ServiceI :
type SpotV1ServiceI interface {
	// Market Data Endpoints
	SpotSymbols() (*SpotSymbolsResponse, error)
	SpotQuoteDepth(SpotQuoteDepthParam) (*SpotQuoteDepthResponse, error)
	SpotQuoteDepthMerged(SpotQuoteDepthMergedParam) (*SpotQuoteDepthMergedResponse, error)
	SpotQuoteTrades(SpotQuoteTradesParam) (*SpotQuoteTradesResponse, error)
	SpotQuoteKline(SpotQuoteKlineParam) (*SpotQuoteKlineResponse, error)
	SpotQuoteTicker24hr(SpotQuoteTicker24hrParam) (*SpotQuoteTicker24hrResponse, error)
	SpotQuoteTickerPrice(SpotQuoteTickerPriceParam) (*SpotQuoteTickerPriceResponse, error)
	SpotQuoteTickerBookTicker(SpotQuoteTickerBookTickerParam) (*SpotQuoteTickerBookTickerResponse, error)

	// Account Data Endpoints
	SpotPostOrder(SpotPostOrderParam) (*SpotPostOrderResponse, error)
	SpotGetOrder(SpotGetOrderParam) (*SpotGetOrderResponse, error)
	SpotDeleteOrder(SpotDeleteOrderParam) (*SpotDeleteOrderResponse, error)
	SpotDeleteOrderFast(SpotDeleteOrderFastParam) (*SpotDeleteOrderFastResponse, error)
	SpotOrderBatchCancel(SpotOrderBatchCancelParam) (*SpotOrderBatchCancelResponse, error)
	SpotOrderBatchFastCancel(SpotOrderBatchFastCancelParam) (*SpotOrderBatchFastCancelResponse, error)
	SpotOrderBatchCancelByIDs(orderIDs []string) (*SpotOrderBatchCancelByIDsResponse, error)
	SpotOpenOrders(SpotOpenOrdersParam) (*SpotOpenOrdersResponse, error)

	// Wallet Data Endpoints
	SpotGetWalletBalance() (*SpotGetWalletBalanceResponse, error)
}

// SpotV1Service :
type SpotV1Service struct {
	client *Client
}

// SpotSymbolsResponse :
type SpotSymbolsResponse struct {
	CommonResponse `json:",inline"`
	Result         []SpotSymbolsResult `json:"result"`
}

// SpotSymbolsResult :
type SpotSymbolsResult struct {
	Name              string `json:"name"`
	Alias             string `json:"alias"`
	BaseCurrency      string `json:"baseCurrency"`
	QuoteCurrency     string `json:"quoteCurrency"`
	BasePrecision     string `json:"basePrecision"`
	QuotePrecision    string `json:"quotePrecision"`
	MinTradeQuantity  string `json:"minTradeQuantity"`
	MinTradeAmount    string `json:"minTradeAmount"`
	MinPricePrecision string `json:"minPricePrecision"`
	MaxTradeQuantity  string `json:"maxTradeQuantity"`
	MaxTradeAmount    string `json:"maxTradeAmount"`
	Category          int    `json:"category"`
}

// SpotSymbols :
func (s *SpotV1Service) SpotSymbols() (*SpotSymbolsResponse, error) {
	var res SpotSymbolsResponse

	if err := s.client.getPublicly("/spot/v1/symbols", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteDepthParam :
type SpotQuoteDepthParam struct {
	Symbol SymbolSpot `url:"symbol"`

	Limit *int `url:"limit,omitempty"`
}

// SpotQuoteDepthResponse :
type SpotQuoteDepthResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotQuoteDepthResult `json:"result"`
}

// SpotQuoteDepthResult :
type SpotQuoteDepthResult struct {
	Time int                    `json:"time"`
	Bids SpotQuoteDepthBidsAsks `json:"bids"`
	Asks SpotQuoteDepthBidsAsks `json:"asks"`
}

// SpotQuoteDepthBidsAsks :
type SpotQuoteDepthBidsAsks []SpotQuoteDepthBidAsk

// UnmarshalJSON :
func (r *SpotQuoteDepthBidsAsks) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	items := SpotQuoteDepthBidsAsks{}
	for _, item := range parsedData {
		item := item
		if len(item) != 2 {
			return errors.New("so far len(item) must be 2, please check it on documents")
		}
		items = append(items, SpotQuoteDepthBidAsk{
			Price:    item[0],
			Quantity: item[1],
		})
	}
	*r = items
	return nil
}

// SpotQuoteDepthBidAsk :
type SpotQuoteDepthBidAsk struct {
	Price    string
	Quantity string
}

// SpotQuoteDepth :
func (s *SpotV1Service) SpotQuoteDepth(param SpotQuoteDepthParam) (*SpotQuoteDepthResponse, error) {
	var res SpotQuoteDepthResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/spot/quote/v1/depth", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteDepthMergedParam :
type SpotQuoteDepthMergedParam struct {
	Symbol SymbolSpot `url:"symbol"`

	Scale *int `url:"scale,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// SpotQuoteDepthMergedResponse :
type SpotQuoteDepthMergedResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotQuoteDepthMergedResult `json:"result"`
}

// SpotQuoteDepthMergedResult :
type SpotQuoteDepthMergedResult struct {
	Time int                    `json:"time"`
	Bids SpotQuoteDepthBidsAsks `json:"bids"`
	Asks SpotQuoteDepthBidsAsks `json:"asks"`
}

// SpotQuoteDepthMerged :
func (s *SpotV1Service) SpotQuoteDepthMerged(param SpotQuoteDepthMergedParam) (*SpotQuoteDepthMergedResponse, error) {
	var res SpotQuoteDepthMergedResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/spot/quote/v1/depth/merged", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTradesParam :
type SpotQuoteTradesParam struct {
	Symbol SymbolSpot `url:"symbol"`

	Limit *int `url:"limit,omitempty"`
}

// SpotQuoteTradesResponse :
type SpotQuoteTradesResponse struct {
	CommonResponse `json:",inline"`
	Result         []SpotQuoteTradesResult `json:"result"`
}

// SpotQuoteTradesResult :
type SpotQuoteTradesResult struct {
	Price        string `json:"price"`
	Time         int    `json:"time"`
	Qty          string `json:"qty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

// SpotQuoteTrades :
func (s *SpotV1Service) SpotQuoteTrades(param SpotQuoteTradesParam) (*SpotQuoteTradesResponse, error) {
	var res SpotQuoteTradesResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/spot/quote/v1/trades", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteKlineParam :
type SpotQuoteKlineParam struct {
	Symbol   SymbolSpot `url:"symbol"`
	Interval Interval   `url:"interval"`

	Limit     *int `url:"limit,omitempty"`
	StartTime *int `url:"startTime,omitempty"`
	EndTime   *int `url:"endTime,omitempty"`
}

// SpotQuoteKlineResponse :
type SpotQuoteKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []SpotQuoteKlineResult `json:"result"`
}

// SpotQuoteKlineResult :
type SpotQuoteKlineResult struct {
	SpotQuoteKline SpotQuoteKline
}

// UnmarshalJSON :
func (r *SpotQuoteKlineResult) UnmarshalJSON(data []byte) error {
	parsedData := []interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	if len(parsedData) != 11 {
		return errors.New("so far len(items) must be 11, please check it on documents")
	}
	r.SpotQuoteKline = SpotQuoteKline{
		StartTime:        int(parsedData[0].(float64)),
		Open:             parsedData[1].(string),
		High:             parsedData[2].(string),
		Low:              parsedData[3].(string),
		Close:            parsedData[4].(string),
		Volume:           parsedData[5].(string),
		EndTime:          int(parsedData[6].(float64)),
		QuoteAssetVolume: parsedData[7].(string),
		Trades:           int(parsedData[8].(float64)),
		TakerBaseVolume:  parsedData[9].(float64),
		TakerQuoteVolume: parsedData[10].(float64),
	}
	return nil
}

// SpotQuoteKline :
type SpotQuoteKline struct {
	StartTime        int
	Open             string
	High             string
	Low              string
	Close            string
	Volume           string
	EndTime          int
	QuoteAssetVolume string
	Trades           int
	TakerBaseVolume  float64
	TakerQuoteVolume float64
}

// SpotQuoteKline :
func (s *SpotV1Service) SpotQuoteKline(param SpotQuoteKlineParam) (*SpotQuoteKlineResponse, error) {
	var res SpotQuoteKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/spot/quote/v1/kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTicker24hrParam :
type SpotQuoteTicker24hrParam struct {
	Symbol *SymbolSpot `url:"symbol,omitempty"`
}

// SpotQuoteTicker24hrResponse :
type SpotQuoteTicker24hrResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotQuoteTicker24hrResult `json:"result"`
}

// SpotQuoteTicker24hrResult :
type SpotQuoteTicker24hrResult struct {
	Time         int    `json:"time"`
	Symbol       string `json:"symbol"`
	BestBidPrice string `json:"bestBidPrice"`
	BestAskPrice string `json:"bestAskPrice"`
	LastPrice    string `json:"lastPrice"`
	OpenPrice    string `json:"openPrice"`
	HighPrice    string `json:"highPrice"`
	LowPrice     string `json:"lowPrice"`
	Volume       string `json:"volume"`
	QuoteVolume  string `json:"quoteVolume"`
}

// SpotQuoteTicker24hr :
func (s *SpotV1Service) SpotQuoteTicker24hr(param SpotQuoteTicker24hrParam) (*SpotQuoteTicker24hrResponse, error) {
	var res SpotQuoteTicker24hrResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/spot/quote/v1/ticker/24hr", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTickerPriceParam :
type SpotQuoteTickerPriceParam struct {
	Symbol *SymbolSpot `url:"symbol,omitempty"`
}

// SpotQuoteTickerPriceResponse :
type SpotQuoteTickerPriceResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotQuoteTickerPriceResult `json:"result"`
}

// SpotQuoteTickerPriceResult :
type SpotQuoteTickerPriceResult struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// SpotQuoteTickerPrice :
func (s *SpotV1Service) SpotQuoteTickerPrice(param SpotQuoteTickerPriceParam) (*SpotQuoteTickerPriceResponse, error) {
	var res SpotQuoteTickerPriceResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/spot/quote/v1/ticker/price", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTickerBookTickerParam :
type SpotQuoteTickerBookTickerParam struct {
	Symbol *SymbolSpot `url:"symbol,omitempty"`
}

// SpotQuoteTickerBookTickerResponse :
type SpotQuoteTickerBookTickerResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotQuoteTickerBookTickerResult `json:"result"`
}

// SpotQuoteTickerBookTickerResult :
type SpotQuoteTickerBookTickerResult struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int    `json:"time"`
}

// SpotQuoteTickerBookTicker :
func (s *SpotV1Service) SpotQuoteTickerBookTicker(param SpotQuoteTickerBookTickerParam) (*SpotQuoteTickerBookTickerResponse, error) {
	var res SpotQuoteTickerBookTickerResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/spot/quote/v1/ticker/book_ticker", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

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
func (s *SpotV1Service) SpotPostOrder(param SpotPostOrderParam) (*SpotPostOrderResponse, error) {
	var res SpotPostOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postForm("/spot/v1/order", queryString, &res); err != nil {
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
func (s *SpotV1Service) SpotGetOrder(param SpotGetOrderParam) (*SpotGetOrderResponse, error) {
	var res SpotGetOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/spot/v1/order", queryString, &res); err != nil {
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
func (s *SpotV1Service) SpotDeleteOrder(param SpotDeleteOrderParam) (*SpotDeleteOrderResponse, error) {
	var res SpotDeleteOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately("/spot/v1/order", queryString, &res); err != nil {
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
func (s *SpotV1Service) SpotDeleteOrderFast(param SpotDeleteOrderFastParam) (*SpotDeleteOrderFastResponse, error) {
	var res SpotDeleteOrderFastResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately("/spot/v1/order/fast", queryString, &res); err != nil {
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

func (s *SpotV1Service) SpotOrderBatchCancel(param SpotOrderBatchCancelParam) (*SpotOrderBatchCancelResponse, error) {
	var res SpotOrderBatchCancelResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately("/spot/order/batch-cancel", queryString, &res); err != nil {
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

func (s *SpotV1Service) SpotOrderBatchFastCancel(param SpotOrderBatchFastCancelParam) (*SpotOrderBatchFastCancelResponse, error) {
	var res SpotOrderBatchFastCancelResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately("/spot/order/batch-fast-cancel", queryString, &res); err != nil {
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
func (s *SpotV1Service) SpotOrderBatchCancelByIDs(orderIDs []string) (*SpotOrderBatchCancelByIDsResponse, error) {
	var res SpotOrderBatchCancelByIDsResponse

	if len(orderIDs) > 100 {
		return nil, errors.New("orderIDs length must be no more than 100")
	}

	query := url.Values{}
	query.Add("orderIds", strings.Join(orderIDs, ","))
	if err := s.client.deletePrivately("/spot/order/batch-cancel-by-ids", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotOpenOrdersParam :
type SpotOpenOrdersParam struct {
	Symbol  *SymbolSpot `url:"symbol,omitempty"`
	OrderID *string     `url:"orderId,omitempty"`
	Limit   *int        `url:"limit,omitempty"`
}

// SpotOpenOrdersResponse :
type SpotOpenOrdersResponse struct {
	CommonResponse `json:",inline"`
	Result         []SpotOpenOrdersResult `json:"result"`
}

// SpotOpenOrdersResult :
type SpotOpenOrdersResult struct {
	AccountID           string `json:"accountId"`
	ExchangeID          string `json:"exchangeId"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbolName"`
	OrderLinkID         string `json:"orderLinkId"`
	OrderID             string `json:"orderId"`
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

// SpotOpenOrders :
func (s *SpotV1Service) SpotOpenOrders(param SpotOpenOrdersParam) (*SpotOpenOrdersResponse, error) {
	var res SpotOpenOrdersResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/spot/v1/open-orders", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotGetWalletBalanceResponse :
type SpotGetWalletBalanceResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotGetWalletBalanceResult `json:"result"`
}

// SpotGetWalletBalanceResult :
type SpotGetWalletBalanceResult struct {
	Balances []SpotGetWalletBalanceResultBalance `json:"balances"`
}

// SpotGetWalletBalanceResultBalance :
type SpotGetWalletBalanceResultBalance struct {
	Coin     string `json:"coin"`
	CoinID   string `json:"coinId"`
	CoinName string `json:"coinName"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Locked   string `json:"locked"`
}

// SpotGetWalletBalance :
func (s *SpotV1Service) SpotGetWalletBalance() (*SpotGetWalletBalanceResponse, error) {
	var res SpotGetWalletBalanceResponse

	queryString, err := query.Values(nil)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately("/spot/v1/account", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
