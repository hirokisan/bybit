package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
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
