package bybit

import "github.com/google/go-querystring/query"

// V5ExecutionServiceI :
type V5ExecutionServiceI interface {
	GetExecutionList(V5GetExecutionParam) (*V5GetExecutionListResponse, error)
}

// V5ExecutionService :
type V5ExecutionService struct {
	client *Client
}

// V5GetExecutionParam :
type V5GetExecutionParam struct {
	Category CategoryV5 `url:"category"`

	Symbol      *SymbolV5   `url:"symbol,omitempty"`
	OrderID     *string     `url:"orderId,omitempty"`
	OrderLinkID *string     `url:"orderLinkId,omitempty"`
	BaseCoin    *Coin       `url:"baseCoin,omitempty"`
	StartTime   *int        `url:"startTime,omitempty"`
	EndTime     *int        `url:"endTime,omitempty"`
	ExecType    *ExecTypeV5 `url:"execType,omitempty"`
	Limit       *int        `url:"limit,omitempty"`
	Cursor      *string     `url:"cursor,omitempty"`
}

// V5GetExecutionListResponse :
type V5GetExecutionListResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetExecutionListResult `json:"result"`
}

// V5GetExecutionListResult :
type V5GetExecutionListResult struct {
	NextPageCursor string                   `json:"nextPageCursor"`
	Category       string                   `json:"category"`
	List           []V5GetExecutionListItem `json:"list"`
}

// V5GetExecutionListItem :
type V5GetExecutionListItem struct {
	Symbol          SymbolV5   `json:"symbol"`
	OrderID         string     `json:"orderId"`
	OrderLinkID     string     `json:"orderLinkId"`
	Side            Side       `json:"side"`
	OrderPrice      string     `json:"orderPrice"`
	OrderQty        string     `json:"orderQty"`
	LeavesQty       string     `json:"leavesQty"`
	OrderType       OrderType  `json:"orderType"`
	StopOrderType   string     `json:"stopOrderType"`
	ExecFee         string     `json:"execFee"`
	ExecID          string     `json:"execId"`
	ExecPrice       string     `json:"execPrice"`
	ExecQty         string     `json:"execQty"`
	ExecType        ExecTypeV5 `json:"execType"`
	ExecValue       string     `json:"execValue"`
	ExecTime        string     `json:"execTime"`
	IsMaker         bool       `json:"isMaker"`
	FeeRate         string     `json:"feeRate"`
	TradeIv         string     `json:"tradeIv"`
	MarkIv          string     `json:"markIv"`
	MarkPrice       string     `json:"markPrice"`
	IndexPrice      string     `json:"indexPrice"`
	UnderlyingPrice string     `json:"underlyingPrice"`
	BlockTradeID    string     `json:"blockTradeId"`
	ClosedSize      string     `json:"closedSize"`
}

func (s *V5ExecutionService) GetExecutionList(param V5GetExecutionParam) (*V5GetExecutionListResponse, error) {
	var res V5GetExecutionListResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately("/v5/execution/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
