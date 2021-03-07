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

// Symbol :
type Symbol string

const (
	// SymbolBTCUSD :
	SymbolBTCUSD = Symbol("BTCUSD")
	// SymbolETHUSD :
	SymbolETHUSD = Symbol("ETHUSD")
	// SymbolEOSUSD :
	SymbolEOSUSD = Symbol("EOSUSD")
	// SymbolXRPUSD :
	SymbolXRPUSD = Symbol("XRPUSD")
)

// Side :
type Side string

const (
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

type Interval string

const (
	Interval1   = Interval("1")
	Interval3   = Interval("3")
	Interval5   = Interval("5")
	Interval15  = Interval("15")
	Interval30  = Interval("30")
	Interval60  = Interval("60")
	Interval120 = Interval("120")
	Interval240 = Interval("240")
	Interval360 = Interval("360")
	Interval720 = Interval("720")
	IntervalD   = Interval("D")
	IntervalW   = Interval("W")
	IntervalM   = Interval("M")
)
