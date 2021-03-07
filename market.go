package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// MarketService :
type MarketService struct {
	Client *Client
}

// OrderBookResponse :
type OrderBookResponse struct {
	CommonResponse `json:",inline"`
	Result         []OrderBookResult `json:"result"`
}

// UnmarshalJSON :
func (r *OrderBookResponse) UnmarshalJSON(data []byte) error {
	fmt.Println(string(data))
	return nil
}

// OrderBookResult :
type OrderBookResult struct {
	Symbol Symbol  `json:"symbol"`
	Price  float64 `json:"price"`
	Size   float64 `json:"size"`
	Side   Side    `json:"side"`
}

// OrderBook :
func (s *MarketService) OrderBook(symbol Symbol) (*OrderBookResponse, error) {
	var res OrderBookResponse

	params := map[string]string{
		"symbol": string(symbol),
	}

	url, err := s.Client.BuildPublicURL("/v2/public/orderBook/L2", params)
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

// ListKlineParam :
type ListKlineParam struct {
	Symbol   Symbol   `json:"symbol"`
	Interval Interval `json:"interval"`
	From     int      `json:"from"`

	Limit *int `json:"limit"`
}

func (p *ListKlineParam) build() map[string]string {
	result := map[string]string{
		"symbol":   string(p.Symbol),
		"interval": string(p.Interval),
		"from":     strconv.Itoa(p.From),
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	return result
}

// ListKlineResponse :
type ListKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListKlineResult `json:"result"`
}

// ListKlineResult :
type ListKlineResult struct {
	Symbol   Symbol `json:"symbol"`
	Interval string `json:"interval"`
	OpenTime int    `json:"open_time`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
	Turnover string `json:"turnover"`
}

// ListKline :
func (s *MarketService) ListKline(param ListKlineParam) (*ListKlineResponse, error) {
	var res ListKlineResponse

	url, err := s.Client.BuildPublicURL("/v2/public/kline/list", param.build())
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
