package bybit

// DerivativeContractServiceI :
type DerivativeContractServiceI interface {
	// Market Data Endpoints
	DerivativesOrderBook(DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
	DerivativesKline(DerivativesKlineParam) (*DerivativesKlineResponse, error)
	DerivativesTickers(DerivativesTickersParam) (*DerivativesTickersResponse, error)
	DerivativesTickersForOption(DerivativesTickersForOptionParam) (*DerivativesTickersForOptionResponse, error)
	DerivativesInstruments(DerivativesInstrumentsParam) (*DerivativesInstrumentsResponse, error)
	DerivativesInstrumentsForOption(DerivativesInstrumentsForOptionParam) (*DerivativesInstrumentsForOptionResponse, error)
}

// DerivativeContractService :
type DerivativeContractService struct {
	client *Client

	*DerivativeCommonService
}
