package bybit

import "github.com/google/go-querystring/query"

// V5AssetServiceI :
type V5AssetServiceI interface {
	GetInternalTransferRecords(V5GetInternalTransferRecordsParam) (*V5GetInternalTransferRecordsResponse, error)
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
