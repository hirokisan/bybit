# bybit

bybit is an bybit client for the Go programing language.

## Usage

```
import "github.com/hirokisan/bybit"

client := bybit.NewClient().WithAuth("your api key", "your api secret")
res, err := client.Wallet().Balance(bybit.CoinBTC)
// do as you want
```

## Status

### Market Data Endpoints

- `/v2/public/orderBook/L2`
- `/v2/public/kline/list`
- `/v2/public/tickers`
- `/v2/public/trading-records`
- `/v2/public/symbols`
- `/v2/public/liq-records`
- `/v2/public/mark-price-kline`
- `/v2/public/index-price-kline`
- `/v2/public/premium-index-kline`
- `/v2/public/open-interest`
- `/v2/public/big-deal`

### Account Data Endpoints

- `/v2/private/order/create`
- `/v2/private/order/cancel`
- `/v2/private/order/list`

### Wallet Data Endpoints

- `/v2/private/wallet/balance`
