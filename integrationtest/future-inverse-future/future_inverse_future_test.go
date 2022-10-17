//go:build integrationtestfutureinversefuture

package integrationtestfutureinversefuture

import (
	"testing"
	"time"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestBalance(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InverseFuture().Balance(bybit.CoinUSDT)
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-wallet-balance.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().Balance(bybit.CoinBTC)
		require.Error(t, err)
	})
}

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().OrderBook(bybit.SymbolFutureBTCUSD)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestListKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().ListKline(bybit.ListKlineParam{
		Symbol:   bybit.SymbolFutureBTCUSD,
		Interval: bybit.Interval120,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-kline-list.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().Tickers(bybit.SymbolFutureBTCUSD)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTradingRecords(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().InverseFuture().TradingRecords(bybit.TradingRecordsParam{
		Symbol: bybit.SymbolFutureBTCUSD,
		Limit:  &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-trading-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().Symbols()
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestMarkPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().MarkPriceKline(bybit.MarkPriceKlineParam{
		Symbol:   bybit.SymbolFutureBTCUSD,
		Interval: bybit.IntervalD,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-mark-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestIndexPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().IndexPriceKline(bybit.IndexPriceKlineParam{
		Symbol:   bybit.SymbolFutureBTCUSD,
		Interval: bybit.IntervalD,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().OpenInterest(bybit.OpenInterestParam{
		Symbol: bybit.SymbolFutureBTCUSD,
		Period: bybit.Period1h,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-open-interest.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestBigDeal(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().BigDeal(bybit.BigDealParam{
		Symbol: bybit.SymbolFutureBTCUSD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-big-deal.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestAccountRatio(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().InverseFuture().AccountRatio(bybit.AccountRatioParam{
		Symbol: bybit.SymbolFutureBTCUSD,
		Period: bybit.Period1h,
		Limit:  &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-account-ratio.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestCreateFuturesOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-order-create.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
			orderID = res.Result.OrderID
		}
		// clean
		{
			_, err := client.Future().InverseFuture().CancelFuturesOrder(bybit.CancelFuturesOrderParam{
				Symbol:  symbol,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 10000.0
		_, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolFutureBTCUSD,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         1,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListFuturesOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}

		// need to wait until the order status become new
		time.Sleep(10 * time.Second)

		{
			status := bybit.OrderStatusNew
			res, err := client.Future().InverseFuture().ListFuturesOrder(bybit.ListFuturesOrderParam{
				Symbol:      symbol,
				OrderStatus: &status,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-order-list.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InverseFuture().CancelFuturesOrder(bybit.CancelFuturesOrderParam{
				Symbol:  symbol,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().ListFuturesOrder(bybit.ListFuturesOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestCancelFuturesOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}
		{
			res, err := client.Future().InverseFuture().CancelFuturesOrder(bybit.CancelFuturesOrderParam{
				Symbol:  symbol,
				OrderID: &orderID,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-order-cancel.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().CancelFuturesOrder(bybit.CancelFuturesOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}
