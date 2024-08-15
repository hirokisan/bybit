package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
)

// V5SpotMarginTradeServiceI :
type V5SpotMarginTradeServiceI interface {
	GetVIPMarginData(context.Context, V5GetVIPMarginDataParam) (*V5GetVIPMarginDataResponse, error)
	SetLeverage(ctx context.Context, param V5SetMarginLeverageParam) (*CommonV5Response, error)
	GetStatusAndLeverage(ctx context.Context) (*V5MarginStatusAndLeverageResponse, error)
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

	if err := s.client.getPublicly(ctx, "/v5/spot-margin-trade/data", queryString, res); err != nil {
		return nil, err
	}

	return res, nil
}

type V5SetMarginLeverageParam struct {
	Leverage string `json:"leverage"`
}

func (s *V5SpotMarginTradeService) SetLeverage(ctx context.Context, param V5SetMarginLeverageParam) (*CommonV5Response, error) {
	res := new(CommonV5Response)

	body, err := json.Marshal(param)
	if err != nil {
		return res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/spot-margin-trade/set-leverage", body, res); err != nil {
		return res, err
	}
	return res, nil
}

type V5MarginStatusAndLeverageResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5MarginStatusAndLeverag `json:"result"`
}

type V5MarginStatusAndLeverag struct {
	SpotLeverage      string `json:"spotLeverage"`
	SpotMarginMode    string `json:"spotMarginMode"`
	EffectiveLeverage string `json:"effectiveLeverage"`
}

func (s *V5SpotMarginTradeService) GetStatusAndLeverage(ctx context.Context) (*V5MarginStatusAndLeverageResponse, error) {
	res := new(V5MarginStatusAndLeverageResponse)

	queryString, err := query.Values(nil)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately(ctx, "/v5/spot-margin-trade/state", queryString, res); err != nil {
		return nil, err
	}

	return res, nil
}
