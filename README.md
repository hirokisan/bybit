[![Go Report Card](https://goreportcard.com/badge/github.com/hirokisan/bybit)](https://goreportcard.com/report/github.com/hirokisan/bybit)
[![golangci-lint](https://github.com/hirokisan/bybit/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/hirokisan/bybit/actions/workflows/golangci-lint.yml)
[![test](https://github.com/hirokisan/bybit/actions/workflows/test.yml/badge.svg)](https://github.com/hirokisan/bybit/actions/workflows/test.yml)

# bybit

bybit is an bybit client for the Go programming language.

## Usage

### REST API

```golang
import "github.com/hirokisan/bybit/v2"

client := bybit.NewClient().WithAuth("your api key", "your api secret")
res, err := client.Future().InversePerpetual().Balance(bybit.CoinBTC)
// do as you want
```

### WebSocket API

for single use
```golang
import "github.com/hirokisan/bybit/v2"

wsClient := bybit.NewWebsocketClient()
svc, err := wsClient.Spot().V1().PublicV1()
if err != nil {
	return err
}
_, err = svc.SubscribeTrade(bybit.SymbolSpotBTCUSDT, func(response bybit.SpotWebsocketV1PublicV1TradeResponse) error {
	// do as you want
})
if err != nil {
	return err
}
svc.Start(context.Background())
```

for multiple use
```golang
import "github.com/hirokisan/bybit/v2"

wsClient := bybit.NewWebsocketClient()

executors := []bybit.WebsocketExecutor{}

svcRoot := wsClient.Spot().V1()
{
	svc, err := svcRoot.PublicV1()
	if err != nil {
		return err
	}
	_, err = svc.SubscribeTrade(bybit.SymbolSpotBTCUSDT, func(response bybit.SpotWebsocketV1PublicV1TradeResponse) error {
		// do as you want
	})
	if err != nil {
		return err
	}
	executors = append(executors, svc)
}
{
	svc, err := svcRoot.PublicV2()
	if err != nil {
		return err
	}
	_, err = svc.SubscribeTrade(bybit.SymbolSpotBTCUSDT, func(response bybit.SpotWebsocketV1PublicV2TradeResponse) error {
		// do as you want
	})
	if err != nil {
		return err
	}
	executors = append(executors, svc)
}

wsClient.Start(context.Background(), executors)
```

V5 usage
```golang
import "github.com/hirokisan/bybit/v2"

wsClient := bybit.NewWebsocketClient().WithBaseURL("wss://stream-testnet.bybit.com").WithAuth("key", "secret")
svc, err := wsClient.V5().Private()
if err != nil {
	// handle dialing error
}

err = svc.Subscribe()
if err != nil {
	// handle subscription error
}

err = svc.SubscribePosition(func(position bybit.V5WebsocketPrivatePositionResponse) error {
	// handle new position information
})
if err != nil {
	// handle registration error
}

errHandler := func(isWebsocketClosed bool, err error) {
	// Connection issue (timeout, etc.).

	// At this point, the connection is dead and you must handle the reconnection yourself
}

err = svc.Start(context.Background(), errHandler)
if err != nil {
	// handle reconnection (ping issue, etc.). Probably can be ignored as the errHandler would be notified too
}
```

## Implemented

The following API endpoints have been implemented

### REST API V5

#### Market

- [`/v5/market/kline` Get Kline](https://bybit-exchange.github.io/docs/v5/market/kline)
- [`/v5/market/mark-price-kline` Get Mark Price Kline](https://bybit-exchange.github.io/docs/v5/market/mark-kline)
- [`/v5/market/index-price-kline` Get Index Price Kline](https://bybit-exchange.github.io/docs/v5/market/index-kline)
- [`/v5/market/premium-index-price-kline` Get Premium Index Price Kline](https://bybit-exchange.github.io/docs/v5/market/preimum-index-kline)
- [`/v5/market/instruments-info` Get Instruments Info](https://bybit-exchange.github.io/docs/v5/market/instrument)
- [`/v5/market/orderbook` Get Orderbook](https://bybit-exchange.github.io/docs/v5/market/orderbook)
- [`/v5/market/tickers` Get Tickers](https://bybit-exchange.github.io/docs/v5/market/tickers)
- [`/v5/market/funding/history` Get Funding Rate History](https://bybit-exchange.github.io/docs/v5/market/history-fund-rate)
- [`/v5/market/recent-trade` Get Public Trading History](https://bybit-exchange.github.io/docs/v5/market/recent-trade)
- [`/v5/market/open-interest` Get Open Interest](https://bybit-exchange.github.io/docs/v5/market/open-interest)
- [`/v5/market/historical-volatility` Get Historical Volatility](https://bybit-exchange.github.io/docs/v5/market/iv)

#### Position

- [`/v5/position/list` Get Position Info](https://bybit-exchange.github.io/docs/v5/position)
- [`/v5/position/set-leverage` Set Leverage](https://bybit-exchange.github.io/docs/v5/position/leverage)
- [`/v5/position/trading-stop` Set Trading Stop](https://bybit-exchange.github.io/docs/v5/position/trading-stop)
- [`/v5/position/switch-mode` Switch Position Mode](https://bybit-exchange.github.io/docs/v5/position/position-mode)
- [`/v5/position/set-tpsl-mode` Set TP/SL Mode](https://bybit-exchange.github.io/docs/v5/position/tpsl-mode)
- [`/v5/position/closed-pnl` Get Closed PnL](https://bybit-exchange.github.io/docs/v5/position/close-pnl)

#### Order

- [`/v5/order/create` Place Order](https://bybit-exchange.github.io/docs/v5/order/create-order)
- [`/v5/order/amend` Amend Order](https://bybit-exchange.github.io/docs/v5/order/amend-order)
- [`/v5/order/cancel` Cancel Order](https://bybit-exchange.github.io/docs/v5/order/cancel-order)
- [`/v5/order/realtime` Get Open Orders](https://bybit-exchange.github.io/docs/v5/order/open-order)
- [`/v5/order/cancel-all` Cancel All Orders](https://bybit-exchange.github.io/docs/v5/order/cancel-all)

#### Account

- [`/v5/account/wallet-balance` Get Wallet Balance](https://bybit-exchange.github.io/docs/v5/account/wallet-balance)
- [`/v5/account/account-info` Get Account Info](https://bybit-exchange.github.io/docs/v5/account/account-info)
- [`/v5/account/transaction-log` Get Transaction Log](https://bybit-exchange.github.io/docs/v5/account/transaction-log)

#### Asset

- [`/v5/asset/transfer/query-inter-transfer-list` Get Internal Transfer Records](https://bybit-exchange.github.io/docs/v5/asset/inter-transfer-list)
- [`/v5/asset/deposit/query-record` Get Deposit Records](https://bybit-exchange.github.io/docs/v5/asset/deposit-record)
- [`/v5/asset/deposit/query-internal-record` Get Internal Deposit Records](https://bybit-exchange.github.io/docs/v5/asset/internal-deposit-record)

#### User

- [`/v5/user/query-api` Get API Key Information](https://bybit-exchange.github.io/docs/v5/user/apikey-info)


### REST API

#### [Derivatives Unified Margin](https://bybit-exchange.github.io/docs/derivativesV3/unified_margin)

##### Market Data Endpoints

- `/derivatives/v3/public/order-book/L2` Get Order Book
- `/derivatives/v3/public/kline` Get Kline
- `/derivatives/v3/public/tickers` Get Latest Information For Symbol
- `/derivatives/v3/public/instruments-info` Get Instrument Info
- `/derivatives/v3/public/mark-price-kline` Get Mark Price Kline
- `/derivatives/v3/public/index-price-kline` Get Index Price Kline

#### [Derivatives Contract](https://bybit-exchange.github.io/docs/derivativesV3/contract)

##### Market Data Endpoints

- `/derivatives/v3/public/order-book/L2` Get Order Book
- `/derivatives/v3/public/kline` Get Kline
- `/derivatives/v3/public/tickers` Get Latest Information For Symbol
- `/derivatives/v3/public/instruments-info` Get Instrument Info
- `/derivatives/v3/public/mark-price-kline` Get Mark Price Kline
- `/derivatives/v3/public/index-price-kline` Get Index Price Kline

#### [Inverse Perpetual](https://bybit-exchange.github.io/docs/futuresV2/inverse)

##### Market Data Endpoints

- `/v2/public/orderBook/L2` Order Book
- `/v2/public/kline/list` Query Kline
- `/v2/public/tickers` Latest Information for Symbol
- `/v2/public/trading-records` Public Trading Records
- `/v2/public/symbols` Query Symbol
- `/v2/public/mark-price-kline` Query Mark Price Kline
- `/v2/public/index-price-kline` Query Index Price Kline
- `/v2/public/premium-index-kline` Query Premium Index Kline
- `/v2/public/open-interest` Open Interest
- `/v2/public/big-deal` Latest Big Deal
- `/v2/public/account-ratio` Long-Short Ratio

##### Account Data Endpoints

- `/v2/private/order/create` Place Active Order
- `/v2/private/order/list` Get Active Order
- `/v2/private/order/cancel` Cancel Active Order
- `/v2/private/order/cancelAll` Cancel All Active Orders
- `/v2/private/order` Query Active Order (real-time)
- `/v2/private/stop-order/create` Place Conditional Order
- `/v2/private/stop-order/list` Get Conditional Order
- `/v2/private/stop-order/cancel` Cancel Conditional Order
- `/v2/private/stop-order/cancelAll` Cancel All Conditional Orders
- `/v2/private/stop-order` Query Conditional Order (real-time)
- `/v2/private/position/list` My Position
- `/v2/private/position/trading-stop` Set Trading-Stop
- `/v2/private/position/leverage/save` Set Leverage
- `/v2/private/account/api-key` API Key info

##### Wallet Data Endpoints

- `/v2/private/wallet/balance` Get Wallet Balance

#### [USDT Perpetual](https://bybit-exchange.github.io/docs/futuresV2/linear)

##### Market Data Endpoints

- `/v2/public/orderBook/L2` Order Book
- `/public/linear/kline` Query Kline
- `/v2/public/tickers` Latest Information for Symbol
- `/v2/public/symbols` Query Symbol
- `/v2/public/open-interest` Open Interest
- `/v2/public/big-deal` Latest Big Deal
- `/v2/public/account-ratio` Long-Short Ratio

##### Account Data Endpoints

- `/private/linear/order/create` Place Active Order
- `/private/linear/order/list` Get Active Order
- `/private/linear/order/cancel` Cancel Active Order
- `/private/linear/order/cancel-all` Cancel All Active Orders
- `/private/linear/order/replace` Replace Active Order
- `/private/linear/order/search` Query Active Order (real-time)
- `/private/linear/stop-order/create` Place Conditional Order
- `/private/linear/stop-order/list` Get Conditional Order
- `/private/linear/stop-order/cancel` Cancel Conditional Order
- `/private/linear/stop-order/cancel-all` Cancel All Conditional Orders
- `/private/linear/stop-order/search` Query Conditional Order (real-time)
- `/private/linear/position/list` My Position
- `/private/linear/position/set-leverage` Set Leverage
- `/private/linear/position/trading-stop` Set Trading-Stop
- `/private/linear/trade/execution/list` User Trade Records
- `/v2/private/account/api-key` API Key info

##### Wallet Data Endpoints

- `/v2/private/wallet/balance` Get Wallet Balance

#### [Inverse Future](https://bybit-exchange.github.io/docs/futuresV2/inverse_futures)

##### Market Data Endpoints

- `/v2/public/orderBook/L2` Order Book
- `/v2/public/kline/list` Query Kline
- `/v2/public/tickers` Latest Information for Symbol
- `/v2/public/trading-records` Public Trading Records
- `/v2/public/symbols` Query Symbol
- `/v2/public/mark-price-kline` Query Index Price Kline
- `/v2/public/index-price-kline` Query Index Price Kline
- `/v2/public/open-interest` Open Interest
- `/v2/public/big-deal` Latest Big Deal
- `/v2/public/account-ratio` Long-Short Ratio

##### Account Data Endpoints

- `/futures/private/order/create` Place Active Order
- `/futures/private/order/list` Get Active Order
- `/futures/private/order/cancel` Cancel Active Order
- `/futures/private/order/cancelAll` Cancel All Active Orders
- `/futures/private/order` Query Active Order (real-time)
- `/futures/private/stop-order/create` Place Conditional Order
- `/futures/private/stop-order/list` Get Conditional Order
- `/futures/private/stop-order/cancel` Cancel Conditional Order
- `/futures/private/stop-order/cancelAll` Cancel All Conditional Orders
- `/futures/private/stop-order` Query Conditional Order (real-time)
- `/futures/private/position/list` My Position
- `/futures/private/position/trading-stop` Set Trading-Stop
- `/futures/private/position/leverage/save` Set Leverage
- `/v2/private/account/api-key` API Key info

##### Wallet Data Endpoints

- `/v2/private/wallet/balance` Get Wallet Balance

#### [Spot v1](https://bybit-exchange.github.io/docs/spot/v1)

##### Market Data Endpoints

- `/spot/v1/symbols` Query Symbol
- `/spot/quote/v1/depth` Order Book
- `/spot/quote/v1/depth/merged` Merged Order Book
- `/spot/quote/v1/trades` Public Trading Records
- `/spot/quote/v1/kline` Query Kline
- `/spot/quote/v1/ticker/24hr` Latest Information for Symbol
- `/spot/quote/v1/ticker/price` Last Traded Price
- `/spot/quote/v1/ticker/book_ticker` Best Bid/Ask Price

##### Account Data Endpoints

- `/spot/v1/order`
  - Place Active Order
  - Get Active Order
  - Cancel Active Order
  - Fast Cancel Active Order
- `/spot/v1/order/fast` Fast Cancel Active Order
- `/spot/order/batch-cancel` Batch Cancel Active Order
- `/spot/order/batch-fast-cancel` Batch Fast Cancel Active Order
- `/spot/order/batch-cancel-by-ids` Batch Cancel Active Order By IDs
- `/spot/v1/open-orders` Open Orders

##### Wallet Data Endpoints

- `/spot/v1/account` Get Wallet Balance

### WebSocket API

#### Public Topics V5

- [Orderbook](https://bybit-exchange.github.io/docs/v5/websocket/public/orderbook)
- [Kline](https://bybit-exchange.github.io/docs/v5/websocket/public/kline)
- [Ticker](https://bybit-exchange.github.io/docs/v5/websocket/public/ticker)

#### Private Topics V5

- [Position](https://bybit-exchange.github.io/docs/v5/websocket/private/position)
- [Order](https://bybit-exchange.github.io/docs/v5/websocket/private/order)
- [Wallet](https://bybit-exchange.github.io/docs/v5/websocket/private/wallet)

#### [Spot v1](https://bybit-exchange.github.io/docs/spot/v1/#t-websocket)

##### Public Topics

- trade

##### Public Topics V2

- trade

##### Private Topics

- outboundAccountInfo

## Integration Tests

There are tests so that we can get to know the changes of bybit api response.

See below

- [REST API Integration Test](./integrationtest/README.md)
- [Websocket API Integration Test](./integrationtest-ws/README.md)

## Contributing

I would like to cover Bybit API and contributions are always welcome. The calling pattern is established, so adding new methods is relatively straightforward. See some PRs like https://github.com/hirokisan/bybit/pull/44.

To submit issues, PRs, and every other help is welcome.
