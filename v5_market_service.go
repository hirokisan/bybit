package bybit

import (
	"encoding/json"
	"errors"

	"github.com/google/go-querystring/query"
)

// V5MarketServiceI :
type V5MarketServiceI interface {
	GetKline(V5GetKlineParam) (*V5GetKlineResponse, error)
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
