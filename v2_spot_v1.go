package bybit

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-querystring/query"
)

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
