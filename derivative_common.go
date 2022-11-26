package bybit

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-querystring/query"
)

// DerivativeCommonService :
type DerivativeCommonService struct {
	client *Client
}

// DerivativesOrderBookResponse :
type DerivativesOrderBookResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesOrderBookResult `json:"result"`
}

// DerivativesOrderBookResult :
type DerivativesOrderBookResult struct {
	Symbol    SymbolDerivative                  `json:"s"`
	Buyers    DerivativesOrderBookResultBuyers  `json:"b"`
	Sellers   DerivativesOrderBookResultSellers `json:"a"`
	Timestamp int                               `json:"ts"`
	ID        int                               `json:"u"`
}

// DerivativesOrderBookResultBuyers :
type DerivativesOrderBookResultBuyers []DerivativesOrderBookResultBuyer

// UnmarshalJSON :
func (r *DerivativesOrderBookResultBuyers) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	items := DerivativesOrderBookResultBuyers{}
	for _, item := range parsedData {
		item := item
		price, qty := item[0], item[1]
		items = append(items, DerivativesOrderBookResultBuyer{
			Price: price,
			Qty:   qty,
		})
	}
	*r = items
	return nil
}

// DerivativesOrderBookResultBuyer :
type DerivativesOrderBookResultBuyer struct {
	Price string
	Qty   string
}

// DerivativesOrderBookResultSellers :
type DerivativesOrderBookResultSellers []DerivativesOrderBookResultSeller

// UnmarshalJSON :
func (r *DerivativesOrderBookResultSellers) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	items := DerivativesOrderBookResultSellers{}
	for _, item := range parsedData {
		item := item
		price, qty := item[0], item[1]
		items = append(items, DerivativesOrderBookResultSeller{
			Price: price,
			Qty:   qty,
		})
	}
	*r = items
	return nil
}

// DerivativesOrderBookResultSeller :
type DerivativesOrderBookResultSeller struct {
	Price string
	Qty   string
}

// DerivativesOrderBookParam :
type DerivativesOrderBookParam struct {
	Symbol   SymbolDerivative   `url:"symbol"`
	Category CategoryDerivative `url:"category"`

	Limit *int `url:"limit,omitempty"`
}

// DerivativesOrderBook :
func (s *DerivativeCommonService) DerivativesOrderBook(param DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error) {
	var res DerivativesOrderBookResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/derivatives/v3/public/order-book/L2", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DerivativesKlineResponse :
type DerivativesKlineResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesKlineResult `json:"result"`
}

// DerivativesKlineResult :
type DerivativesKlineResult struct {
	Category CategoryDerivative           `json:"category"`
	Symbol   SymbolDerivative             `json:"symbol"`
	Lists    []DerivativesKlineResultList `json:"list"`
}

// DerivativesKlineResultList :
type DerivativesKlineResultList struct {
	Start    string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	Turnover string
}

// UnmarshalJSON :
func (r *DerivativesKlineResultList) UnmarshalJSON(data []byte) error {
	parsedData := []interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	if len(parsedData) != 7 {
		return errors.New("so far len(items) must be 7, please check it on documents")
	}
	*r = DerivativesKlineResultList{
		Start:    parsedData[0].(string),
		Open:     parsedData[1].(string),
		High:     parsedData[2].(string),
		Low:      parsedData[3].(string),
		Close:    parsedData[4].(string),
		Volume:   parsedData[5].(string),
		Turnover: parsedData[6].(string),
	}
	return nil
}

// DerivativesKlineParam :
type DerivativesKlineParam struct {
	Symbol   SymbolDerivative   `url:"symbol"`
	Category CategoryDerivative `url:"category"`
	Interval Interval           `url:"interval"`
	Start    int                `url:"start"` // timestamp point for result, in milliseconds
	End      int                `url:"end"`   // timestamp point for result, in milliseconds

	Limit *int `url:"limit,omitempty"`
}

// DerivativesKline :
func (s *DerivativeCommonService) DerivativesKline(param DerivativesKlineParam) (*DerivativesKlineResponse, error) {
	var res DerivativesKlineResponse

	if param.Category != CategoryDerivativeInverse && param.Category != CategoryDerivativeLinear {
		return nil, fmt.Errorf("only inverse and linear supported, but %s given for category", param.Category)
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/derivatives/v3/public/kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DerivativesTickersResponse :
type DerivativesTickersResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesTickersResult `json:"result"`
}

// DerivativesTickersResult :
type DerivativesTickersResult struct {
	Category CategoryDerivative             `json:"category"`
	Lists    []DerivativesTickersResultList `json:"list"`
}

// DerivativesTickersResultList :
type DerivativesTickersResultList struct {
	Symbol                 SymbolDerivative `json:"symbol"`
	BidPrice               string           `json:"bidPrice"`
	AskPrice               string           `json:"askPrice"`
	LastPrice              string           `json:"lastPrice"`
	LastTickDirection      string           `json:"lastTickDirection"`
	PrevPrice24h           string           `json:"prevPrice24h"`
	Price24hPcnt           string           `json:"price24hPcnt"`
	HighPrice24h           string           `json:"highPrice24h"`
	LowPrice24h            string           `json:"lowPrice24h"`
	PrevPrice1h            string           `json:"prevPrice1h"`
	MarkPrice              string           `json:"markPrice"`
	IndexPrice             string           `json:"indexPrice"`
	OpenInterest           string           `json:"openInterest"`
	Turnover24h            string           `json:"turnover24h"`
	Volume24h              string           `json:"volume24h"`
	FundingRate            string           `json:"fundingRate"`
	NextFundingTime        string           `json:"nextFundingTime"`
	PredictedDeliveryPrice string           `json:"predictedDeliveryPrice"` // Applicable to inverse future and option
	BasisRate              string           `json:"basisRate"`
	DeliveryFeeRate        string           `json:"deliveryFeeRate"`
	DeliveryTime           string           `json:"deliveryTime"`
}

// DerivativesTickersParam :
type DerivativesTickersParam struct {
	Category CategoryDerivative `url:"category"`

	Symbol *SymbolDerivative `url:"symbol,omitempty"`
}

// DerivativesTickers :
func (s *DerivativeCommonService) DerivativesTickers(param DerivativesTickersParam) (*DerivativesTickersResponse, error) {
	var res DerivativesTickersResponse

	if param.Category == CategoryDerivativeOption {
		return nil, errors.New("call DerivativesTickersForOption instead")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/derivatives/v3/public/tickers", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DerivativesTickersForOptionResponse :
type DerivativesTickersForOptionResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesTickersForOptionResult `json:"result"`
}

// DerivativesTickersForOptionResult :
type DerivativesTickersForOptionResult struct {
	Category               CategoryDerivative `json:"category"`
	Symbol                 SymbolDerivative   `json:"symbol"`
	BidPrice               string             `json:"bidPrice"`
	BidSize                string             `json:"bidSize"`
	BidIv                  string             `json:"bidIv"`
	AskPrice               string             `json:"askPrice"`
	AskSize                string             `json:"askSize"`
	AskIv                  string             `json:"askIv"`
	LastPrice              string             `json:"lastPrice"`
	HighPrice24h           string             `json:"highPrice24h"`
	LowPrice24h            string             `json:"lowPrice24h"`
	MarkPrice              string             `json:"markPrice"`
	IndexPrice             string             `json:"indexPrice"`
	MarkPriceIv            string             `json:"markPriceIv"`
	UnderlyingPrice        string             `json:"underlyingPrice"`
	OpenInterest           string             `json:"openInterest"`
	Turnover24h            string             `json:"turnover24h"`
	Volume24h              string             `json:"volume24h"`
	TotalVolume            string             `json:"totalVolume"`
	TotalTurnover          string             `json:"totalTurnover"`
	Delta                  string             `json:"delta"`
	Gamma                  string             `json:"gamma"`
	Vega                   string             `json:"vega"`
	Theta                  string             `json:"theta"`
	PredictedDeliveryPrice string             `json:"predictedDeliveryPrice"`
	Change24h              string             `json:"change24h"`
}

// DerivativesTickersForOptionParam :
type DerivativesTickersForOptionParam struct {
	Symbol SymbolDerivative `url:"symbol"`
}

// DerivativesTickersForOption :
func (s *DerivativeCommonService) DerivativesTickersForOption(param DerivativesTickersForOptionParam) (*DerivativesTickersForOptionResponse, error) {
	var res DerivativesTickersForOptionResponse

	queryString, err := query.Values(DerivativesTickersParam{
		Category: CategoryDerivativeOption,
		Symbol:   &param.Symbol,
	})
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/derivatives/v3/public/tickers", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DerivativesInstrumentsResponse :
type DerivativesInstrumentsResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesInstrumentsResult `json:"result"`
}

// DerivativesInstrumentsResult :
type DerivativesInstrumentsResult struct {
	Category CategoryDerivative `json:"category"`
	List     []struct {
		Symbol          SymbolDerivative       `json:"symbol"`
		ContractType    ContractTypeDerivative `json:"contractType"`
		Status          StatusDerivative       `json:"status"`
		BaseCoin        string                 `json:"baseCoin"`
		QuoteCoin       string                 `json:"quoteCoin"`
		LaunchTime      string                 `json:"launchTime"`
		DeliveryTime    string                 `json:"deliveryTime"`
		DeliveryFeeRate string                 `json:"deliveryFeeRate"`
		PriceScale      string                 `json:"priceScale"`
		LeverageFilter  struct {
			MinLeverage  string `json:"minLeverage"`
			MaxLeverage  string `json:"maxLeverage"`
			LeverageStep string `json:"leverageStep"`
		} `json:"leverageFilter"`
		PriceFilter struct {
			MinPrice string `json:"minPrice"`
			MaxPrice string `json:"maxPrice"`
			TickSize string `json:"tickSize"`
		} `json:"priceFilter"`
		LotSizeFilter struct {
			MaxTradingQty string `json:"maxTradingQty"`
			MinTradingQty string `json:"minTradingQty"`
			QtyStep       string `json:"qtyStep"`
		} `json:"lotSizeFilter"`
	} `json:"list"`
	NextPageCursor string `json:"nextPageCursor"`
}

// DerivativesInstrumentsParam :
type DerivativesInstrumentsParam struct {
	Category CategoryDerivative `url:"category"`

	Symbol *SymbolDerivative `url:"symbol,omitempty"`
	Limit  *int              `url:"limit,omitempty"`
	Cursor *string           `url:"cursor,omitempty"`
}

// DerivativesInstruments :
func (s *DerivativeCommonService) DerivativesInstruments(param DerivativesInstrumentsParam) (*DerivativesInstrumentsResponse, error) {
	var res DerivativesInstrumentsResponse

	if param.Category == CategoryDerivativeOption {
		return nil, errors.New("call DerivativesInstrumentsForOption instead")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/derivatives/v3/public/instruments-info", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DerivativesInstrumentsForOptionResponse :
type DerivativesInstrumentsForOptionResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesInstrumentsForOptionResult `json:"result"`
}

// DerivativesInstrumentsForOptionResult :
type DerivativesInstrumentsForOptionResult struct {
	ResultTotalSize int    `json:"resultTotalSize"`
	Cursor          string `json:"cursor"`
	List            []struct {
		Category        CategoryDerivative `json:"category"`
		Symbol          SymbolDerivative   `json:"symbol"`
		Status          StatusDerivative   `json:"status"`
		BaseCoin        string             `json:"baseCoin"`
		QuoteCoin       string             `json:"quoteCoin"`
		SettleCoin      string             `json:"settleCoin"`
		OptionsType     string             `json:"optionsType"`
		LaunchTime      string             `json:"launchTime"`
		DeliveryTime    string             `json:"deliveryTime"`
		DeliveryFeeRate string             `json:"deliveryFeeRate"`
		PriceFilter     struct {
			MinPrice string `json:"minPrice"`
			MaxPrice string `json:"maxPrice"`
			TickSize string `json:"tickSize"`
		} `json:"priceFilter"`
		LotSizeFilter struct {
			MaxOrderQty string `json:"maxOrderQty"`
			MinOrderQty string `json:"minOrderQty"`
			QtyStep     string `json:"qtyStep"`
		} `json:"lotSizeFilter"`
	} `json:"dataList"`
}

// DerivativesInstrumentsForOptionParam :
type DerivativesInstrumentsForOptionParam struct {
	Symbol *SymbolDerivative `url:"symbol,omitempty"`
	Limit  *int              `url:"limit,omitempty"`
	Cursor *string           `url:"cursor,omitempty"`
}

// DerivativesInstrumentsForOption :
func (s *DerivativeCommonService) DerivativesInstrumentsForOption(param DerivativesInstrumentsForOptionParam) (*DerivativesInstrumentsForOptionResponse, error) {
	var res DerivativesInstrumentsForOptionResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	queryString.Add("category", string(CategoryDerivativeOption))

	if err := s.client.getPublicly("/derivatives/v3/public/instruments-info", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DerivativesMarkPriceKlineResponse :
type DerivativesMarkPriceKlineResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesMarkPriceKlineResult `json:"result"`
}

// DerivativesMarkPriceKlineResult :
type DerivativesMarkPriceKlineResult struct {
	Category CategoryDerivative                        `json:"category"`
	Symbol   SymbolDerivative                          `json:"symbol"`
	List     []DerivativesMarkPriceKlineResultListItem `json:"list"`
}

// DerivativesMarkPriceKlineResultListItem :
type DerivativesMarkPriceKlineResultListItem struct {
	Start string `json:"start"`
	Open  string `json:"open"`
	High  string `json:"high"`
	Low   string `json:"low"`
	Close string `json:"close"`
}

// UnmarshalJSON :
func (r *DerivativesMarkPriceKlineResultListItem) UnmarshalJSON(data []byte) error {
	parsedData := []interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	if len(parsedData) != 5 {
		return errors.New("so far len(items) must be 5, please check it on documents")
	}
	*r = DerivativesMarkPriceKlineResultListItem{
		Start: parsedData[0].(string),
		Open:  parsedData[1].(string),
		High:  parsedData[2].(string),
		Low:   parsedData[3].(string),
		Close: parsedData[4].(string),
	}
	return nil
}

// DerivativesMarkPriceKlineParam :
type DerivativesMarkPriceKlineParam struct {
	Category CategoryDerivative `url:"category"`
	Symbol   SymbolDerivative   `url:"symbol"`
	Interval Interval           `url:"interval"`
	Start    int                `url:"start"` // timestamp point for result, in milliseconds
	End      int                `url:"end"`   // timestamp point for result, in milliseconds

	Limit *int `url:"limit,omitempty"`
}

// DerivativesMarkPriceKline :
func (s *DerivativeCommonService) DerivativesMarkPriceKline(param DerivativesMarkPriceKlineParam) (*DerivativesMarkPriceKlineResponse, error) {
	var res DerivativesMarkPriceKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/derivatives/v3/public/mark-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DerivativesIndexPriceKlineResponse :
type DerivativesIndexPriceKlineResponse struct {
	CommonV3Response `json:",inline"`
	Result           DerivativesIndexPriceKlineResult `json:"result"`
}

// DerivativesIndexPriceKlineResult :
type DerivativesIndexPriceKlineResult struct {
	Category CategoryDerivative                         `json:"category"`
	Symbol   SymbolDerivative                           `json:"symbol"`
	List     []DerivativesIndexPriceKlineResultListItem `json:"list"`
}

// DerivativesIndexPriceKlineResultListItem :
type DerivativesIndexPriceKlineResultListItem struct {
	Start string `json:"start"`
	Open  string `json:"open"`
	High  string `json:"high"`
	Low   string `json:"low"`
	Close string `json:"close"`
}

// UnmarshalJSON :
func (r *DerivativesIndexPriceKlineResultListItem) UnmarshalJSON(data []byte) error {
	parsedData := []interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	if len(parsedData) != 5 {
		return errors.New("so far len(items) must be 5, please check it on documents")
	}
	*r = DerivativesIndexPriceKlineResultListItem{
		Start: parsedData[0].(string),
		Open:  parsedData[1].(string),
		High:  parsedData[2].(string),
		Low:   parsedData[3].(string),
		Close: parsedData[4].(string),
	}
	return nil
}

// DerivativesIndexPriceKlineParam :
type DerivativesIndexPriceKlineParam struct {
	Category CategoryDerivative `url:"category"`
	Symbol   SymbolDerivative   `url:"symbol"`
	Interval Interval           `url:"interval"`
	Start    int                `url:"start"` // timestamp point for result, in milliseconds
	End      int                `url:"end"`   // timestamp point for result, in milliseconds

	Limit *int `url:"limit,omitempty"`
}

// DerivativesIndexPriceKline :
func (s *DerivativeCommonService) DerivativesIndexPriceKline(param DerivativesIndexPriceKlineParam) (*DerivativesIndexPriceKlineResponse, error) {
	var res DerivativesIndexPriceKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/derivatives/v3/public/index-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
