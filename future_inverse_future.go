package bybit

// FutureInverseFutureServiceI :
type FutureInverseFutureServiceI interface {
	Balance(Coin) (*BalanceResponse, error)
	OrderBook(SymbolInverse) (*OrderBookResponse, error)
	ListKline(ListKlineParam) (*ListKlineResponse, error)
	Tickers(SymbolInverse) (*TickersResponse, error)
	TradingRecords(TradingRecordsParam) (*TradingRecordsResponse, error)
	Symbols() (*SymbolsResponse, error)
	IndexPriceKline(IndexPriceKlineParam) (*IndexPriceKlineResponse, error)
	OpenInterest(OpenInterestParam) (*OpenInterestResponse, error)
	BigDeal(BigDealParam) (*BigDealResponse, error)
	AccountRatio(AccountRatioParam) (*AccountRatioResponse, error)
}

// FutureInverseFutureService :
type FutureInverseFutureService struct {
	client *Client

	*FutureCommonService
}
