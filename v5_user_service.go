package bybit

import (
	"encoding/json"
	"net/url"
	"time"
)

// V5UserServiceI :
type V5UserServiceI interface {
	CreateSubUID(param V5CreateSubUIDParam) (*V5CreateSubUIDResponse, error)
	GetSubUIDList() (*V5GetSubUIDListResponse, error)
	CreateSubUIDAPIKey(param V5CreateSubUIDAPIKeyParam) (*V5CreateSubUIDAPIKeyResponse, error)
	GetAPIKey() (*V5APIKeyResponse, error)
}

// V5UserService :
type V5UserService struct {
	client *Client
}

// V5CreateSubUIDParam :
type V5CreateSubUIDParam struct {
	Username   string  `json:"username"`
	MemberType int     `json:"memberType"`
	Password   *string `json:"password,omitempty"`
	Switch     *int    `json:"switch,omitempty"`
	Note       *string `json:"note,omitempty"`
}

// V5CreateSubUIDResponse :
type V5CreateSubUIDResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5CreateSubUIDResult `json:"result"`
}

// V5CreateSubUIDResult :
type V5CreateSubUIDResult struct {
	UID        string `json:"uid"`
	Username   string `json:"username"`
	MemberType int    `json:"memberType"`
	Status     int    `json:"status"`
	Remark     string `json:"remark"`
}

// V5GetSubUIDListResponse :
type V5GetSubUIDListResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetSubUIDListResult `json:"result"`
}

// V5GetSubUIDListResult :
type V5GetSubUIDListResult struct {
	SubMembers []V5SubMember `json:"subMembers"`
}

// V5SubMember :
type V5SubMember struct {
	UID         string `json:"uid"`
	Username    string `json:"username"`
	MemberType  int    `json:"memberType"`
	Status      int    `json:"status"`
	AccountMode int    `json:"accountMode"`
	Remark      string `json:"remark"`
}

// V5CreateSubUIDAPIKeyParam :
type V5CreateSubUIDAPIKeyParam struct {
	Subuid      int                     `json:"subuid"`
	ReadOnly    int                     `json:"readOnly"`
	Note        *string                 `json:"note,omitempty"`
	Ips         *string                 `json:"ips,omitempty"`
	Permissions V5APIKeyPermissionsParam `json:"permissions"`
}

// V5APIKeyPermissionsParam :
type V5APIKeyPermissionsParam struct {
	ContractTrade []string `json:"ContractTrade,omitempty"`
	Spot          []string `json:"Spot,omitempty"`
	Wallet        []string `json:"Wallet,omitempty"`
	Options       []string `json:"Options,omitempty"`
	Derivatives   []string `json:"Derivatives,omitempty"`
	CopyTrading   []string `json:"CopyTrading,omitempty"`
	BlockTrade    []string `json:"BlockTrade,omitempty"`
	Exchange      []string `json:"Exchange,omitempty"`
	NFT           []string `json:"NFT,omitempty"`
}

// V5CreateSubUIDAPIKeyResponse :
type V5CreateSubUIDAPIKeyResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5CreateSubUIDAPIKeyResult `json:"result"`
}

// V5CreateSubUIDAPIKeyResult :
type V5CreateSubUIDAPIKeyResult struct {
	ID          string `json:"id"`
	Note        string `json:"note"`
	APIKey      string `json:"apiKey"`
	ReadOnly    int    `json:"readOnly"`
	Secret      string `json:"secret"`
	Permissions V5APIKeyPermissions `json:"permissions"`
}

// V5APIKeyPermissions :
type V5APIKeyPermissions struct {
	ContractTrade []string `json:"ContractTrade"`
	Spot          []string `json:"Spot"`
	Wallet        []string `json:"Wallet"`
	Options       []string `json:"Options"`
	Derivatives   []string `json:"Derivatives"`
	CopyTrading   []string `json:"CopyTrading"`
	BlockTrade    []string `json:"BlockTrade"`
	Exchange      []string `json:"Exchange"`
	NFT           []string `json:"NFT"`
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
		ContractTrade []string `json:"ContractTrade"`
		Spot          []string `json:"Spot"`
		Wallet        []string `json:"Wallet"`
		Options       []string `json:"Options"`
		Derivatives   []string `json:"Derivatives"`
		CopyTrading   []string `json:"CopyTrading"`
		BlockTrade    []string `json:"BlockTrade"`
		Exchange      []string `json:"Exchange"`
		Nft           []string `json:"NFT"`
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

// CreateSubUID :
func (s *V5UserService) CreateSubUID(param V5CreateSubUIDParam) (*V5CreateSubUIDResponse, error) {
	var (
		res V5CreateSubUIDResponse
	)

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postV5JSON("/v5/user/create-sub-member", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetSubUIDList :
func (s *V5UserService) GetSubUIDList() (*V5GetSubUIDListResponse, error) {
	var (
		res V5GetSubUIDListResponse
	)

	if err := s.client.getV5Privately("/v5/user/query-sub-members", url.Values{}, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateSubUIDAPIKey :
func (s *V5UserService) CreateSubUIDAPIKey(param V5CreateSubUIDAPIKeyParam) (*V5CreateSubUIDAPIKeyResponse, error) {
	var (
		res V5CreateSubUIDAPIKeyResponse
	)

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postV5JSON("/v5/user/create-sub-api", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetAPIKey :
func (s *V5UserService) GetAPIKey() (*V5APIKeyResponse, error) {
	var (
		res V5APIKeyResponse
	)

	if err := s.client.getV5Privately("/v5/user/query-api", url.Values{}, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
