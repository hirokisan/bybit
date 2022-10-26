package bybit

// DerivativeUnifiedMarginServiceI :
type DerivativeUnifiedMarginServiceI interface {
	DerivativesOrderBook(DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
}

// DerivativeUnifiedMarginService :
type DerivativeUnifiedMarginService struct {
	client *Client

	*DerivativeCommonService
}
