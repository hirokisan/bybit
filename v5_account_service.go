package bybit

import (
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// V5AccountServiceI :
type V5AccountServiceI interface {
	GetWalletBalance(AccountType, []Coin) (*V5GetWalletBalanceResponse, error)
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
// at: UNIFIED or CONTRACT
//
// coin:
// If not passed, it returns non-zero asset info
// You can pass multiple coins to query, separated by comma. "USDT,USDC".
func (s *V5AccountService) GetWalletBalance(at AccountType, coins []Coin) (*V5GetWalletBalanceResponse, error) {
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
