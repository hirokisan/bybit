# bybit

bybit is an bybit client for the Go programing language.

## Usage

```
import "github.com/hirokisan/bybit"

client := bybit.NewClient().WithAuth("your api key", "your api secret")
res, err := client.Wallet().Balance(CoinBTC)
// do as you want
```

## Status

### Account Data Endpoints

- `/v2/private/order/create`
- `/v2/private/order/cancel`
- `/v2/private/order/list`

### Wallet Data Endpoints

- `/v2/private/wallet/balance`
