package bybit

// DerivativeContractServiceI :
type DerivativeContractServiceI interface {
	DerivativesOrderBook(DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
}

// DerivativeContractService :
type DerivativeContractService struct {
	client *Client

	*DerivativeCommonService
}
