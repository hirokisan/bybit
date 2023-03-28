package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/go-querystring/query"
)

// V5MarketServiceI :
type V5MarketServiceI interface {
	GetKline(V5GetKlineParam) (*V5GetKlineResponse, error)
	GetMarkPriceKline(V5GetMarkPriceKlineParam) (*V5GetMarkPriceKlineResponse, error)
	GetIndexPriceKline(V5GetIndexPriceKlineParam) (*V5GetIndexPriceKlineResponse, error)
	GetPremiumIndexPriceKline(V5GetPremiumIndexPriceKlineParam) (*V5GetPremiumIndexPriceKlineResponse, error)
	GetInstrumentsInfo(V5GetInstrumentsInfoParam) (*V5GetInstrumentsInfoResponse, error)
	GetOrderbook(V5GetOrderbookParam) (*V5GetOrderbookResponse, error)
	GetTickers(V5GetTickersParam) (*V5GetTickersResponse, error)
	GetFundingRateHistory(V5GetFundingRateHistoryParam) (*V5GetFundingRateHistoryResponse, error)
	GetPublicTradingHistory(V5GetPublicTradingHistoryParam) (*V5GetPublicTradingHistoryResponse, error)
	GetOpenInterest(V5GetOpenInterestParam) (*V5GetOpenInterestResponse, error)
	GetHistoricalVolatility(V5GetHistoricalVolatilityParam) (*V5GetHistoricalVolatilityResponse, error)
}

// V5MarketService :
type V5MarketService struct {
	client *Client
}

// V5GetKlineParam :
type V5GetKlineParam struct {
	Category CategoryV5 `url:"category"`
	Symbol   SymbolV5   `url:"symbol"`
	Interval Interval   `url:"interval"`
	Start    *int       `url:"start,omitempty"` // timestamp point for result, in milliseconds
	End      *int       `url:"end,omitempty"`   // timestamp point for result, in milliseconds

	Limit *int `url:"limit,omitempty"` // Limit for data size per page. [1, 200]. Default: 200
}

// V5GetKlineResponse :
type V5GetKlineResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetKlineResult `json:"result"`
}

// V5GetKlineResult :
type V5GetKlineResult struct {
	Category CategoryV5     `json:"category"`
	Symbol   SymbolV5       `json:"symbol"`
	List     V5GetKlineList `json:"list"`
}

// V5GetKlineList :
type V5GetKlineList []V5GetKlineItem

// V5GetKlineItem :
type V5GetKlineItem struct {
	StartTime string
	Open      string
	High      string
	Low       string
	Close     string
	Volume    string
	Turnover  string
}

// UnmarshalJSON :
func (l *V5GetKlineList) UnmarshalJSON(data []byte) error {
	parsedData := [][]interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	for _, d := range parsedData {
		if len(d) != 7 {
			return errors.New("so far len(items) must be 7, please check it on documents")
		}
		*l = append(*l, V5GetKlineItem{
			StartTime: d[0].(string),
			Open:      d[1].(string),
			High:      d[2].(string),
			Low:       d[3].(string),
			Close:     d[4].(string),
			Volume:    d[5].(string),
			Turnover:  d[5].(string),
		})
	}
	return nil
}

// GetKline :
func (s *V5MarketService) GetKline(param V5GetKlineParam) (*V5GetKlineResponse, error) {
	var res V5GetKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetMarkPriceKlineParam :
type V5GetMarkPriceKlineParam struct {
	Category CategoryV5 `url:"category"`
	Symbol   SymbolV5   `url:"symbol"`
	Interval Interval   `url:"interval"`
	Start    *int       `url:"start,omitempty"` // timestamp point for result, in milliseconds
	End      *int       `url:"end,omitempty"`   // timestamp point for result, in milliseconds
	Limit    *int       `url:"limit,omitempty"` // Limit for data size per page. [1, 200]. Default: 200
}

// V5GetMarkPriceKlineResponse :
type V5GetMarkPriceKlineResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetMarkPriceKlineResult `json:"result"`
}

// V5GetMarkPriceKlineResult :
type V5GetMarkPriceKlineResult struct {
	Category CategoryV5              `json:"category"`
	Symbol   SymbolV5                `json:"symbol"`
	List     V5GetMarkPriceKlineList `json:"list"`
}

// V5GetMarkPriceKlineList :
type V5GetMarkPriceKlineList []V5GetMarkPriceKlineItem

// V5GetMarkPriceKlineItem :
type V5GetMarkPriceKlineItem struct {
	StartTime string
	Open      string
	High      string
	Low       string
	Close     string
}

// UnmarshalJSON :
func (l *V5GetMarkPriceKlineList) UnmarshalJSON(data []byte) error {
	parsedData := [][]interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	for _, d := range parsedData {
		if len(d) != 5 {
			return errors.New("so far len(items) must be 5, please check it on documents")
		}
		*l = append(*l, V5GetMarkPriceKlineItem{
			StartTime: d[0].(string),
			Open:      d[1].(string),
			High:      d[2].(string),
			Low:       d[3].(string),
			Close:     d[4].(string),
		})
	}
	return nil
}

// GetMarkPriceKline :
func (s *V5MarketService) GetMarkPriceKline(param V5GetMarkPriceKlineParam) (*V5GetMarkPriceKlineResponse, error) {
	var res V5GetMarkPriceKlineResponse

	if param.Category != CategoryV5Linear && param.Category != CategoryV5Inverse {
		return nil, fmt.Errorf("category should be linear or inverse")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/mark-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetIndexPriceKlineParam :
type V5GetIndexPriceKlineParam struct {
	Category CategoryV5 `url:"category"`
	Symbol   SymbolV5   `url:"symbol"`
	Interval Interval   `url:"interval"`
	Start    *int       `url:"start,omitempty"` // timestamp point for result, in milliseconds
	End      *int       `url:"end,omitempty"`   // timestamp point for result, in milliseconds
	Limit    *int       `url:"limit,omitempty"` // Limit for data size per page. [1, 200]. Default: 200
}

// V5GetIndexPriceKlineResponse :
type V5GetIndexPriceKlineResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetIndexPriceKlineResult `json:"result"`
}

// V5GetIndexPriceKlineResult :
type V5GetIndexPriceKlineResult struct {
	Category CategoryV5               `json:"category"`
	Symbol   SymbolV5                 `json:"symbol"`
	List     V5GetIndexPriceKlineList `json:"list"`
}

// V5GetIndexPriceKlineList :
type V5GetIndexPriceKlineList []V5GetIndexPriceKlineItem

// V5GetIndexPriceKlineItem :
type V5GetIndexPriceKlineItem struct {
	StartTime string
	Open      string
	High      string
	Low       string
	Close     string
}

// UnmarshalJSON :
func (l *V5GetIndexPriceKlineList) UnmarshalJSON(data []byte) error {
	parsedData := [][]interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	for _, d := range parsedData {
		if len(d) != 5 {
			return errors.New("so far len(items) must be 5, please check it on documents")
		}
		*l = append(*l, V5GetIndexPriceKlineItem{
			StartTime: d[0].(string),
			Open:      d[1].(string),
			High:      d[2].(string),
			Low:       d[3].(string),
			Close:     d[4].(string),
		})
	}
	return nil
}

// GetIndexPriceKline :
func (s *V5MarketService) GetIndexPriceKline(param V5GetIndexPriceKlineParam) (*V5GetIndexPriceKlineResponse, error) {
	var res V5GetIndexPriceKlineResponse

	if param.Category != CategoryV5Linear && param.Category != CategoryV5Inverse {
		return nil, fmt.Errorf("category should be linear or inverse")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/index-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetPremiumIndexPriceKlineParam :
type V5GetPremiumIndexPriceKlineParam struct {
	Category CategoryV5 `url:"category"`
	Symbol   SymbolV5   `url:"symbol"`
	Interval Interval   `url:"interval"`
	Start    *int       `url:"start,omitempty"` // timestamp point for result, in milliseconds
	End      *int       `url:"end,omitempty"`   // timestamp point for result, in milliseconds
	Limit    *int       `url:"limit,omitempty"` // Limit for data size per page. [1, 200]. Default: 200
}

// V5GetPremiumIndexPriceKlineResponse :
type V5GetPremiumIndexPriceKlineResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetPremiumIndexPriceKlineResult `json:"result"`
}

// V5GetPremiumIndexPriceKlineResult :
type V5GetPremiumIndexPriceKlineResult struct {
	Category CategoryV5                      `json:"category"`
	Symbol   SymbolV5                        `json:"symbol"`
	List     V5GetPremiumIndexPriceKlineList `json:"list"`
}

// V5GetPremiumIndexPriceKlineList :
type V5GetPremiumIndexPriceKlineList []V5GetPremiumIndexPriceKlineItem

// V5GetPremiumIndexPriceKlineItem :
type V5GetPremiumIndexPriceKlineItem struct {
	StartTime string
	Open      string
	High      string
	Low       string
	Close     string
}

// UnmarshalJSON :
func (l *V5GetPremiumIndexPriceKlineList) UnmarshalJSON(data []byte) error {
	parsedData := [][]interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	for _, d := range parsedData {
		if len(d) != 5 {
			return errors.New("so far len(items) must be 5, please check it on documents")
		}
		log.Println(d)
		*l = append(*l, V5GetPremiumIndexPriceKlineItem{
			StartTime: d[0].(string),
			Open:      d[1].(string),
			High:      d[2].(string),
			Low:       d[3].(string),
			Close:     d[4].(string),
		})
	}
	return nil
}

// GetPremiumIndexPriceKline :
func (s *V5MarketService) GetPremiumIndexPriceKline(param V5GetPremiumIndexPriceKlineParam) (*V5GetPremiumIndexPriceKlineResponse, error) {
	var res V5GetPremiumIndexPriceKlineResponse

	if param.Category != CategoryV5Linear {
		return nil, fmt.Errorf("category should be linear")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/premium-index-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetInstrumentsInfoParam :
// Spot does not support pagination, so limit, cursor are invalid.
// When query by baseCoin, regardless of category=linear or inverse, the result will have Linear contract and Inverse contract symbols.
type V5GetInstrumentsInfoParam struct {
	Category CategoryV5 `url:"category"`

	Symbol   *SymbolV5 `url:"symbol,omitempty"`
	BaseCoin *Coin     `url:"baseCoin,omitempty"` // Base coin. linear,inverse,option only
	Limit    *int      `url:"limit,omitempty"`    // Limit for data size per page. [1, 1000]. Default: 500
	Cursor   *string   `url:"cursor,omitempty"`
}

// V5GetInstrumentsInfoResponse :
type V5GetInstrumentsInfoResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetInstrumentsInfoResult `json:"result"`
}

// V5GetInstrumentsInfoResult :
// Responses are filled according to category.
type V5GetInstrumentsInfoResult struct {
	LinearInverse *V5GetInstrumentsInfoLinearInverseResult
	Option        *V5GetInstrumentsInfoOptionResult
	Spot          *V5GetInstrumentsInfoSpotResult
}

// UnmarshalJSON :
func (r *V5GetInstrumentsInfoResult) UnmarshalJSON(data []byte) error {
	var categoryJudge struct {
		Category CategoryV5 `json:"category"`
	}
	if err := json.Unmarshal(data, &categoryJudge); err != nil {
		return err
	}
	switch categoryJudge.Category {
	case CategoryV5Linear, CategoryV5Inverse:
		if err := json.Unmarshal(data, &r.LinearInverse); err != nil {
			return err
		}
	case CategoryV5Option:
		if err := json.Unmarshal(data, &r.Option); err != nil {
			return err
		}
	case CategoryV5Spot:
		if err := json.Unmarshal(data, &r.Spot); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unexpected category %s given", categoryJudge.Category)
	}
	return nil
}

// V5GetInstrumentsInfoLinearInverseResult :
type V5GetInstrumentsInfoLinearInverseResult struct {
	Category       CategoryV5 `json:"category"`
	NextPageCursor string     `json:"nextPageCursor"`
	List           []struct {
		Symbol          SymbolV5         `json:"symbol"`
		ContractType    ContractType     `json:"contractType"`
		Status          InstrumentStatus `json:"status"`
		BaseCoin        Coin             `json:"baseCoin"`
		QuoteCoin       Coin             `json:"quoteCoin"`
		SettleCoin      Coin             `json:"settleCoin"`
		LaunchTime      string           `json:"launchTime"`
		DeliveryTime    string           `json:"deliveryTime"`
		DeliveryFeeRate string           `json:"deliveryFeeRate"`
		PriceScale      string           `json:"priceScale"`
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
			MaxOrderQty         string `json:"maxOrderQty"`
			MinOrderQty         string `json:"minOrderQty"`
			QtyStep             string `json:"qtyStep"`
			PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
		} `json:"lotSizeFilter"`
		UnifiedMarginTrade bool `json:"unifiedMarginTrade"`
		FundingInterval    int  `json:"fundingInterval"`
	} `json:"list"`
}

// V5GetInstrumentsInfoOptionResult :
type V5GetInstrumentsInfoOptionResult struct {
	Category       CategoryV5 `json:"category"`
	NextPageCursor string     `json:"nextPageCursor"`
	List           []struct {
		Symbol          SymbolV5         `json:"symbol"`
		OptionsType     OptionsType      `json:"optionsType"`
		Status          InstrumentStatus `json:"status"`
		BaseCoin        Coin             `json:"baseCoin"`
		QuoteCoin       Coin             `json:"quoteCoin"`
		SettleCoin      Coin             `json:"settleCoin"`
		LaunchTime      string           `json:"launchTime"`
		DeliveryTime    string           `json:"deliveryTime"`
		DeliveryFeeRate string           `json:"deliveryFeeRate"`
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
	} `json:"list"`
}

// V5GetInstrumentsInfoSpotResult :
type V5GetInstrumentsInfoSpotResult struct {
	Category CategoryV5 `json:"category"`
	List     []struct {
		Symbol        SymbolV5         `json:"symbol"`
		BaseCoin      Coin             `json:"baseCoin"`
		QuoteCoin     Coin             `json:"quoteCoin"`
		Innovation    Innovation       `json:"innovation"`
		Status        InstrumentStatus `json:"status"`
		LotSizeFilter struct {
			BasePrecision  string `json:"basePrecision"`
			QuotePrecision string `json:"quotePrecision"`
			MaxOrderQty    string `json:"maxOrderQty"`
			MinOrderQty    string `json:"minOrderQty"`
			MinOrderAmt    string `json:"minOrderAmt"`
			MaxOrderAmt    string `json:"maxOrderAmt"`
		} `json:"lotSizeFilter"`
		PriceFilter struct {
			TickSize string `json:"tickSize"`
		} `json:"priceFilter"`
	} `json:"list"`
}

// GetInstrumentsInfo :
func (s *V5MarketService) GetInstrumentsInfo(param V5GetInstrumentsInfoParam) (*V5GetInstrumentsInfoResponse, error) {
	var res V5GetInstrumentsInfoResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/instruments-info", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetOrderbookParam :
type V5GetOrderbookParam struct {
	Category CategoryV5 `url:"category"`
	Symbol   SymbolV5   `url:"symbol"`

	// spot: [1, 50]. Default: 1.
	// linear&inverse: [1, 200]. Default: 25.
	// option: [1, 25]. Default: 1.
	Limit *int `url:"limit,omitempty"`
}

// V5GetOrderbookResponse :
type V5GetOrderbookResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetOrderbookResult `json:"result"`
}

// V5GetOrderbookResult :
type V5GetOrderbookResult struct {
	Symbol    SymbolV5              `json:"s"`
	Bids      V5GetOrderbookBidAsks `json:"b"`
	Asks      V5GetOrderbookBidAsks `json:"a"`
	Timestamp int64                 `json:"ts"`
	UpdateID  int                   `json:"u"`
}

// V5GetOrderbookBidAsks :
type V5GetOrderbookBidAsks []V5GetOrderbookBidAsk

// UnmarshalJSON :
func (b *V5GetOrderbookBidAsks) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	items := V5GetOrderbookBidAsks{}
	for _, item := range parsedData {
		item := item
		if len(item) != 2 {
			return errors.New("so far len(item) must be 2, please check it on documents")
		}
		items = append(items, V5GetOrderbookBidAsk{
			Price:    item[0],
			Quantity: item[1],
		})
	}
	*b = items
	return nil
}

// V5GetOrderbookBidAsk :
type V5GetOrderbookBidAsk struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

// GetOrderbook :
func (s *V5MarketService) GetOrderbook(param V5GetOrderbookParam) (*V5GetOrderbookResponse, error) {
	var res V5GetOrderbookResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/orderbook", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetTickersParam :
type V5GetTickersParam struct {
	Category CategoryV5 `url:"category"`

	Symbol   *SymbolV5 `url:"symbol,omitempty"`
	BaseCoin *Coin     `url:"baseCoin,omitempty"` // Base coin. For option only
	ExpDate  *string   `url:"expDate,omitempty"`  // Expiry date. e.g., 25DEC22. For option only
}

func (p V5GetTickersParam) validate() error {
	if p.Category == CategoryV5Option && (p.Symbol == nil && p.BaseCoin == nil) {
		return fmt.Errorf("symbol or baseCoin must be passed for option")
	}
	if p.BaseCoin != nil && p.Category != CategoryV5Option {
		return fmt.Errorf("baseCoin is for option only")
	}
	if p.ExpDate != nil && p.Category != CategoryV5Option {
		return fmt.Errorf("expDate is for option only")
	}
	return nil
}

// V5GetTickersResponse :
type V5GetTickersResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetTickersResult `json:"result"`
}

// V5GetTickersResult :
// Responses are filled according to category.
type V5GetTickersResult struct {
	LinearInverse *V5GetTickersLinearInverseResult
	Option        *V5GetTickersOptionResult
	Spot          *V5GetTickersSpotResult
}

// UnmarshalJSON :
func (r *V5GetTickersResult) UnmarshalJSON(data []byte) error {
	var categoryJudge struct {
		Category CategoryV5 `json:"category"`
	}
	if err := json.Unmarshal(data, &categoryJudge); err != nil {
		return err
	}
	switch categoryJudge.Category {
	case CategoryV5Linear, CategoryV5Inverse:
		if err := json.Unmarshal(data, &r.LinearInverse); err != nil {
			return err
		}
	case CategoryV5Option:
		if err := json.Unmarshal(data, &r.Option); err != nil {
			return err
		}
	case CategoryV5Spot:
		if err := json.Unmarshal(data, &r.Spot); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unexpected category %s given", categoryJudge.Category)
	}
	return nil
}

// V5GetTickersLinearInverseResult :
type V5GetTickersLinearInverseResult struct {
	Category CategoryV5 `json:"category"`
	List     []struct {
		Symbol                 SymbolV5 `json:"symbol"`
		LastPrice              string   `json:"lastPrice"`
		IndexPrice             string   `json:"indexPrice"`
		MarkPrice              string   `json:"markPrice"`
		PrevPrice24H           string   `json:"prevPrice24h"`
		Price24HPcnt           string   `json:"price24hPcnt"`
		HighPrice24H           string   `json:"highPrice24h"`
		LowPrice24H            string   `json:"lowPrice24h"`
		PrevPrice1H            string   `json:"prevPrice1h"`
		OpenInterest           string   `json:"openInterest"`
		OpenInterestValue      string   `json:"openInterestValue"`
		Turnover24H            string   `json:"turnover24h"`
		Volume24H              string   `json:"volume24h"`
		FundingRate            string   `json:"fundingRate"`
		NextFundingTime        string   `json:"nextFundingTime"`
		PredictedDeliveryPrice string   `json:"predictedDeliveryPrice"`
		BasisRate              string   `json:"basisRate"`
		DeliveryFeeRate        string   `json:"deliveryFeeRate"`
		DeliveryTime           string   `json:"deliveryTime"`
		Ask1Size               string   `json:"ask1Size"`
		Bid1Price              string   `json:"bid1Price"`
		Ask1Price              string   `json:"ask1Price"`
		Bid1Size               string   `json:"bid1Size"`
	} `json:"list"`
}

// V5GetTickersOptionResult :
type V5GetTickersOptionResult struct {
	Category CategoryV5 `json:"category"`
	List     []struct {
		Symbol                 SymbolV5 `json:"symbol"`
		Bid1Price              string   `json:"bid1Price"`
		Bid1Size               string   `json:"bid1Size"`
		Bid1Iv                 string   `json:"bid1Iv"`
		Ask1Price              string   `json:"ask1Price"`
		Ask1Size               string   `json:"ask1Size"`
		Ask1Iv                 string   `json:"ask1Iv"`
		LastPrice              string   `json:"lastPrice"`
		HighPrice24H           string   `json:"highPrice24h"`
		LowPrice24H            string   `json:"lowPrice24h"`
		MarkPrice              string   `json:"markPrice"`
		IndexPrice             string   `json:"indexPrice"`
		MarkIv                 string   `json:"markIv"`
		UnderlyingPrice        string   `json:"underlyingPrice"`
		OpenInterest           string   `json:"openInterest"`
		Turnover24H            string   `json:"turnover24h"`
		Volume24H              string   `json:"volume24h"`
		TotalVolume            string   `json:"totalVolume"`
		TotalTurnover          string   `json:"totalTurnover"`
		Delta                  string   `json:"delta"`
		Gamma                  string   `json:"gamma"`
		Vega                   string   `json:"vega"`
		Theta                  string   `json:"theta"`
		PredictedDeliveryPrice string   `json:"predictedDeliveryPrice"`
		Change24H              string   `json:"change24h"`
	} `json:"list"`
}

// V5GetTickersSpotResult :
type V5GetTickersSpotResult struct {
	Category CategoryV5 `json:"category"`
	List     []struct {
		Symbol        SymbolV5 `json:"symbol"`
		Bid1Price     string   `json:"bid1Price"`
		Bid1Size      string   `json:"bid1Size"`
		Ask1Price     string   `json:"ask1Price"`
		Ask1Size      string   `json:"ask1Size"`
		LastPrice     string   `json:"lastPrice"`
		PrevPrice24H  string   `json:"prevPrice24h"`
		Price24HPcnt  string   `json:"price24hPcnt"`
		HighPrice24H  string   `json:"highPrice24h"`
		LowPrice24H   string   `json:"lowPrice24h"`
		Turnover24H   string   `json:"turnover24h"`
		Volume24H     string   `json:"volume24h"`
		UsdIndexPrice string   `json:"usdIndexPrice"`
	} `json:"list"`
}

// GetTickers :
func (s *V5MarketService) GetTickers(param V5GetTickersParam) (*V5GetTickersResponse, error) {
	var res V5GetTickersResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/tickers", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetFundingRateHistoryParam :
type V5GetFundingRateHistoryParam struct {
	Category CategoryV5 `url:"category"`
	Symbol   SymbolV5   `url:"symbol"`

	StartTime *int64 `url:"startTime,omitempty"` // The start timestamp (ms)
	EndTime   *int64 `url:"endTime,omitempty"`   // The start timestamp (ms)
	Limit     *int   `url:"limit,omitempty"`     // Limit for data size per page. [1, 200]. Default: 200
}

func (p V5GetFundingRateHistoryParam) validate() error {
	if p.Category != CategoryV5Linear && p.Category != CategoryV5Inverse {
		return fmt.Errorf("only linear and inverse are supported for category")
	}
	return nil
}

// V5GetFundingRateHistoryResponse :
type V5GetFundingRateHistoryResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetFundingRateHistoryResult `json:"result"`
}

// V5GetFundingRateHistoryResult :
type V5GetFundingRateHistoryResult struct {
	Category CategoryV5 `json:"category"`
	List     []struct {
		Symbol               SymbolV5 `json:"symbol"`
		FundingRate          string   `json:"fundingRate"`
		FundingRateTimestamp string   `json:"fundingRateTimestamp"`
	} `json:"list"`
}

// GetFundingRateHistory :
func (s *V5MarketService) GetFundingRateHistory(param V5GetFundingRateHistoryParam) (*V5GetFundingRateHistoryResponse, error) {
	var res V5GetFundingRateHistoryResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/funding/history", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetPublicTradingHistoryParam :
type V5GetPublicTradingHistoryParam struct {
	Category CategoryV5 `url:"category"`
	Symbol   SymbolV5   `url:"symbol"`

	BaseCoin   *Coin        `url:"baseCoin,omitempty"` // For option only. If not passed, return BTC data by default
	OptionType *OptionsType `url:"optionType,omitempty"`

	// Limit for data size per page.
	// - spot: [1,60], default: 60
	// - others: [1,1000], default: 500
	Limit *int `url:"limit,omitempty"`
}

func (p V5GetPublicTradingHistoryParam) validate() error {
	if p.BaseCoin != nil && p.Category != CategoryV5Option {
		return fmt.Errorf("baseCoin is for option only")
	}
	return nil
}

// V5GetPublicTradingHistoryResponse :
type V5GetPublicTradingHistoryResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetPublicTradingHistoryResult `json:"result"`
}

// V5GetPublicTradingHistoryResult :
type V5GetPublicTradingHistoryResult struct {
	Category CategoryV5 `json:"category"`
	List     []struct {
		ExecID       string   `json:"execId"`
		Symbol       SymbolV5 `json:"symbol"`
		Price        string   `json:"price"`
		Size         string   `json:"size"`
		Side         Side     `json:"side"`
		Time         string   `json:"time"`
		IsBlockTrade bool     `json:"isBlockTrade"`
	} `json:"list"`
}

// GetPublicTradingHistory :
func (s *V5MarketService) GetPublicTradingHistory(param V5GetPublicTradingHistoryParam) (*V5GetPublicTradingHistoryResponse, error) {
	var res V5GetPublicTradingHistoryResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/recent-trade", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetOpenInterestParam :
type V5GetOpenInterestParam struct {
	Category     CategoryV5 `url:"category"`
	Symbol       SymbolV5   `url:"symbol"`
	IntervalTime Period     `url:"intervalTime"`

	StartTime *int64  `url:"startTime,omitempty"` // The start timestamp (ms)
	EndTime   *int64  `url:"endTime,omitempty"`   // The start timestamp (ms)
	Limit     *int    `url:"limit,omitempty"`     // Limit for data size per page. [1, 200]. Default: 50
	Cursor    *string `url:"cursor,omitempty"`
}

func (p V5GetOpenInterestParam) validate() error {
	if p.Category != CategoryV5Linear && p.Category != CategoryV5Inverse {
		return fmt.Errorf("only linear and inverse are supported for category")
	}
	return nil
}

// V5GetOpenInterestResponse :
type V5GetOpenInterestResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetOpenInterestResult `json:"result"`
}

// V5GetOpenInterestResult :
type V5GetOpenInterestResult struct {
	Category CategoryV5 `json:"category"`
	Symbol   SymbolV5   `json:"symbol"`
	List     []struct {
		OpenInterest string `json:"openInterest"`
		Timestamp    string `json:"timestamp"`
	} `json:"list"`
	NextPageCursor string `json:"nextPageCursor"`
}

// GetOpenInterest :
func (s *V5MarketService) GetOpenInterest(param V5GetOpenInterestParam) (*V5GetOpenInterestResponse, error) {
	var res V5GetOpenInterestResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/open-interest", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetHistoricalVolatilityParam :
type V5GetHistoricalVolatilityParam struct {
	Category CategoryV5 `url:"category"` // option only

	BaseCoin  *Coin  `url:"baseCoin,omitempty"`
	Period    *int   `url:"period,omitempty"`
	StartTime *int64 `url:"startTime,omitempty"` // The start timestamp (ms)
	EndTime   *int64 `url:"endTime,omitempty"`   // The start timestamp (ms)
}

func (p V5GetHistoricalVolatilityParam) validate() error {
	if p.Category != CategoryV5Option {
		return fmt.Errorf("only option is supported")
	}
	return nil
}

// V5GetHistoricalVolatilityResponse :
type V5GetHistoricalVolatilityResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetHistoricalVolatilityResult `json:"result"`
}

// UnmarshalJSON : Because the response structure is different from others
func (r *V5GetHistoricalVolatilityResponse) UnmarshalJSON(data []byte) error {
	var tmp struct {
		CommonV5Response `json:",inline"`
		Category         CategoryV5                          `json:"category"`
		List             []V5GetHistoricalVolatilityListItem `json:"result"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = V5GetHistoricalVolatilityResponse{
		CommonV5Response: tmp.CommonV5Response,
		Result: V5GetHistoricalVolatilityResult{
			Category: tmp.Category,
			List:     tmp.List,
		},
	}

	return nil
}

// V5GetHistoricalVolatilityResult :
type V5GetHistoricalVolatilityResult struct {
	Category CategoryV5                          `json:"category"`
	List     []V5GetHistoricalVolatilityListItem `json:"list"`
}

// V5GetHistoricalVolatilityListItem :
type V5GetHistoricalVolatilityListItem struct {
	Period int    `json:"period"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

// GetHistoricalVolatility :
func (s *V5MarketService) GetHistoricalVolatility(param V5GetHistoricalVolatilityParam) (*V5GetHistoricalVolatilityResponse, error) {
	var res V5GetHistoricalVolatilityResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/market/historical-volatility", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
