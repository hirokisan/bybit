package bybit

// DerivativeUnifiedMarginServiceI :
type DerivativeUnifiedMarginServiceI interface {
	// Market Data Endpoints
	DerivativesOrderBook(DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
	DerivativesKline(DerivativesKlineParam) (*DerivativesKlineResponse, error)
	DerivativesTickers(DerivativesTickersParam) (*DerivativesTickersResponse, error)
	DerivativesTickersForOption(DerivativesTickersForOptionParam) (*DerivativesTickersForOptionResponse, error)
}

// DerivativeUnifiedMarginService :
type DerivativeUnifiedMarginService struct {
	client *Client

	*DerivativeCommonService
}
