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
	SetTradingStop(V5SetTradingStopParam) (*V5SetTradingStopResponse, error)
	SetTpSlMode(V5SetTpSlModeParam) (*V5SetTpSlModeResponse, error)
	SwitchPositionMode(V5SwitchPositionModeParam) (*V5SwitchPositionModeResponse, error)
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

// V5SetTradingStopParam :
type V5SetTradingStopParam struct {
	Category    CategoryV5  `json:"category"`
	Symbol      SymbolV5    `json:"symbol"`
	PositionIdx PositionIdx `json:"positionIdx"`

	TakeProfit   *string    `json:"takeProfit,omitempty"`
	StopLoss     *string    `json:"stopLoss,omitempty"`
	TrailingStop *string    `json:"trailingStop,omitempty"`
	TpTriggerBy  *TriggerBy `json:"tpTriggerBy,omitempty"`
	SlTriggerBy  *TriggerBy `json:"slTriggerBy,omitempty"`
	ActivePrice  *string    `json:"activePrice,omitempty"`
	TpSize       *string    `json:"tpSize,omitempty"`
	SlSize       *string    `json:"slSize,omitempty"`
}

func (p V5SetTradingStopParam) validate() error {
	if p.Category != CategoryV5Linear && p.Category != CategoryV5Inverse {
		return fmt.Errorf("only linear and inverse are supported for category")
	}
	if p.TakeProfit == nil && p.StopLoss == nil {
		return fmt.Errorf("takeProfit or stopLoss needed")
	}
	return nil
}

// V5SetTradingStopResponse :
type V5SetTradingStopResponse struct {
	CommonV5Response `json:",inline"`
	Result           interface{} `json:"result"` // no content
}

// SetTradingStop :
func (s *V5PositionService) SetTradingStop(param V5SetTradingStopParam) (*V5SetTradingStopResponse, error) {
	var res V5SetTradingStopResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/position/trading-stop", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5SetTpSlModeParam :
type V5SetTpSlModeParam struct {
	Category CategoryV5 `json:"category"`
	Symbol   SymbolV5   `json:"symbol"`
	TpSlMode TpSlMode   `json:"tpSlMode"`
}

func (p V5SetTpSlModeParam) validate() error {
	if p.Category != CategoryV5Linear && p.Category != CategoryV5Inverse {
		return fmt.Errorf("only linear and inverse are supported for category")
	}
	return nil
}

// V5SetTpSlModeResponse :
type V5SetTpSlModeResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5SetTpSlModeResult `json:"result"`
}

// V5SetTpSlModeResult :
type V5SetTpSlModeResult struct {
	TpSlMode TpSlMode `json:"tpSlMode"`
}

// SetTpSlMode :
func (s *V5PositionService) SetTpSlMode(param V5SetTpSlModeParam) (*V5SetTpSlModeResponse, error) {
	var res V5SetTpSlModeResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/position/set-tpsl-mode", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5SwitchPositionModeParam :
type V5SwitchPositionModeParam struct {
	Category CategoryV5   `json:"category"`
	Mode     PositionMode `json:"mode"`

	Symbol *SymbolV5 `json:"symbol,omitempty"`
	Coin   *Coin     `json:"coin,omitempty"`
}

func (p V5SwitchPositionModeParam) validate() error {
	if p.Symbol == nil && p.Coin == nil {
		return fmt.Errorf("symbol or coin is required")
	}
	return nil
}

// V5SwitchPositionModeResponse :
type V5SwitchPositionModeResponse struct {
	CommonV5Response `json:",inline"`
	Result           interface{} `json:"result"` // no content
}

// SwitchPositionMode :
func (s *V5PositionService) SwitchPositionMode(param V5SwitchPositionModeParam) (*V5SwitchPositionModeResponse, error) {
	var res V5SwitchPositionModeResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON("/v5/position/switch-mode", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}
