package bybit

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// FutureCommonService :
type FutureCommonService struct {
	client *Client
}

// APIKeyInfoResult :
type APIKeyInfoResult struct {
	APIKey        string    `json:"api_key"`
	Type          string    `json:"type"`
	UserID        int       `json:"user_id"`
	InviterID     int       `json:"inviter_id"`
	Ips           []string  `json:"ips"`
	Note          string    `json:"note"`
	Permissions   []string  `json:"permissions"`
	CreatedAt     time.Time `json:"created_at"`
	ExpiredAt     time.Time `json:"expired_at"`
	ReadOnly      bool      `json:"read_only"`
	VipLevel      string    `json:"vip_level"`
	MktMakerLevel string    `json:"mkt_maker_level"`
	AffiliateID   int       `json:"affiliate_id"`
}

// APIKeyInfoResponse :
type APIKeyInfoResponse struct {
	CommonResponse `json:",inline"`
	Result         []APIKeyInfoResult `json:"result"`
}

// APIKeyInfo :
func (s *FutureCommonService) APIKeyInfo() (*APIKeyInfoResponse, error) {
	var (
		res   APIKeyInfoResponse
		query url.Values
	)

	if err := s.client.getPrivately("/v2/private/account/api-key", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// BalanceResponse :
type BalanceResponse struct {
	CommonResponse `json:",inline"`
	Result         BalanceResult `json:"result"`
}

// BalanceResult :
type BalanceResult struct {
	Balance map[Coin]Balance
}

// UnmarshalJSON :
func (r *BalanceResult) UnmarshalJSON(data []byte) error {
	parsedData := map[string]Balance{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	r.Balance = map[Coin]Balance{}
	for coin, balanceData := range parsedData {
		r.Balance[Coin(coin)] = balanceData
	}
	return nil
}

// Balance :
type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

// Balance :
func (s *FutureCommonService) Balance(coin Coin) (*BalanceResponse, error) {
	var res BalanceResponse

	query := url.Values{}
	query.Add("coin", string(coin))
	if err := s.client.getPrivately("/v2/private/wallet/balance", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// OrderBookResponse :
type OrderBookResponse struct {
	CommonResponse `json:",inline"`
	Result         []OrderBookResult `json:"result"`
}

// OrderBookResult :
type OrderBookResult struct {
	Symbol SymbolFuture `json:"symbol"`
	Price  string       `json:"price"`
	Size   float64      `json:"size"`
	Side   Side         `json:"side"`
}

// OrderBook :
func (s *FutureCommonService) OrderBook(symbol SymbolFuture) (*OrderBookResponse, error) {
	var res OrderBookResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPublicly("/v2/public/orderBook/L2", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListKlineParam :
type ListKlineParam struct {
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`

	Limit *int `url:"limit,omitempty"`
}

// ListKlineResponse :
type ListKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListKlineResult `json:"result"`
}

// ListKlineResult :
type ListKlineResult struct {
	Symbol   SymbolFuture `json:"symbol"`
	Interval string       `json:"interval"`
	OpenTime int          `json:"open_time"`
	Open     string       `json:"open"`
	High     string       `json:"high"`
	Low      string       `json:"low"`
	Close    string       `json:"close"`
	Volume   string       `json:"volume"`
	Turnover string       `json:"turnover"`
}

// ListKline :
func (s *FutureCommonService) ListKline(param ListKlineParam) (*ListKlineResponse, error) {
	var res ListKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/kline/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// TickersResponse :
type TickersResponse struct {
	CommonResponse `json:",inline"`
	Result         []TickersResult `json:"result"`
}

// TickersResult :
type TickersResult struct {
	Symbol               SymbolFuture  `json:"symbol"`
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

// Tickers :
func (s *FutureCommonService) Tickers(symbol SymbolFuture) (*TickersResponse, error) {
	var res TickersResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPublicly("/v2/public/tickers", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// TradingRecordsParam :
type TradingRecordsParam struct {
	Symbol SymbolFuture `url:"symbol"`

	From  *int `url:"from,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// TradingRecordsResponse :
type TradingRecordsResponse struct {
	CommonResponse `json:",inline"`
	Result         []TradingRecordsResult `json:"result"`
}

// TradingRecordsResult :
type TradingRecordsResult struct {
	ID     float64      `json:"id"`
	Symbol SymbolFuture `json:"symbol"`
	Price  float64      `json:"price"`
	Qty    float64      `json:"qty"`
	Side   Side         `json:"side"`
	Time   string       `json:"time"`
}

// TradingRecords :
func (s *FutureCommonService) TradingRecords(param TradingRecordsParam) (*TradingRecordsResponse, error) {
	var res TradingRecordsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/trading-records", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SymbolsResponse :
type SymbolsResponse struct {
	CommonResponse `json:",inline"`
	Result         []SymbolsResult `json:"result"`
}

// SymbolsResult :
type SymbolsResult struct {
	Name           string         `json:"name"`
	BaseCurrency   string         `json:"base_currency"`
	QuoteCurrency  string         `json:"quote_currency"`
	PriceScale     float64        `json:"price_scale"`
	TakerFee       string         `json:"taker_fee"`
	MakerFee       string         `json:"maker_fee"`
	LeverageFilter LeverageFilter `json:"leverage_filter"`
	PriceFilter    PriceFilter    `json:"price_filter"`
	LotSizeFilter  LotSizeFilter  `json:"lot_size_filter"`
}

// LeverageFilter :
type LeverageFilter struct {
	MinLeverage  float64 `json:"min_leverage"`
	MaxLeverage  float64 `json:"max_leverage"`
	LeverageStep string  `json:"leverage_step"`
}

// PriceFilter :
type PriceFilter struct {
	MinPrice string `json:"min_price"`
	MaxPrice string `json:"max_price"`
	TickSize string `json:"tick_size"`
}

// LotSizeFilter :
type LotSizeFilter struct {
	MaxTradingQty float64 `json:"max_trading_qty"`
	MinTradingQty float64 `json:"min_trading_qty"`
	QtyStep       float64 `json:"qty_step"`
}

// Symbols :
func (s *FutureCommonService) Symbols() (*SymbolsResponse, error) {
	var res SymbolsResponse

	if err := s.client.getPublicly("/v2/public/symbols", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// MarkPriceKlineResponse :
type MarkPriceKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []MarkPriceKlineResult `json:"result"`
}

// MarkPriceKlineResult :
type MarkPriceKlineResult struct {
	Symbol  SymbolFuture `json:"symbol"`
	Period  Period       `json:"period"`
	StartAt int          `json:"start_at"`
	Open    float64      `json:"open"`
	High    float64      `json:"high"`
	Low     float64      `json:"low"`
	Close   float64      `json:"close"`
}

// MarkPriceKlineParam :
type MarkPriceKlineParam struct {
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`

	Limit *int `url:"limit,omitempty"`
}

// MarkPriceKline :
func (s *FutureCommonService) MarkPriceKline(param MarkPriceKlineParam) (*MarkPriceKlineResponse, error) {
	var res MarkPriceKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/mark-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// IndexPriceKlineResponse :
type IndexPriceKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []IndexPriceKlineResult `json:"result"`
}

// IndexPriceKlineResult :
type IndexPriceKlineResult struct {
	Symbol   SymbolFuture `json:"symbol"`
	Period   Period       `json:"period"`
	OpenTime int          `json:"open_time"`
	Open     string       `json:"open"`
	High     string       `json:"high"`
	Low      string       `json:"low"`
	Close    string       `json:"close"`
}

// IndexPriceKlineParam :
type IndexPriceKlineParam struct {
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`

	Limit *int `url:"limit,omitempty"`
}

// IndexPriceKline :
func (s *FutureCommonService) IndexPriceKline(param IndexPriceKlineParam) (*IndexPriceKlineResponse, error) {
	var res IndexPriceKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/index-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// OpenInterestResponse :
type OpenInterestResponse struct {
	CommonResponse `json:",inline"`
	Result         []OpenInterestResult `json:"result"`
}

// OpenInterestResult :
type OpenInterestResult struct {
	OpenInterest float64      `json:"open_interest"`
	Timestamp    int          `json:"timestamp"`
	Symbol       SymbolFuture `json:"symbol"`
}

// OpenInterestParam :
type OpenInterestParam struct {
	Symbol SymbolFuture `url:"symbol"`
	Period Period       `url:"period"`

	Limit *int `url:"limit,omitempty"`
}

// OpenInterest :
func (s *FutureCommonService) OpenInterest(param OpenInterestParam) (*OpenInterestResponse, error) {
	var res OpenInterestResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/open-interest", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// BigDealResponse :
type BigDealResponse struct {
	CommonResponse `json:",inline"`
	Result         []BigDealResult `json:"result"`
}

// BigDealResult :
type BigDealResult struct {
	Symbol    SymbolFuture `json:"symbol"`
	Side      Side         `json:"side"`
	Timestamp int          `json:"timestamp"`
	Value     float64      `json:"value"`
}

// BigDealParam :
type BigDealParam struct {
	Symbol SymbolFuture `url:"symbol"`

	Limit *int `url:"limit,omitempty"`
}

// BigDeal :
func (s *FutureCommonService) BigDeal(param BigDealParam) (*BigDealResponse, error) {
	var res BigDealResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/big-deal", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// AccountRatioResponse :
type AccountRatioResponse struct {
	CommonResponse `json:",inline"`
	Result         []AccountRatioResult `json:"result"`
}

// AccountRatioResult :
type AccountRatioResult struct {
	Symbol    SymbolFuture `json:"symbol"`
	BuyRatio  float64      `json:"buy_ratio"`
	SellRatio float64      `json:"sell_ratio"`
	Timestamp int          `json:"timestamp"`
}

// AccountRatioParam :
type AccountRatioParam struct {
	Symbol SymbolFuture `url:"symbol"`
	Period Period       `url:"period"`

	Limit *int `url:"limit,omitempty"`
}

// AccountRatio :
func (s *FutureCommonService) AccountRatio(param AccountRatioParam) (*AccountRatioResponse, error) {
	var res AccountRatioResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v2/public/account-ratio", queryString, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
