package bybit

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// V5AccountServiceI :
type V5AccountServiceI interface {
	GetWalletBalance(AccountTypeV5, []Coin) (*V5GetWalletBalanceResponse, error)
	SetCollateralCoin(V5SetCollateralCoinParam) (*V5SetCollateralCoinResponse, error)
	GetCollateralInfo(V5GetCollateralInfoParam) (*V5GetCollateralInfoResponse, error)
	GetAccountInfo() (*V5GetAccountInfoResponse, error)
	GetTransactionLog(V5GetTransactionLogParam) (*V5GetTransactionLogResponse, error)
}

// V5AccountService :
type V5AccountService struct {
	client *Client
}

// V5GetWalletBalanceResponse :
type V5GetWalletBalanceResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5WalletBalanceResult `json:"result"`
}

// V5WalletBalanceResult :
type V5WalletBalanceResult struct {
	List []V5WalletBalanceList `json:"list"`
}

// V5WalletBalanceCoin :
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
	Free                string `json:"free"`
	Locked              string `json:"locked"`
	Coin                Coin   `json:"coin"`
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

// GetWalletBalance :
//
// at: UNIFIED, CONTRACT, SPOT
//
// coin:
// If not passed, it returns non-zero asset info
// You can pass multiple coins to query, separated by comma. "USDT,USDC".
func (s *V5AccountService) GetWalletBalance(at AccountTypeV5, coins []Coin) (*V5GetWalletBalanceResponse, error) {
	switch at {
	case AccountTypeV5UNIFIED, AccountTypeV5CONTRACT, AccountTypeV5SPOT:
	default:
		return nil, fmt.Errorf("wrong account type")
	}
	var (
		res   V5GetWalletBalanceResponse
		query = make(url.Values)
	)

	query.Add("accountType", string(at))
	if len(coins) > 0 {
		var coinsStr []string
		for _, c := range coins {
			coinsStr = append(coinsStr, string(c))
		}
		query.Add("coin", strings.Join(coinsStr, ","))
	}

	if err := s.client.getV5Privately("/v5/account/wallet-balance", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5SetCollateralCoinParam :
type V5SetCollateralCoinParam struct {
	// Coin:
	// You cannot pass multiple coins to query
	// USDT,USDC cannot be switched off
	Coin Coin `json:"coin"`

	// CollateralSwitch: CollateralSwitchV5On or CollateralSwitchV5Off
	CollateralSwitch CollateralSwitchV5 `json:"collateralSwitch"`
}

// V5SetCollateralCoinResponse :
type V5SetCollateralCoinResponse struct {
	CommonV5Response `json:",inline"`
	Result           interface{} `json:"result"`
}

// SetCollateralCoin :
func (s *V5AccountService) SetCollateralCoin(param V5SetCollateralCoinParam) (*V5SetCollateralCoinResponse, error) {
	var res V5SetCollateralCoinResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postV5JSON("/v5/account/set-collateral-switch", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetCollateralInfoParam :
type V5GetCollateralInfoParam struct {
	Currency *string `url:"currency,omitempty"`
}

// V5GetCollateralInfoResponse :
type V5GetCollateralInfoResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetCollateralInfoResult
}

// V5GetCollateralInfoResult :
type V5GetCollateralInfoResult struct {
	List []V5GetCollateralInfoList `json:"list"`
}

// V5GetCollateralInfoList :
type V5GetCollateralInfoList struct {
	Currency            string `json:"currency"`
	HourlyBorrowRate    string `json:"hourlyBorrowRate"`
	MaxBorrowingAmount  string `json:"maxBorrowingAmount"`
	FreeBorrowingLimit  string `json:"freeBorrowingLimit"`
	FreeBorrowAmount    string `json:"freeBorrowAmount"`
	BorrowAmount        string `json:"borrowAmount"`
	FreeBorrowingAmount string `json:"freeBorrowingAmount"`
	AvailableToBorrow   string `json:"availableToBorrow"`
	Borrowable          bool   `json:"borrowable"`
	BorrowUsageRate     string `json:"borrowUsageRate"`
	MarginCollateral    bool   `json:"marginCollateral"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	CollateralRatio     string `json:"collateralRatio"`
}

// GetCollateralInfo :
func (s *V5AccountService) GetCollateralInfo(param V5GetCollateralInfoParam) (*V5GetCollateralInfoResponse, error) {
	var res V5GetCollateralInfoResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err = s.client.getV5Privately("/v5/account/collateral-info", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetAccountInfoResponse :
type V5GetAccountInfoResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5AccountInfoResult `json:"result"`
}

// V5AccountInfoResult :
type V5AccountInfoResult struct {
	MarginMode          MarginMode          `json:"marginMode"`
	UpdatedTime         string              `json:"updatedTime"`
	UnifiedMarginStatus UnifiedMarginStatus `json:"unifiedMarginStatus"`
}

// GetAccountInfo :
func (s *V5AccountService) GetAccountInfo() (*V5GetAccountInfoResponse, error) {
	var (
		res   V5GetAccountInfoResponse
		query = make(url.Values)
	)

	if err := s.client.getV5Privately("/v5/account/info", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetTransactionLogParam :
type V5GetTransactionLogParam struct {
	AccountType *AccountTypeV5        `url:"accountType,omitempty"`
	Category    *CategoryV5           `url:"category,omitempty"`
	Currency    *string               `url:"currency,omitempty"`
	BaseCoin    *Coin                 `url:"baseCoin,omitempty"`
	Type        *TransactionLogTypeV5 `url:"type,omitempty"`
	StartTime   *int64                `url:"startTime,omitempty"` // The start timestamp (ms)
	EndTime     *int64                `url:"endTime,omitempty"`   // The start timestamp (ms)
	Limit       *int                  `url:"limit,omitempty"`     // Limit for data size per page. [1, 50]. Default: 20
	Cursor      *string               `url:"cursor,omitempty"`
}

// V5GetTransactionLogResponse :
type V5GetTransactionLogResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetTransactionLogResult `json:"result"`
}

// V5GetTransactionLogResult :
type V5GetTransactionLogResult struct {
	NextPageCursor string                  `json:"nextPageCursor"`
	List           V5GetTransactionLogList `json:"list"`
}

// V5GetTransactionLogList :
type V5GetTransactionLogList []V5GetTransactionLogItem

// V5GetTransactionLogItem :
type V5GetTransactionLogItem struct {
	Symbol          SymbolV5             `json:"symbol"`
	Category        CategoryV5           `json:"category"`
	Side            Side                 `json:"side"`
	TransactionTime string               `json:"transactionTime"`
	Type            TransactionLogTypeV5 `json:"type"`
	Qty             string               `json:"qty"`
	Size            string               `json:"size"`
	Currency        string               `json:"currency"`
	TradePrice      string               `json:"tradePrice"`
	Funding         string               `json:"funding"`
	Fee             string               `json:"fee"`
	CashFlow        string               `json:"cashFlow"`
	Change          string               `json:"change"`
	CashBalance     string               `json:"cashBalance"`
	FeeRate         string               `json:"feeRate"`
	BonusChange     string               `json:"bonusChange"`
	TradeID         string               `json:"tradeId"`
	OrderID         string               `json:"orderId"`
	OrderLinkID     string               `json:"orderLinkId"`
}

// GetTransactionLog :
func (s *V5AccountService) GetTransactionLog(param V5GetTransactionLogParam) (*V5GetTransactionLogResponse, error) {
	var res V5GetTransactionLogResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/account/transaction-log", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
