package bybit

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

// UnmarshalJSON :
func (r *SpotQuoteDepthBidsAsks) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	r.Price = parsedData[0][0]
	r.Quantity = parsedData[0][1]
	return nil
}

// SpotQuoteDepthBidsAsks :
type SpotQuoteDepthBidsAsks struct {
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
