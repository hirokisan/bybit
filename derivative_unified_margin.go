package bybit

import "context"

// DerivativeUnifiedMarginServiceI :
type DerivativeUnifiedMarginServiceI interface {
	// Market Data Endpoints
	DerivativesOrderBook(context.Context, DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
	DerivativesKline(context.Context, DerivativesKlineParam) (*DerivativesKlineResponse, error)
	DerivativesTickers(context.Context, DerivativesTickersParam) (*DerivativesTickersResponse, error)
	DerivativesTickersForOption(context.Context, DerivativesTickersForOptionParam) (*DerivativesTickersForOptionResponse, error)
	DerivativesInstruments(context.Context, DerivativesInstrumentsParam) (*DerivativesInstrumentsResponse, error)
	DerivativesInstrumentsForOption(context.Context, DerivativesInstrumentsForOptionParam) (*DerivativesInstrumentsForOptionResponse, error)
	DerivativesMarkPriceKline(context.Context, DerivativesMarkPriceKlineParam) (*DerivativesMarkPriceKlineResponse, error)
	DerivativesIndexPriceKline(context.Context, DerivativesIndexPriceKlineParam) (*DerivativesIndexPriceKlineResponse, error)
}

// DerivativeUnifiedMarginService :
type DerivativeUnifiedMarginService struct {
	client *Client

	*DerivativeCommonService
}
