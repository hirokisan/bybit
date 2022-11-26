package bybit

// SymbolDerivative :
type SymbolDerivative string

const (
	// SymbolDerivativeBTCUSDT :
	SymbolDerivativeBTCUSDT = SymbolDerivative("BTCUSDT")
	// SymbolDerivativeBTC31MAR23_40000C :
	SymbolDerivativeBTC31MAR23_40000C = SymbolDerivative("BTC-31MAR23-40000-C")
)

// CategoryDerivative :
type CategoryDerivative string

const (
	// CategoryDerivativeLinear :
	CategoryDerivativeLinear = CategoryDerivative("linear")
	// CategoryDerivativeInverse :
	CategoryDerivativeInverse = CategoryDerivative("inverse")
	// CategoryDerivativeOption :
	CategoryDerivativeOption = CategoryDerivative("option")
)

// ContractTypeDerivative :
type ContractTypeDerivative string

const (
	// ContractTypeDerivativeInversePerpetual :
	ContractTypeDerivativeInversePerpetual = ContractTypeDerivative("InversePerpetual")
	// ContractTypeDerivativeLinearPerpetual :
	ContractTypeDerivativeLinearPerpetual = ContractTypeDerivative("LinearPerpetual")
	// ContractTypeDerivativeInverseFutures :
	ContractTypeDerivativeInverseFutures = ContractTypeDerivative("InverseFutures")
)

// StatusDerivative :
type StatusDerivative string

const (
	// StatusDerivativePending :
	StatusDerivativePending = StatusDerivative("Pending")
	// StatusDerivativeTrading :
	StatusDerivativeTrading = StatusDerivative("Trading")
	// StatusDerivativeSettling :
	StatusDerivativeSettling = StatusDerivative("Settling")
	// StatusDerivativeClosed :
	StatusDerivativeClosed = StatusDerivative("Closed")
)
