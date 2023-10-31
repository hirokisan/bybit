package bybit

import "github.com/google/go-querystring/query"

// V5AssetServiceI :
type V5AssetServiceI interface {
	GetInternalTransferRecords(V5GetInternalTransferRecordsParam) (*V5GetInternalTransferRecordsResponse, error)
	GetDepositRecords(V5GetDepositRecordsParam) (*V5GetDepositRecordsResponse, error)
	GetSubDepositRecords(V5GetSubDepositRecordsParam) (*V5GetSubDepositRecordsResponse, error)
	GetInternalDepositRecords(V5GetInternalDepositRecordsParam) (*V5GetInternalDepositRecordsResponse, error)
	GetWithdrawalRecords(V5GetWithdrawalRecordsParam) (*V5GetWithdrawalRecordsResponse, error)
	GetCoinInfo(V5GetCoinInfoParam) (*V5GetCoinInfoResponse, error)
}

// V5AssetService :
type V5AssetService struct {
	client *Client
}

// V5GetInternalTransferRecordsParam :
type V5GetInternalTransferRecordsParam struct {
	TransferID *string           `url:"transferId,omitempty"`
	Coin       *Coin             `url:"coin,omitempty"`
	Status     *TransferStatusV5 `url:"status,omitempty"`
	StartTime  *int64            `url:"startTime,omitempty"` // The start timestamp (ms)
	EndTime    *int64            `url:"endTime,omitempty"`   // The start timestamp (ms)
	Limit      *int              `url:"limit,omitempty"`     // Limit for data size per page. [1, 50]. Default: 20
	Cursor     *string           `url:"cursor,omitempty"`
}

// V5GetInternalTransferRecordsResponse :
type V5GetInternalTransferRecordsResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetInternalTransferRecordsResult `json:"result"`
}

// V5GetInternalTransferRecordsResult :
type V5GetInternalTransferRecordsResult struct {
	List           V5GetInternalTransferRecordsList `json:"list"`
	NextPageCursor string                           `json:"nextPageCursor"`
}

// V5GetInternalTransferRecordsList :
type V5GetInternalTransferRecordsList []V5GetInternalTransferRecordsItem

// V5GetInternalTransferRecordsItem :
type V5GetInternalTransferRecordsItem struct {
	TransferID      string           `json:"transferId"`
	Coin            Coin             `json:"coin"`
	Amount          string           `json:"amount"`
	FromAccountType AccountTypeV5    `json:"fromAccountType"`
	ToAccountType   AccountTypeV5    `json:"toAccountType"`
	Timestamp       string           `json:"timestamp"`
	Status          TransferStatusV5 `json:"status"`
}

// GetInternalTransferRecords :
func (s *V5AssetService) GetInternalTransferRecords(param V5GetInternalTransferRecordsParam) (*V5GetInternalTransferRecordsResponse, error) {
	var res V5GetInternalTransferRecordsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/asset/transfer/query-inter-transfer-list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetDepositRecordsParam :
type V5GetDepositRecordsParam struct {
	Coin      *Coin   `url:"coin,omitempty"`
	StartTime *int64  `url:"startTime,omitempty"` // Start time (ms). Default value: 30 days before the current time
	EndTime   *int64  `url:"endTime,omitempty"`   // End time (ms). Default value: current time
	Limit     *int    `url:"limit,omitempty"`     // Number of items per page, [1, 50]. Default value: 50
	Cursor    *string `url:"cursor,omitempty"`
}

// V5GetDepositRecordsResponse :
type V5GetDepositRecordsResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetDepositRecordsResult `json:"result"`
}

// V5GetDepositRecordsResult :
type V5GetDepositRecordsResult struct {
	Rows           V5GetDepositRecordsRows `json:"rows"`
	NextPageCursor string                  `json:"nextPageCursor"`
}

// V5GetDepositRecordsRows :
type V5GetDepositRecordsRows []V5GetDepositRecordsRow

// V5GetDepositRecordsRow :
type V5GetDepositRecordsRow struct {
	Coin          Coin            `json:"coin"`
	Chain         string          `json:"chain"`
	Amount        string          `json:"amount"`
	TxID          string          `json:"txID"`
	Status        DepositStatusV5 `json:"status"`
	ToAddress     string          `json:"toAddress"`
	Tag           string          `json:"tag"`
	DepositFee    string          `json:"depositFee"`
	SuccessAt     string          `json:"successAt"`
	Confirmations string          `json:"confirmations"`
	TxIndex       string          `json:"txIndex"`
	BlockHash     string          `json:"blockHash"`
}

// GetDepositRecords :
func (s *V5AssetService) GetDepositRecords(param V5GetDepositRecordsParam) (*V5GetDepositRecordsResponse, error) {
	var res V5GetDepositRecordsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/asset/deposit/query-record", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetSubDepositRecordsParam :
type V5GetSubDepositRecordsParam struct {
	SubMemberID string `url:"subMemberId"`

	Coin      *Coin   `url:"coin,omitempty"`
	StartTime *int64  `url:"startTime,omitempty"` // Start time (ms). Default value: 30 days before the current time
	EndTime   *int64  `url:"endTime,omitempty"`   // The start timestamp (ms)
	Limit     *int    `url:"limit,omitempty"`     // Limit for data size per page. [1, 50]. Default: 50
	Cursor    *string `url:"cursor,omitempty"`
}

// V5GetSubDepositRecordsResponse :
type V5GetSubDepositRecordsResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetSubDepositRecordsResult `json:"result"`
}

// V5GetSubDepositRecordsResult :
type V5GetSubDepositRecordsResult struct {
	Rows           V5GetSubDepositRecordsRows `json:"rows"`
	NextPageCursor string                     `json:"nextPageCursor"`
}

// V5GetSubDepositRecordsRows :
type V5GetSubDepositRecordsRows []V5GetSubDepositRecordsRow

// V5GetSubDepositRecordsRow :
type V5GetSubDepositRecordsRow struct {
	Coin          Coin            `json:"coin"`
	Chain         string          `json:"chain"`
	Amount        string          `json:"amount"`
	TxID          string          `json:"txID"`
	Status        DepositStatusV5 `json:"status"`
	ToAddress     string          `json:"toAddress"`
	Tag           string          `json:"tag"`
	DepositFee    string          `json:"depositFee"`
	SuccessAt     string          `json:"successAt"`
	Confirmations string          `json:"confirmations"`
	TxIndex       string          `json:"txIndex"`
	BlockHash     string          `json:"blockHash"`
}

// GetSubDepositRecords :
func (s *V5AssetService) GetSubDepositRecords(param V5GetSubDepositRecordsParam) (*V5GetSubDepositRecordsResponse, error) {
	var res V5GetSubDepositRecordsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/asset/deposit/query-sub-member-record", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetInternalDepositRecordsParam :
type V5GetInternalDepositRecordsParam struct {
	StartTime *int64  `url:"startTime,omitempty"` // Start time (ms). Default value: 30 days before the current time
	EndTime   *int64  `url:"endTime,omitempty"`   // End time (ms). Default value: current time
	Coin      *Coin   `url:"coin,omitempty"`
	Cursor    *string `url:"cursor,omitempty"`
	Limit     *int    `url:"limit,omitempty"` // Number of items per page, [1, 50]. Default value: 50
}

// V5GetInternalDepositRecordsResponse :
type V5GetInternalDepositRecordsResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetInternalDepositRecordsResult `json:"result"`
}

// V5GetInternalDepositRecordsResult :
type V5GetInternalDepositRecordsResult struct {
	Rows           V5GetInternalDepositRecordsRows `json:"rows"`
	NextPageCursor string                          `json:"nextPageCursor"`
}

// V5GetInternalDepositRecordsRows :
type V5GetInternalDepositRecordsRows []V5GetInternalDepositRecordsRow

// V5GetInternalDepositRecordsRow :
type V5GetInternalDepositRecordsRow struct {
	ID          string                  `json:"id"`
	Type        string                  `json:"type"`
	Coin        Coin                    `json:"coin"`
	Amount      string                  `json:"amount"`
	Status      InternalDepositStatusV5 `json:"status"`
	Address     string                  `json:"address"` // Email address or phone number
	CreatedTime string                  `json:"createdTime"`
}

// GetInternalDepositRecords :
func (s *V5AssetService) GetInternalDepositRecords(param V5GetInternalDepositRecordsParam) (*V5GetInternalDepositRecordsResponse, error) {
	var res V5GetInternalDepositRecordsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/asset/deposit/query-internal-record", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetWithdrawalRecordsParam :
type V5GetWithdrawalRecordsParam struct {
	WithdrawID   *string         `url:"withdrawId,omitempty"`
	Coin         *Coin           `url:"coin,omitempty"`
	WithdrawType *WithdrawTypeV5 `url:"withdrawType,omitempty"`
	StartTime    *int64          `url:"startTime,omitempty"` // The start timestamp (ms)
	EndTime      *int64          `url:"endTime,omitempty"`   // The start timestamp (ms)
	Limit        *int            `url:"limit,omitempty"`     // Limit for data size per page. [1, 50]. Default: 50
	Cursor       *string         `url:"cursor,omitempty"`
}

// V5GetWithdrawalRecordsResponse :
type V5GetWithdrawalRecordsResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetWithdrawalRecordsResult `json:"result"`
}

// V5GetWithdrawalRecordsResult :
type V5GetWithdrawalRecordsResult struct {
	Rows           V5GetWithdrawalRecordsRows `json:"rows"`
	NextPageCursor string                     `json:"nextPageCursor"`
}

// V5GetWithdrawalRecordsRows :
type V5GetWithdrawalRecordsRows []V5GetWithdrawalRecordsRow

// V5GetWithdrawalRecordsRow :
type V5GetWithdrawalRecordsRow struct {
	WithdrawId   string           `json:"withdrawId"`
	TxID         string           `json:"txId"`
	WithdrawType WithdrawTypeV5   `json:"withdrawType"`
	Coin         Coin             `json:"coin"`
	Chain        string           `json:"chain"`
	Amount       string           `json:"amount"`
	WithdrawFee  string           `json:"withdrawFee"`
	Status       WithdrawStatusV5 `json:"status"`
	ToAddress    string           `json:"toAddress"`
	Tag          string           `json:"tag"`
	CreatedTime  string           `json:"createTime"`
	UpdatedTime  string           `json:"updateTime"`
}

// GetWithdrawalRecords :
func (s *V5AssetService) GetWithdrawalRecords(param V5GetWithdrawalRecordsParam) (*V5GetWithdrawalRecordsResponse, error) {
	var res V5GetWithdrawalRecordsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/asset/withdraw/query-record", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetCoinInfoParam :
type V5GetCoinInfoParam struct {
	Coin *Coin `url:"coin,omitempty"`
}

// V5GetCoinInfoResponse :
type V5GetCoinInfoResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetCoinInfoResult `json:"result"`
}

// V5GetCoinInfoResult :
type V5GetCoinInfoResult struct {
	Rows V5GetCoinInfoRows `json:"rows"`
}

// V5GetCoinInfoRows :
type V5GetCoinInfoRows []V5GetCoinInfoRow

// V5GetCoinInfoRow :
type V5GetCoinInfoRow struct {
	Name         string              `json:"name"`
	Coin         Coin                `json:"coin"`
	RemainAmount string              `json:"remainAmount"`
	Chains       V5GetCoinInfoChains `json:"chains"`
}

// V5GetCoinInfoChains :
type V5GetCoinInfoChains []V5GetCoinInfoChain

// V5GetCoinInfoChain :
type V5GetCoinInfoChain struct {
	Chain                 string `json:"chain"`
	ChainType             string `json:"chainType"`
	Confirmation          string `json:"confirmation"`
	WithdrawFee           string `json:"withdrawFee"`
	DepositMin            string `json:"depositMin"`
	WithdrawMin           string `json:"withdrawMin"`
	MinAccuracy           string `json:"minAccuracy"`
	ChainDeposit          string `json:"chainDeposit"`
	ChainWithdraw         string `json:"chainWithdraw"`
	WithdrawPercentageFee string `json:"withdrawPercentageFee"`
}

// GetCoinInfo :
func (s *V5AssetService) GetCoinInfo(param V5GetCoinInfoParam) (*V5GetCoinInfoResponse, error) {
	var res V5GetCoinInfoResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/asset/coin/query-info", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
