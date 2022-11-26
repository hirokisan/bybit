package bybit

// DerivativeUnifiedMarginServiceI :
type DerivativeUnifiedMarginServiceI interface {
	// Market Data Endpoints
	DerivativesOrderBook(DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
	DerivativesKline(DerivativesKlineParam) (*DerivativesKlineResponse, error)
	DerivativesTickers(DerivativesTickersParam) (*DerivativesTickersResponse, error)
	DerivativesTickersForOption(DerivativesTickersForOptionParam) (*DerivativesTickersForOptionResponse, error)
	DerivativesInstruments(DerivativesInstrumentsParam) (*DerivativesInstrumentsResponse, error)
	DerivativesInstrumentsForOption(DerivativesInstrumentsForOptionParam) (*DerivativesInstrumentsForOptionResponse, error)
	DerivativesMarkPriceKline(DerivativesMarkPriceKlineParam) (*DerivativesMarkPriceKlineResponse, error)
	DerivativesIndexPriceKline(DerivativesIndexPriceKlineParam) (*DerivativesIndexPriceKlineResponse, error)
}

// DerivativeUnifiedMarginService :
type DerivativeUnifiedMarginService struct {
	client *Client

	*DerivativeCommonService
}
