package bybit

import (
	"net/url"
	"time"
)

// V5UserServiceI :
type V5UserServiceI interface {
	GetAPIKey() (*V5APIKeyResponse, error) }

// V5UserService :
type V5UserService struct {
	client *Client
}

// V5APIKeyResponse :
type V5APIKeyResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5ApiKeyResult `json:"result"`
}

// V5ApiKeyResult :
type V5ApiKeyResult struct {
	ID          string `json:"id"`
	Note        string `json:"note"`
	APIKey      string `json:"apiKey"`
	ReadOnly    int    `json:"readOnly"`
	Secret      string `json:"secret"`
	Permissions struct {
		ContractTrade []interface{} `json:"ContractTrade"`
		Spot          []interface{} `json:"Spot"`
		Wallet        []string      `json:"Wallet"`
		Options       []interface{} `json:"Options"`
		Derivatives   []interface{} `json:"Derivatives"`
		CopyTrading   []interface{} `json:"CopyTrading"`
		BlockTrade    []interface{} `json:"BlockTrade"`
		Exchange      []interface{} `json:"Exchange"`
		Nft           []interface{} `json:"NFT"`
	} `json:"permissions"`
	Ips           []string  `json:"ips"`
	Type          int       `json:"type"`
	DeadlineDay   int       `json:"deadlineDay"`
	ExpiredAt     time.Time `json:"expiredAt"`
	CreatedAt     time.Time `json:"createdAt"`
	Unified       int       `json:"unified"`
	Uta           int       `json:"uta"`
	UserID        int       `json:"userID"`
	InviterID     int       `json:"inviterID"`
	VipLevel      string    `json:"vipLevel"`
	MktMakerLevel string    `json:"mktMakerLevel"`
	AffiliateID   int       `json:"affiliateID"`
}

// GetAPIKey :
func (s *V5AccountService) GetAPIKey() (*V5APIKeyResponse, error) {
	var (
		res V5APIKeyResponse
	)

	if err := s.client.getV5Privately("/v5/user/query-api", url.Values{}, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
