package bybit

import (
	"encoding/json"

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

	Limit int `url:"limit,omitempty"`
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
