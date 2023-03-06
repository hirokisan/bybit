package bybit

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
)

// V5PositionServiceI :
type V5PositionServiceI interface {
	GetPositionInfo(V5GetPositionInfoParam) (*V5GetPositionInfoResponse, error)
	SetLeverage(V5SetLeverageParam) (*V5SetLeverageResponse, error)
}

// V5PositionService :
type V5PositionService struct {
	client *Client
}

// V5GetPositionInfoParam :
type V5GetPositionInfoParam struct {
	Category CategoryV5 `url:"category"`

	Symbol     *SymbolV5 `url:"symbol,omitempty"`
	BaseCoin   *Coin     `url:"baseCoin,omitempty"`   // option only
	SettleCoin *Coin     `url:"settleCoin,omitempty"` // Settle coin. For linear & inverse, either symbol or settleCon is required. symbol has a higher priority
	Limit      *int      `url:"limit,omitempty"`      // Limit for data size per page. [1, 200]. Default: 200
	Cursor     *string   `url:"cursor,omitempty"`     // Cursor. Used for pagination
}

// V5GetPositionInfoResponse :
type V5GetPositionInfoResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetPositionInfoResult `json:"result"`
}

// V5GetPositionInfoResult :
type V5GetPositionInfoResult struct {
	Category       CategoryV5            `json:"category"`
	NextPageCursor string                `json:"nextPageCursor"`
	List           V5GetPositionInfoList `json:"list"`
}

// V5GetPositionInfoList :
type V5GetPositionInfoList []V5GetPositionInfoItem

// V5GetPositionInfoItem :
type V5GetPositionInfoItem struct {
	Symbol         SymbolV5 `json:"symbol"`
	Leverage       string   `json:"leverage"`
	AvgPrice       string   `json:"avgPrice"`
	LiqPrice       string   `json:"liqPrice"`
	RiskLimitValue string   `json:"riskLimitValue"`
	TakeProfit     string   `json:"takeProfit"`
	PositionValue  string   `json:"positionValue"`
	TpSlMode       TpSlMode `json:"tpslMode"`
	RiskID         int      `json:"riskId"`
	TrailingStop   string   `json:"trailingStop"`
	UnrealisedPnl  string   `json:"unrealisedPnl"`
	MarkPrice      string   `json:"markPrice"`
	CumRealisedPnl string   `json:"cumRealisedPnl"`
	PositionMM     string   `json:"positionMM"`
	CreatedTime    string   `json:"createdTime"`
	PositionIdx    int      `json:"positionIdx"`
	PositionIM     string   `json:"positionIM"`
	UpdatedTime    string   `json:"updatedTime"`
	Side           Side     `json:"side"`
	BustPrice      string   `json:"bustPrice"`
	Size           string   `json:"size"`
	PositionStatus string   `json:"positionStatus"`
	StopLoss       string   `json:"stopLoss"`
	TradeMode      int      `json:"tradeMode"`
}

// GetPositionInfo :
func (s *V5PositionService) GetPositionInfo(param V5GetPositionInfoParam) (*V5GetPositionInfoResponse, error) {
	var res V5GetPositionInfoResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/position/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5SetLeverageParam :
type V5SetLeverageParam struct {
	Category     CategoryV5 `json:"category"`
	Symbol       SymbolV5   `json:"symbol"`
	BuyLeverage  string     `json:"buyLeverage"`
	SellLeverage string     `json:"sellLeverage"`
}

// V5SetLeverageResponse :
type V5SetLeverageResponse struct {
	CommonV5Response `json:",inline"`
	Result           interface{} `json:"result"`
}

// SetLeverage :
func (s *V5PositionService) SetLeverage(param V5SetLeverageParam) (*V5SetLeverageResponse, error) {
	var res V5SetLeverageResponse

	if param.Category == "" || param.Symbol == "" || param.BuyLeverage == "" || param.SellLeverage == "" {
		return nil, fmt.Errorf("Category, Symbol, BuyLeverage and SellLeverage needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/position/set-leverage", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}
