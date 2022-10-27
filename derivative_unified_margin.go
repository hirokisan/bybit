package bybit

// DerivativeUnifiedMarginServiceI :
type DerivativeUnifiedMarginServiceI interface {
	DerivativesOrderBook(DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
	DerivativesKline(DerivativesKlineParam) (*DerivativesKlineResponse, error)
}

// DerivativeUnifiedMarginService :
type DerivativeUnifiedMarginService struct {
	client *Client

	*DerivativeCommonService
}
