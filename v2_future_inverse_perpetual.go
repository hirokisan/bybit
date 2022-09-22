package bybit

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

// FutureInversePerpetualService :
type FutureInversePerpetualService struct {
	client *Client

	*FutureCommonService
}

// BalanceResponse :
type BalanceResponse struct {
	CommonResponse `json:",inline"`
	Result         BalanceResult `json:"result"`
}

// BalanceResult :
type BalanceResult struct {
	Balance map[Coin]Balance
}

// UnmarshalJSON :
func (r *BalanceResult) UnmarshalJSON(data []byte) error {
	parsedData := map[string]Balance{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	r.Balance = map[Coin]Balance{}
	for coin, balanceData := range parsedData {
		r.Balance[Coin(coin)] = balanceData
	}
	return nil
}

// Balance :
type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

// Balance :
func (s *FutureInversePerpetualService) Balance(coin Coin) (*BalanceResponse, error) {
	var res BalanceResponse

	query := url.Values{}
	query.Add("coin", string(coin))
	if err := s.client.getPrivately("/v2/private/wallet/balance", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// PremiumIndexKlineResponse :
type PremiumIndexKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []PremiumIndexKlineResult `json:"result"`
}

// PremiumIndexKlineResult :
type PremiumIndexKlineResult struct {
	Symbol   SymbolInverse `json:"symbol"`
	Period   Period        `json:"period"`
	OpenTime int           `json:"open_time"`
	Open     string        `json:"open"`
	High     string        `json:"high"`
	Low      string        `json:"low"`
	Close    string        `json:"close"`
}

// PremiumIndexKlineParam :
type PremiumIndexKlineParam struct {
	Symbol   SymbolInverse `url:"symbol"`
	Interval Interval      `url:"interval"`
	From     int           `url:"from"`

	Limit *int `url:"limit,omitempty"`
}

// PremiumIndexKline :
func (s *FutureInversePerpetualService) PremiumIndexKline(param PremiumIndexKlineParam) (*PremiumIndexKlineResponse, error) {
	var res PremiumIndexKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/premium-index-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
