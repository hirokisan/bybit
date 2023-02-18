package bybit

// CategoryV5 :
type CategoryV5 string

const (
	// CategoryV5Spot :
	CategoryV5Spot = CategoryV5("spot")
	// CategoryV5Linear :
	CategoryV5Linear = CategoryV5("linear")
	// CategoryV5Inverse :
	CategoryV5Inverse = CategoryV5("inverse")
	// CategoryV5Option :
	CategoryV5Option = CategoryV5("option")
)

// SymbolV5 :
type SymbolV5 string

// SymbolV5 :
const (
	// USDT Perpetual:
	SymbolV5BTCUSDT = SymbolV5("BTCUSDT")
	SymbolV5ETHUSDT = SymbolV5("ETHUSDT")

	// USDC Perpetual
	SymbolV5BTCPERP = SymbolV5("BTCPERP")
	SymbolV5ETHPERP = SymbolV5("ETHPERP")

	// Inverse Perpetual
	SymbolV5BTCUSD = SymbolV5("BTCUSD")
	SymbolV5ETHUSD = SymbolV5("ETHUSD")

	// Inverse Futures
	SymbolV5BTCUSDH23 = SymbolV5("BTCUSDH23")
	SymbolV5BTCUSDM23 = SymbolV5("BTCUSDM23")
	SymbolV5BTCUSDU23 = SymbolV5("BTCUSDU23")
	SymbolV5BTCUSDZ23 = SymbolV5("BTCUSDZ23")

	// Spot
	SymbolV5ETHUSDC = SymbolV5("ETHUSDC")
)
