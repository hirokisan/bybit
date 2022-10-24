//go:build integrationtestfutureinverseperpetual

package integrationtestfutureinverseperpetual

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
		res, err := client.Future().InversePerpetual().Balance(bybit.CoinUSDT)
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-wallet-balance.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().Balance(bybit.CoinBTC)
		require.Error(t, err)
	})
}

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().OrderBook(bybit.SymbolFutureBTCUSD)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestListKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().ListKline(bybit.ListKlineParam{
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
	res, err := client.Future().InversePerpetual().Tickers(bybit.SymbolFutureBTCUSD)
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
	res, err := client.Future().InversePerpetual().TradingRecords(bybit.TradingRecordsParam{
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
	res, err := client.Future().InversePerpetual().Symbols()
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
	res, err := client.Future().InversePerpetual().IndexPriceKline(bybit.IndexPriceKlineParam{
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
	res, err := client.Future().InversePerpetual().OpenInterest(bybit.OpenInterestParam{
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
	res, err := client.Future().InversePerpetual().BigDeal(bybit.BigDealParam{
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
	res, err := client.Future().InversePerpetual().AccountRatio(bybit.AccountRatioParam{
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

func TestPremiumIndexKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().PremiumIndexKline(bybit.PremiumIndexKlineParam{
		Symbol:   bybit.SymbolFutureBTCUSD,
		Interval: bybit.Interval120,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	require.NoError(t, err)
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
		res, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolFutureBTCUSD,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         1,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order-create.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
		// clean
		{
			orderID := res.Result.OrderID
			_, err := client.Future().InversePerpetual().CancelOrder(bybit.CancelOrderParam{
				Symbol:  bybit.SymbolFutureBTCUSD,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 28383.5
		_, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
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

func TestListOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolFutureBTCUSD,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}
		{
			res, err := client.Future().InversePerpetual().ListOrder(bybit.ListOrderParam{
				Symbol: bybit.SymbolFutureBTCUSD,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-order-list.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InversePerpetual().CancelOrder(bybit.CancelOrderParam{
				Symbol:  bybit.SymbolFutureBTCUSD,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListOrder(bybit.ListOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestListPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual().ListPosition(bybit.SymbolFutureBTCUSD)
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-position-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListPosition(bybit.SymbolFutureBTCUSD)
		require.Error(t, err)
	})
}

func TestListPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual().ListPositions()
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-position-lists.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListPositions()
		require.Error(t, err)
	})
}

func TestCancelOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var orderID string
		{
			price := 28383.5
			res, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolFutureBTCUSD,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}

		res, err := client.Future().InversePerpetual().CancelOrder(bybit.CancelOrderParam{
			Symbol:  bybit.SymbolFutureBTCUSD,
			OrderID: &orderID,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelOrder(bybit.CancelOrderParam{})
		require.Error(t, err)
	})
}

func TestCancelAllOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			price := 10000.0
			_, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolFutureBTCUSD,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
		}
		res, err := client.Future().InversePerpetual().CancelAllOrder(bybit.CancelAllOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order-cancel-all.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelAllOrder(bybit.CancelAllOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestQueryOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
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

		res, err := client.Future().InversePerpetual().QueryOrder(bybit.QueryOrderParam{
			Symbol: symbol,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}

		{
			_, err := client.Future().InversePerpetual().CancelOrder(bybit.CancelOrderParam{
				Symbol:  symbol,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().QueryOrder(bybit.QueryOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestTradingStop(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		{
			_, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}

		{
			price := 20000.0
			res, err := client.Future().InversePerpetual().TradingStop(bybit.TradingStopParam{
				Symbol:     symbol,
				TakeProfit: &price,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-position-trading-stop.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InversePerpetual().CreateOrder(bybit.CreateOrderParam{
				Side:        bybit.SideSell,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().TradingStop(bybit.TradingStopParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestSaveLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			res, err := client.Future().InversePerpetual().SaveLeverage(bybit.SaveLeverageParam{
				Symbol:   bybit.SymbolFutureBTCUSD,
				Leverage: 2.0,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-position-leverage-save.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelOrder(bybit.CancelOrderParam{})
		require.Error(t, err)
	})
}

func TestCreateStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(bybit.CreateStopOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolFutureBTCUSD,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-stop-order-create.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
			stopOrderID = res.Result.StopOrderID
		}
		// clean
		{
			_, err := client.Future().InversePerpetual().CancelStopOrder(bybit.CancelStopOrderParam{
				Symbol:      bybit.SymbolFutureBTCUSD,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 10000.0
		_, err := client.Future().InversePerpetual().CreateStopOrder(bybit.CreateStopOrderParam{
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

func TestListStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var stopOrderID string
		symbol := bybit.SymbolFutureBTCUSD
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(bybit.CreateStopOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}

		// need to wait until the order status becode untriggered
		time.Sleep(10 * time.Second)

		status := bybit.OrderStatusUntriggered
		res, err := client.Future().InversePerpetual().ListStopOrder(bybit.ListStopOrderParam{
			Symbol:          symbol,
			StopOrderStatus: &status,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}

		// clean
		{
			_, err := client.Future().InversePerpetual().CancelStopOrder(bybit.CancelStopOrderParam{
				Symbol:      bybit.SymbolFutureBTCUSD,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListStopOrder(bybit.ListStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestCancelStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(bybit.CreateStopOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolFutureBTCUSD,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}
		res, err := client.Future().InversePerpetual().CancelStopOrder(bybit.CancelStopOrderParam{
			Symbol:      bybit.SymbolFutureBTCUSD,
			StopOrderID: &stopOrderID,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelStopOrder(bybit.CancelStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestCancelAllStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			price := 19400.5
			_, err := client.Future().InversePerpetual().CreateStopOrder(bybit.CreateStopOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolFutureBTCUSD,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}
		res, err := client.Future().InversePerpetual().CancelAllStopOrder(bybit.CancelAllStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order-cancel-all.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelAllStopOrder(bybit.CancelAllStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestQueryStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(bybit.CreateStopOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}

		res, err := client.Future().InversePerpetual().QueryStopOrder(bybit.QueryStopOrderParam{
			Symbol: symbol,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}

		{
			_, err := client.Future().InversePerpetual().CancelStopOrder(bybit.CancelStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().QueryStopOrder(bybit.QueryStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}
