package bybit

import (
	"net/url"
)

// V5AccountServiceI :
type V5AccountServiceI interface {
	GetWalletBalance(AccountType, string) (*V5WalletBalanceResponse, error)
}

// V5AccountService :
type V5AccountService struct {
	client *Client
}

// V5WalletBalanceResponse :
type V5WalletBalanceResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5WalletBalanceResult `json:"result"`
}

// V5WalletBalanceResult :
type V5WalletBalanceResult struct {
	List []V5WalletBalanceList `json:"list"`
}

// VV5WalletBalanceCoin :
type V5WalletBalanceCoin struct {
	AvailableToBorrow   string `json:"availableToBorrow"`
	AccruedInterest     string `json:"accruedInterest"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	TotalOrderIM        string `json:"totalOrderIM"`
	Equity              string `json:"equity"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UsdValue            string `json:"usdValue"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	BorrowAmount        string `json:"borrowAmount"`
	TotalPositionIM     string `json:"totalPositionIM"`
	WalletBalance       string `json:"walletBalance"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Coin                string `json:"coin"`
}

// V5WalletBalanceList :
type V5WalletBalanceList struct {
	TotalEquity            string                `json:"totalEquity"`
	AccountIMRate          string                `json:"accountIMRate"`
	TotalMarginBalance     string                `json:"totalMarginBalance"`
	TotalInitialMargin     string                `json:"totalInitialMargin"`
	AccountType            string                `json:"accountType"`
	TotalAvailableBalance  string                `json:"totalAvailableBalance"`
	AccountMMRate          string                `json:"accountMMRate"`
	TotalPerpUPL           string                `json:"totalPerpUPL"`
	TotalWalletBalance     string                `json:"totalWalletBalance"`
	TotalMaintenanceMargin string                `json:"totalMaintenanceMargin"`
	Coin                   []V5WalletBalanceCoin `json:"coin"`
}

// GetPositionInfo :
//
// at: UNIFIED or CONTRACT
//
// coin:
// If not passed, it returns non-zero asset info
// You can pass multiple coins to query, separated by comma. "USDT,USDC".
func (s *V5AccountService) GetWalletBalance(at AccountType, coin string) (*V5WalletBalanceResponse, error) {
	var (
		res   V5WalletBalanceResponse
		query = make(url.Values)
	)

	query.Add("accountType", string(at))
	if coin != "" {
		query.Add("coin", coin)
	}

	if err := s.client.getV5Privately("/v5/account/wallet-balance", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
