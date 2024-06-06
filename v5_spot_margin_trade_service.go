package bybit

import (
	"context"
	"github.com/google/go-querystring/query"
)

// V5SpotMarginTradeServiceI :
type V5SpotMarginTradeServiceI interface {
	GetVIPMarginData(context.Context, V5GetVIPMarginDataParam) (*V5GetVIPMarginDataResponse, error)
}

// V5SpotMarginTradeService :
type V5SpotMarginTradeService struct {
	client *Client
}

// V5GetVIPMarginDataParam:
type V5GetVIPMarginDataParam struct {
	VipLevel string `url:"vipLevel"`
	Currency string `url:"currency"`
}

// V5GetVIPMarginDataResponse :
type V5GetVIPMarginDataResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5VIPMarginDataResult `json:"result"`
}

type V5VIPMarginDataResult struct {
	VipCoinList []struct {
		List     []V5VIPMarginDataVipCoin `json:"list"`
		VipLevel string                   `json:"vipLevel"`
	} `json:"vipCoinList"`
}

type V5VIPMarginDataVipCoin struct {
	Borrowable         bool   `json:"borrowable"`
	CollateralRatio    string `json:"collateralRatio"`
	Currency           string `json:"currency"`
	HourlyBorrowRate   string `json:"hourlyBorrowRate"`
	LiquidationOrder   string `json:"liquidationOrder"`
	MarginCollateral   bool   `json:"marginCollateral"`
	MaxBorrowingAmount string `json:"maxBorrowingAmount"`
}

func (s *V5SpotMarginTradeService) GetVIPMarginData(ctx context.Context, param V5GetVIPMarginDataParam) (*V5GetVIPMarginDataResponse, error) {
	res := new(V5GetVIPMarginDataResponse)

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPubliclyWithContext(ctx, "/v5/spot-margin-trade/data", queryString, res); err != nil {
		return nil, err
	}

	return res, nil
}
