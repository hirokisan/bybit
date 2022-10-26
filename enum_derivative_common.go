package bybit

// SymbolDerivative :
type SymbolDerivative string

const (
	// SymbolDerivativeBTCUSDT :
	SymbolDerivativeBTCUSDT = SymbolDerivative("BTCUSDT")
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
