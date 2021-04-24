package bybit

import (
	"encoding/json"
	"net/http"
)

// LinearTickers :
func (s *MarketService) LinearTickers(symbol SymbolUSDT) (*TickersResponse, error) {
	var res TickersResponse

	params := map[string]string{
		"symbol": string(symbol),
	}

	url, err := s.Client.BuildPublicURL("/v2/public/tickers", params)
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
