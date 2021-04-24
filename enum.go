package bybit

// Coin :
type Coin string

const (
	// CoinBTC :
	CoinBTC = "BTC"
	// CoinETH :
	CoinETH = "ETH"
	// CoinEOS :
	CoinEOS = "EOS"
	// CoinXRP :
	CoinXRP = "XRP"
	// CoinUSDT :
	CoinUSDT = "USDT"
)

// SymbolInverse :
type SymbolInverse string

const (
	// SymbolInverseBTCUSD :
	SymbolInverseBTCUSD = SymbolInverse("BTCUSD")
	// SymbolInverseETHUSD :
	SymbolInverseETHUSD = SymbolInverse("ETHUSD")
	// SymbolInverseEOSUSD :
	SymbolInverseEOSUSD = SymbolInverse("EOSUSD")
	// SymbolInverseXRPUSD :
	SymbolInverseXRPUSD = SymbolInverse("XRPUSD")
)

// SymbolUSDT :
type SymbolUSDT string

const (
	// SymbolUSDTBTC :
	SymbolUSDTBTC = SymbolUSDT("BTCUSDT")
	// SymbolUSDTETH :
	SymbolUSDTETH = SymbolUSDT("ETHUSDT")
	// SymbolUSDTLTC :
	SymbolUSDTLTC = SymbolUSDT("LTCUSDT")
	// SymbolUSDTLINK :
	SymbolUSDTLINK = SymbolUSDT("LINKUSDT")
	// SymbolUSDTXTZ :
	SymbolUSDTXTZ = SymbolUSDT("XTZUSDT")
	// SymbolUSDTBCH :
	SymbolUSDTBCH = SymbolUSDT("BCHUSDT")
	// SymbolUSDTADA :
	SymbolUSDTADA = SymbolUSDT("ADAUSDT")
	// SymbolUSDTDOT :
	SymbolUSDTDOT = SymbolUSDT("DOTUSDT")
	// SymbolUSDTUNI :
	SymbolUSDTUNI = SymbolUSDT("UNIUSDT")
)

// Side :
type Side string

const (
	// SideNone : not defined officially
	SideNone = Side("None")
	// SideBuy :
	SideBuy = Side("Buy")
	// SideSell :
	SideSell = Side("Sell")
)

// OrderType :
type OrderType string

const (
	// OrderTypeLimit :
	OrderTypeLimit = OrderType("Limit")
	// OrderTypeMarket :
	OrderTypeMarket = OrderType("Market")
)

// OrderStatus :
type OrderStatus string

const (
	// OrderStatusCreated :
	OrderStatusCreated = OrderStatus("Created")
	// OrderStatusRejected :
	OrderStatusRejected = OrderStatus("Rejected")
	// OrderStatusNew :
	OrderStatusNew = OrderStatus("New")
	// OrderStatusPartiallyFilled :
	OrderStatusPartiallyFilled = OrderStatus("PartiallyFilled")
	// OrderStatusFilled :
	OrderStatusFilled = OrderStatus("Filled")
	// OrderStatusCancelled :
	OrderStatusCancelled = OrderStatus("Cancelled")
	// OrderStatusPendingCancel :
	OrderStatusPendingCancel = OrderStatus("PendingCancel")
)

// TimeInForce :
type TimeInForce string

const (
	// TimeInForceGoodTillCancel :
	TimeInForceGoodTillCancel = TimeInForce("GoodTillCancel")
	// TimeInForceImmediateOrCancel :
	TimeInForceImmediateOrCancel = TimeInForce("ImmediateOrCancel")
	// TimeInForceFillOrKill :
	TimeInForceFillOrKill = TimeInForce("FillOrKill")
	// TimeInForcePostOnly :
	TimeInForcePostOnly = TimeInForce("PostOnly")
)

// Interval :
type Interval string

const (
	// Interval1 :
	Interval1 = Interval("1")
	// Interval3 :
	Interval3 = Interval("3")
	// Interval5 :
	Interval5 = Interval("5")
	// Interval15 :
	Interval15 = Interval("15")
	// Interval30 :
	Interval30 = Interval("30")
	// Interval60 :
	Interval60 = Interval("60")
	// Interval120 :
	Interval120 = Interval("120")
	// Interval240 :
	Interval240 = Interval("240")
	// Interval360 :
	Interval360 = Interval("360")
	// Interval720 :
	Interval720 = Interval("720")
	// IntervalD :
	IntervalD = Interval("D")
	// IntervalW :
	IntervalW = Interval("W")
	// IntervalM :
	IntervalM = Interval("M")
)

// TickDirection :
type TickDirection string

const (
	// TickDirectionPlusTick :
	TickDirectionPlusTick = TickDirection("PlusTick")
	// TickDirectionZeroPlusTick :
	TickDirectionZeroPlusTick = TickDirection("ZeroPlusTick")
	// TickDirectionMinusTick :
	TickDirectionMinusTick = TickDirection("MinusTick")
	// TickDirectionZeroMinusTick :
	TickDirectionZeroMinusTick = TickDirection("ZeroMinusTick")
)

// Period :
type Period string

const (
	// Period5min :
	Period5min = Period("5min")
	// Period15min :
	Period15min = Period("15min")
	// Period30min :
	Period30min = Period("30min")
	// Period1h :
	Period1h = Period("1h")
	// Period4h :
	Period4h = Period("4h")
	// Period1d :
	Period1d = Period("1d")
)
