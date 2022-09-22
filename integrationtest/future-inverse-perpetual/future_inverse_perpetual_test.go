//go:build integrationtestfutureinverseperpetual

package integrationtestfutureinverseperpetual

import (
	"testing"
	"time"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestBalance(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual.Balance(bybit.CoinUSDT)
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/v2-private-wallet-balance.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual.Balance(bybit.CoinBTC)
		require.Error(t, err)
	})
}

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.OrderBook(bybit.SymbolInverseBTCUSD)
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestListKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.ListKline(bybit.ListKlineParam{
		Symbol:   bybit.SymbolInverseBTCUSD,
		Interval: bybit.Interval120,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-kline-list.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.Tickers(bybit.SymbolInverseBTCUSD)
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTradingRecords(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().InversePerpetual.TradingRecords(bybit.TradingRecordsParam{
		Symbol: bybit.SymbolInverseBTCUSD,
		Limit:  &limit,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-trading-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.Symbols()
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestIndexPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.IndexPriceKline(bybit.IndexPriceKlineParam{
		Symbol:   bybit.SymbolInverseBTCUSD,
		Interval: bybit.IntervalD,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.OpenInterest(bybit.OpenInterestParam{
		Symbol: bybit.SymbolInverseBTCUSD,
		Period: bybit.Period1h,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-open-interest.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestBigDeal(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.BigDeal(bybit.BigDealParam{
		Symbol: bybit.SymbolInverseBTCUSD,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-big-deal.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestAccountRatio(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().InversePerpetual.AccountRatio(bybit.AccountRatioParam{
		Symbol: bybit.SymbolInverseBTCUSD,
		Period: bybit.Period1h,
		Limit:  &limit,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-account-ratio.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestPremiumIndexKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual.PremiumIndexKline(bybit.PremiumIndexKlineParam{
		Symbol:   bybit.SymbolInverseBTCUSD,
		Interval: bybit.Interval120,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-premium-index-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestCreateOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		price := 28383.5
		res, err := client.Future().InversePerpetual.CreateOrder(bybit.CreateOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolInverseBTCUSD,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         1,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/v2-private-order-create.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
		// clean
		{
			orderID := res.Result.OrderID
			res, err := client.Future().InversePerpetual.CancelOrder(bybit.CancelOrderParam{
				Symbol:  bybit.SymbolInverseBTCUSD,
				OrderID: &orderID,
			})
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 28383.5
		_, err := client.Future().InversePerpetual.CreateOrder(bybit.CreateOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolInverseBTCUSD,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         1,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InversePerpetual.CreateOrder(bybit.CreateOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolInverseBTCUSD,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			{
				require.NoError(t, err)
				require.Equal(t, "OK", res.RetMsg)
			}
			orderID = res.Result.OrderID
		}
		res, err := client.Future().InversePerpetual.ListOrder(bybit.ListOrderParam{
			Symbol: bybit.SymbolInverseBTCUSD,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/v2-private-order-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
		// clean
		{
			res, err := client.Future().InversePerpetual.CancelOrder(bybit.CancelOrderParam{
				Symbol:  bybit.SymbolInverseBTCUSD,
				OrderID: &orderID,
			})
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 28383.5
		_, err := client.Future().InversePerpetual.CreateOrder(bybit.CreateOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolInverseBTCUSD,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         1,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual.ListPosition(bybit.SymbolInverseBTCUSD)
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/v2-private-position-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual.ListPosition(bybit.SymbolInverseBTCUSD)
		require.Error(t, err)
	})
}

func TestListPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual.ListPositions()
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/v2-private-position-lists.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual.ListPositions()
		require.Error(t, err)
	})
}

func TestCancelOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var orderID string
		{
			price := 28383.5
			res, err := client.Future().InversePerpetual.CreateOrder(bybit.CreateOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolInverseBTCUSD,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			{
				require.NoError(t, err)
				require.Equal(t, "OK", res.RetMsg)
			}
			orderID = res.Result.OrderID
		}
		res, err := client.Future().InversePerpetual.CancelOrder(bybit.CancelOrderParam{
			Symbol:  bybit.SymbolInverseBTCUSD,
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/v2-private-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual.CancelOrder(bybit.CancelOrderParam{})
		require.Error(t, err)
	})
}

func TestSaveLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			res, err := client.Future().InversePerpetual.SaveLeverage(bybit.SaveLeverageParam{
				Symbol:   bybit.SymbolInverseBTCUSD,
				Leverage: 2.0,
			})
			{
				require.NoError(t, err)
				require.Equal(t, "leverage not modified", res.RetMsg)
			}
			{
				goldenFilename := "./testdata/v2-private-position-leverage-save.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual.CancelOrder(bybit.CancelOrderParam{})
		require.Error(t, err)
	})
}
