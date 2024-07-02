package bybit

// MarginMode :
type MarginMode string

const (
	// MarginModeRegular :
	MarginModeRegular = MarginMode("REGULAR_MARGIN")
	// MarginModePortfolio :
	MarginModePortfolio = MarginMode("PORTFOLIO_MARGIN")
)

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

// TriggerDirection :
type TriggerDirection int

const (
	// TriggerDirectionRise : triggered when market price rises
	TriggerDirectionRise = TriggerDirection(1)
	// TriggerDirectionFall : triggered when market price falls
	TriggerDirectionFall = TriggerDirection(2)
)

// IsLeverage : Valid for spot only
type IsLeverage int

const (
	// IsLeverageFalse : false then spot trading
	IsLeverageFalse = IsLeverage(0)
	// IsLeverageTrue : true then margin trading
	IsLeverageTrue = IsLeverage(1)
)

// OrderFilter : Valid for spot only
type OrderFilter string

const (
	// OrderFilterOrder :
	OrderFilterOrder = OrderFilter("Order")
	// OrderFilterStopOrder :
	OrderFilterStopOrder = OrderFilter("StopOrder")
	// OrderFilterTpSlOrder :
	OrderFilterTpSlOrder = OrderFilter("tpslOrder")
)

// TriggerBy :
type TriggerBy string

const (
	// TriggerByLastPrice :
	TriggerByLastPrice = TriggerBy("LastPrice")
	// TriggerByIndexPrice :
	TriggerByIndexPrice = TriggerBy("IndexPrice")
	// TriggerByMarkPrice :
	TriggerByMarkPrice = TriggerBy("MarkPrice")
)

// PositionIdx :
type PositionIdx int

// PositionIdx :
const (
	PositionIdxOneWay    = PositionIdx(0)
	PositionIdxHedgeBuy  = PositionIdx(1)
	PositionIdxHedgeSell = PositionIdx(2)
)

// ContractType :
type ContractType string

// ContractType :
const (
	ContractTypeInversePerpetual = ContractType("InversePerpetual")
	ContractTypeLinearPerpetual  = ContractType("LinearPerpetual")
	ContractTypeInverseFutures   = ContractType("InverseFutures")
)

// InstrumentStatus :
type InstrumentStatus string

// InstrumentStatus :
const (
	// linear & inverse:
	InstrumentStatusPending  = InstrumentStatus("Pending")
	InstrumentStatusTrading  = InstrumentStatus("Trading")
	InstrumentStatusSettling = InstrumentStatus("Settling")
	InstrumentStatusClosed   = InstrumentStatus("Closed")

	// option
	InstrumentStatusWaitingOnline = InstrumentStatus("WAITING_ONLINE")
	InstrumentStatusOnline        = InstrumentStatus("ONLINE")
	InstrumentStatusDelivering    = InstrumentStatus("DELIVERING")
	InstrumentStatusOffline       = InstrumentStatus("OFFLINE")

	// spot
	InstrumentStatusAvailable = InstrumentStatus("1")
)

// OptionsType :
type OptionsType string

// OptionsType :
const (
	OptionsTypeCall = OptionsType("Call")
	OptionsTypePut  = OptionsType("Put")
)

// Innovation :
type Innovation string

// Innovation :
const (
	InnovationFalse = Innovation("0")
	InnovationTrue  = Innovation("1")
)

// PositionMode :
type PositionMode int

const (
	// PositionModeMergedSingle :
	PositionModeMergedSingle = PositionMode(0)
	// PositionModeBothSides :
	PositionModeBothSides = PositionMode(3)
)

// PositionMarginMode :
type PositionMarginMode int

const (
	// PositionMarginCross :
	PositionMarginCross = PositionMarginMode(0)
	// PositionMarginIsolated :
	PositionMarginIsolated = PositionMarginMode(1)
)

// ExecTypeV5 :
type ExecTypeV5 string

const (
	// ExecTypeV5Trade :
	ExecTypeV5Trade = ExecTypeV5("Trade")
	// ExecTypeV5AdlTrade : Auto-Deleveraging
	// https://www.bybit.com/en/help-center/article/Auto-Deleveraging-ADL
	ExecTypeV5AdlTrade = ExecTypeV5("AdlTrade")
	// ExecTypeV5Funding : Funding Fee
	// https://www.bybit.com/en/help-center/article/Introduction-to-Funding-Rate
	ExecTypeV5Funding = ExecTypeV5("Funding")
	// ExecTypeV5BustTrade : Takeover liquidation
	ExecTypeV5BustTrade = ExecTypeV5("BustTrade")
	// ExecTypeV5Delivery : USDC futures delivery; Position closed by contract delisted
	ExecTypeV5Delivery = ExecTypeV5("Delivery")
	// ExecTypeV5SessionSettlePnL :
	ExecTypeV5SessionSettlePnL = ExecTypeV5("SessionSettlePnL")
	// ExecTypeV5Settle :
	ExecTypeV5Settle = ExecTypeV5("Settle")
	// ExecTypeV5BlockTrade :
	ExecTypeV5BlockTrade = ExecTypeV5("BlockTrade")
	// ExecTypeV5MovePosition :
	ExecTypeV5MovePosition = ExecTypeV5("MovePosition")
	// ExecTypeV5UNKNOWN : May be returned by a classic account. Cannot query by this type
	ExecTypeV5UNKNOWN = ExecTypeV5("UNKNOWN")
)

// TransferStatusV5 :
type TransferStatusV5 string

const (
	// TransferStatusV5SUCCESS :
	TransferStatusV5SUCCESS = TransferStatusV5("SUCCESS")
	// TransferStatusV5PENDING :
	TransferStatusV5PENDING = TransferStatusV5("PENDING")
	// TransferStatusV5FAILED :
	TransferStatusV5FAILED = TransferStatusV5("FAILED")
)

// AccountTypeV5 :
type AccountTypeV5 string

const (
	// AccountTypeV5CONTRACT
	// Classic: Derivatives Account
	// UTA: 	Inverse Derivatives Account
	AccountTypeV5CONTRACT = AccountTypeV5("CONTRACT")
	// AccountTypeV5SPOT
	// Classic: Spot Account
	AccountTypeV5SPOT = AccountTypeV5("SPOT")
	// AccountTypeV5INVESTMENT
	// Classic: ByFi Account
	// Deprecated: this service is now offline
	AccountTypeV5INVESTMENT = AccountTypeV5("INVESTMENT")
	// AccountTypeV5OPTION
	// Classic: USDC Derivatives
	AccountTypeV5OPTION = AccountTypeV5("OPTION")
	// AccountTypeV5UNIFIED
	// UTA: Unified Trading Account
	AccountTypeV5UNIFIED = AccountTypeV5("UNIFIED")
	// AccountTypeV5FUND
	// Classic: Funding Account
	// UTA: 	Funding Account
	AccountTypeV5FUND = AccountTypeV5("FUND")
)

// UnifiedMarginStatus :
type UnifiedMarginStatus int

const (
	// UnifiedMarginStatusRegular : Regular account
	UnifiedMarginStatusRegular = UnifiedMarginStatus(1)
	// UnifiedMarginStatusUnifiedMargin : Unified margin account, it only trades linear perpetual and options.
	// Deprecated: Is not used anymore - Please ignore
	UnifiedMarginStatusUnifiedMargin = UnifiedMarginStatus(2)
	// UnifiedMarginStatusUnifiedTrade : Unified trade account, it can trade linear perpetual, options and spot
	UnifiedMarginStatusUnifiedTrade = UnifiedMarginStatus(3)
	// UnifiedMarginStatusUTAPro : UTA Pro, the pro version of Unified trade account
	UnifiedMarginStatusUTAPro = UnifiedMarginStatus(4)
)

// TransactionLogTypeV5 :
type TransactionLogTypeV5 string

const (
	TransactionLogTypeV5TRANSFERIN   = TransactionLogTypeV5("TRANSFER_IN")
	TransactionLogTypeV5TRANSFEROUT  = TransactionLogTypeV5("TRANSFER_OUT")
	TransactionLogTypeV5TRADE        = TransactionLogTypeV5("TRADE")
	TransactionLogTypeV5SETTLEMENT   = TransactionLogTypeV5("SETTLEMENT")
	TransactionLogTypeV5DELIVERY     = TransactionLogTypeV5("DELIVERY")
	TransactionLogTypeV5LIQUIDATION  = TransactionLogTypeV5("LIQUIDATION")
	TransactionLogTypeV5BONUS        = TransactionLogTypeV5("BONUS")
	TransactionLogTypeV5FEEREFUND    = TransactionLogTypeV5("FEE_REFUND")
	TransactionLogTypeV5INTEREST     = TransactionLogTypeV5("INTEREST")
	TransactionLogTypeV5CURRENCYBUY  = TransactionLogTypeV5("CURRENCY_BUY")
	TransactionLogTypeV5CURRENCYSELL = TransactionLogTypeV5("CURRENCY_SELL")
)

// InternalDepositStatusV5 :
type InternalDepositStatusV5 int

const (
	InternalDepositStatusV5Processing = InternalDepositStatusV5(1)
	InternalDepositStatusV5Success    = InternalDepositStatusV5(2)
	InternalDepositStatusV5Failed     = InternalDepositStatusV5(3)
)

// DepositStatusV5 :
type DepositStatusV5 int

const (
	DepositStatusV5Unknown       = DepositStatusV5(0)
	DepositStatusV5ToBeConfirmed = DepositStatusV5(1)
	DepositStatusV5Processing    = DepositStatusV5(2)
	DepositStatusV5Success       = DepositStatusV5(3)
	DepositStatusV5Failed        = DepositStatusV5(4)
)

type WithdrawTypeV5 int

const (
	WithdrawTypeOnChain  = WithdrawTypeV5(0)
	WithdrawTypeOffChain = WithdrawTypeV5(1)
	WithdrawTypeAll      = WithdrawTypeV5(2)
)

type WithdrawStatusV5 string

const (
	WithdrawStatusV5SecurityCheck       = WithdrawStatusV5("SecurityCheck")
	WithdrawStatusV5Pending             = WithdrawStatusV5("Pending")
	WithdrawStatusV5Success             = WithdrawStatusV5("success")
	WithdrawStatusV5CancelByUser        = WithdrawStatusV5("CancelByUser")
	WithdrawStatusV5Reject              = WithdrawStatusV5("Reject")
	WithdrawStatusV5Fail                = WithdrawStatusV5("Fail")
	WithdrawStatusV5BlockchainConfirmed = WithdrawStatusV5("BlockchainConfirmed")
)

type IsLowestRisk int

const (
	IsLowestRiskFalse = IsLowestRisk(0)
	IsLowestRiskTrue  = IsLowestRisk(1)
)

type CollateralSwitchV5 string

const (
	CollateralSwitchV5On  = CollateralSwitchV5("ON")
	CollateralSwitchV5Off = CollateralSwitchV5("OFF")
)

// AdlRankIndicator : Auto-deleverage rank indicator
type AdlRankIndicator int

const (
	AdlRankIndicator0 = AdlRankIndicator(0) // default value of empty position
	AdlRankIndicator1 = AdlRankIndicator(1)
	AdlRankIndicator2 = AdlRankIndicator(2)
	AdlRankIndicator3 = AdlRankIndicator(3)
	AdlRankIndicator4 = AdlRankIndicator(4)
	AdlRankIndicator5 = AdlRankIndicator(5)
)
