package bybit

import "net/url"

// FutureCommonService :
type FutureCommonService struct {
	client *Client
}

// LinearTickersResponse :
type LinearTickersResponse struct {
	CommonResponse `json:",inline"`
	Result         []LinearTickersResult `json:"result"`
}

// LinearTickersResult :
type LinearTickersResult struct {
	Symbol               SymbolUSDT    `json:"symbol"`
	BidPrice             string        `json:"bid_price"`
	AskPrice             string        `json:"ask_price"`
	LastPrice            string        `json:"last_price"`
	LastTickDirection    TickDirection `json:"last_tick_direction"`
	PrevPrice24h         string        `json:"prev_price_24h"`
	Price24hPcnt         string        `json:"price_24h_pcnt"`
	HighPrice24h         string        `json:"high_price_24h"`
	LowPrice24h          string        `json:"low_price_24h"`
	PrevPrice1h          string        `json:"prev_price_1h"`
	Price1hPcnt          string        `json:"price_1h_pcnt"`
	MarkPrice            string        `json:"mark_price"`
	IndexPrice           string        `json:"index_price"`
	OpenInterest         float64       `json:"open_interest"`
	OpenValue            string        `json:"open_value"`
	TotalTurnover        string        `json:"total_turnover"`
	Turnover24h          string        `json:"turnover_24h"`
	TotalVolume          float64       `json:"total_volume"`
	Volume24h            float64       `json:"volume_24h"`
	FundingRate          string        `json:"funding_rate"`
	PredictedFundingRate string        `json:"predicted_funding_rate"`
	NextFundingTime      string        `json:"next_funding_time"`
	CountdownHour        float64       `json:"countdown_hour"`
}

// LinearTickers :
func (s *FutureCommonService) LinearTickers(symbol SymbolUSDT) (*LinearTickersResponse, error) {
	var res LinearTickersResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPublicly("/v2/public/tickers", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
