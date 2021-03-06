package bybit

import (
	"encoding/json"
	"net/http"
)

// WalletService :
type WalletService struct {
	Client *Client
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
	Equity           int     `json:"equity"`
	AvailableBalance int     `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   int     `json:"position_margin"`
	OccClosingFee    int     `json:"occ_closing_fee"`
	OccFundingFee    int     `json:"occ_funding_fee"`
	WalletBalance    int     `json:"wallet_balance"`
	RealisedPnl      int     `json:"realised_pnl"`
	UnrealisedPnl    int     `json:"unrealised_pnl"`
	CumRealisedPnl   int     `json:"cum_realised_pnl"`
	GivenCash        int     `json:"given_cash"`
	ServiceCash      int     `json:"service_cash"`
}

// Balance :
func (s *WalletService) Balance(coin Coin) (*BalanceResponse, error) {
	var res BalanceResponse

	params := map[string]string{
		"coin": string(coin),
	}
	url := s.Client.BuildURL("/v2/private/wallet/balance", params)
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
