package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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
func (s *MarketService) SpotSymbols() (*SpotSymbolsResponse, error) {
	var res SpotSymbolsResponse

	url, err := s.Client.BuildPublicURL("/spot/v1/symbols", nil)
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

// SpotQuoteDepthParam :
type SpotQuoteDepthParam struct {
	Symbol SymbolSpot `json:"symbol"`

	Limit *int `json:"limit"`
}

func (p *SpotQuoteDepthParam) build() map[string]string {
	result := map[string]string{
		"symbol": string(p.Symbol),
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	return result
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
		fmt.Println("item", item)
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
func (s *MarketService) SpotQuoteDepth(param SpotQuoteDepthParam) (*SpotQuoteDepthResponse, error) {
	var res SpotQuoteDepthResponse

	url, err := s.Client.BuildPublicURL("/spot/quote/v1/depth", param.build())
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

// SpotQuoteDepthMergedParam :
type SpotQuoteDepthMergedParam struct {
	Symbol SymbolSpot `json:"symbol"`

	Scale *int `json:"scale"`
	Limit *int `json:"limit"`
}

func (p *SpotQuoteDepthMergedParam) build() map[string]string {
	result := map[string]string{
		"symbol": string(p.Symbol),
	}
	if p.Scale != nil {
		result["scale"] = strconv.Itoa(*p.Scale)
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	return result
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
func (s *MarketService) SpotQuoteDepthMerged(param SpotQuoteDepthMergedParam) (*SpotQuoteDepthMergedResponse, error) {
	var res SpotQuoteDepthMergedResponse

	url, err := s.Client.BuildPublicURL("/spot/quote/v1/depth/merged", param.build())
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

// SpotQuoteTradesParam :
type SpotQuoteTradesParam struct {
	Symbol SymbolSpot `json:"symbol"`

	Limit *int `json:"limit"`
}

func (p *SpotQuoteTradesParam) build() map[string]string {
	result := map[string]string{
		"symbol": string(p.Symbol),
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	return result
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
func (s *MarketService) SpotQuoteTrades(param SpotQuoteTradesParam) (*SpotQuoteTradesResponse, error) {
	var res SpotQuoteTradesResponse

	url, err := s.Client.BuildPublicURL("/spot/quote/v1/trades", param.build())
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

// SpotQuoteKlineParam :
type SpotQuoteKlineParam struct {
	Symbol   SymbolSpot `json:"symbol"`
	Interval Interval   `json:"interval"`

	Limit     *int `json:"limit"`
	StartTime *int `json:"startTime"`
	EndTime   *int `json:"endTime"`
}

func (p *SpotQuoteKlineParam) build() map[string]string {
	result := map[string]string{
		"symbol":   string(p.Symbol),
		"interval": string(p.Interval),
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	if p.StartTime != nil {
		result["startTime"] = strconv.Itoa(*p.StartTime)
	}
	if p.EndTime != nil {
		result["endTime"] = strconv.Itoa(*p.EndTime)
	}
	return result
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
		TakerBaseVolume:  parsedData[9].(string),
		TakerQuoteVolume: parsedData[10].(string),
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
	TakerBaseVolume  string
	TakerQuoteVolume string
}

// SpotQuoteKline :
func (s *MarketService) SpotQuoteKline(param SpotQuoteKlineParam) (*SpotQuoteKlineResponse, error) {
	var res SpotQuoteKlineResponse

	url, err := s.Client.BuildPublicURL("/spot/quote/v1/kline", param.build())
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

// SpotQuoteTicker24hrParam :
type SpotQuoteTicker24hrParam struct {
	Symbol *SymbolSpot `json:"symbol"`
}

func (p *SpotQuoteTicker24hrParam) build() map[string]string {
	if p.Symbol == nil {
		return nil
	}
	result := map[string]string{
		"symbol": string(*p.Symbol),
	}
	return result
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
func (s *MarketService) SpotQuoteTicker24hr(param SpotQuoteTicker24hrParam) (*SpotQuoteTicker24hrResponse, error) {
	var res SpotQuoteTicker24hrResponse

	url, err := s.Client.BuildPublicURL("/spot/quote/v1/ticker/24hr", param.build())
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

// SpotQuoteTickerPriceParam :
type SpotQuoteTickerPriceParam struct {
	Symbol *SymbolSpot `json:"symbol"`
}

func (p *SpotQuoteTickerPriceParam) build() map[string]string {
	if p.Symbol == nil {
		return nil
	}
	result := map[string]string{
		"symbol": string(*p.Symbol),
	}
	return result
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
func (s *MarketService) SpotQuoteTickerPrice(param SpotQuoteTickerPriceParam) (*SpotQuoteTickerPriceResponse, error) {
	var res SpotQuoteTickerPriceResponse

	url, err := s.Client.BuildPublicURL("/spot/quote/v1/ticker/price", param.build())
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

// SpotQuoteTickerBookTickerParam :
type SpotQuoteTickerBookTickerParam struct {
	Symbol *SymbolSpot `json:"symbol"`
}

func (p *SpotQuoteTickerBookTickerParam) build() map[string]string {
	if p.Symbol == nil {
		return nil
	}
	result := map[string]string{
		"symbol": string(*p.Symbol),
	}
	return result
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
func (s *MarketService) SpotQuoteTickerBookTicker(param SpotQuoteTickerBookTickerParam) (*SpotQuoteTickerBookTickerResponse, error) {
	var res SpotQuoteTickerBookTickerResponse

	url, err := s.Client.BuildPublicURL("/spot/quote/v1/ticker/book_ticker", param.build())
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
func (s *MarketService) SpotPostOrder(param SpotPostOrderParam) (*SpotPostOrderResponse, error) {
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
