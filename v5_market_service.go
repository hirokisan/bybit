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
