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

func TestCancelAllFuturesOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		{
			price := 10000.0
			_, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
		}
		{
			res, err := client.Future().InverseFuture().CancelAllFuturesOrder(bybit.CancelAllFuturesOrderParam{
				Symbol: symbol,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-order-cancel-all.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().CancelAllFuturesOrder(bybit.CancelAllFuturesOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestQueryFuturesOrder(t *testing.T) {
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
			res, err := client.Future().InverseFuture().QueryFuturesOrder(bybit.QueryFuturesOrderParam{
				Symbol:  symbol,
				OrderID: &orderID,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-order.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
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
		_, err := client.Future().InverseFuture().QueryFuturesOrder(bybit.QueryFuturesOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestCreateFuturesStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InverseFuture().CreateFuturesStopOrder(bybit.CreateFuturesStopOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-stop-order-create.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
			stopOrderID = res.Result.StopOrderID
		}
		// clean
		{
			_, err := client.Future().InverseFuture().CancelFuturesStopOrder(bybit.CancelFuturesStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 19400.5
		_, err := client.Future().InverseFuture().CreateFuturesStopOrder(bybit.CreateFuturesStopOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolFutureBTCUSD,
			OrderType:   bybit.OrderTypeMarket,
			Qty:         1,
			BasePrice:   price,
			StopPx:      price + 200,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
		})
		require.Error(t, err)
	})
}

func TestListFuturesStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSDH23
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InverseFuture().CreateFuturesStopOrder(bybit.CreateFuturesStopOrderParam{
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
		{
			status := bybit.OrderStatusUntriggered
			res, err := client.Future().InverseFuture().ListFuturesStopOrder(bybit.ListFuturesStopOrderParam{
				Symbol:          symbol,
				StopOrderStatus: &status,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-stop-order-list.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InverseFuture().CancelFuturesStopOrder(bybit.CancelFuturesStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().ListFuturesStopOrder(bybit.ListFuturesStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSDH23,
		})
		require.Error(t, err)
	})
}

func TestCancelFuturesStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InverseFuture().CreateFuturesStopOrder(bybit.CreateFuturesStopOrderParam{
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
		// clean
		{
			res, err := client.Future().InverseFuture().CancelFuturesStopOrder(bybit.CancelFuturesStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-stop-order-cancel.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().CancelFuturesStopOrder(bybit.CancelFuturesStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestCancelAllFuturesStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSD
		{
			price := 19400.5
			_, err := client.Future().InverseFuture().CreateFuturesStopOrder(bybit.CreateFuturesStopOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}
		// clean
		{
			res, err := client.Future().InverseFuture().CancelAllFuturesStopOrder(bybit.CancelAllFuturesStopOrderParam{
				Symbol: symbol,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-stop-order-cancel-all.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().CancelAllFuturesStopOrder(bybit.CancelAllFuturesStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestQueryFuturesStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSDH23
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InverseFuture().CreateFuturesStopOrder(bybit.CreateFuturesStopOrderParam{
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
		{
			res, err := client.Future().InverseFuture().QueryFuturesStopOrder(bybit.QueryFuturesStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-stop-order.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		{
			_, err := client.Future().InverseFuture().CancelFuturesStopOrder(bybit.CancelFuturesStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().QueryFuturesStopOrder(bybit.QueryFuturesStopOrderParam{
			Symbol: bybit.SymbolFutureBTCUSDH23,
		})
		require.Error(t, err)
	})
}

func TestListFuturesPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSDH23
		{
			_, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      symbol,
				OrderType:   bybit.OrderTypeMarket,
				Qty:         1,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}
		{
			res, err := client.Future().InverseFuture().ListFuturesPositions(symbol)
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-position-list.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
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
		_, err := client.Future().InverseFuture().ListFuturesPositions(bybit.SymbolFutureBTCUSDH23)
		require.Error(t, err)
	})
}

func TestFuturesTradingStop(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := bybit.SymbolFutureBTCUSDH23
		{
			_, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
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
			res, err := client.Future().InverseFuture().FuturesTradingStop(bybit.FuturesTradingStopParam{
				Symbol:     symbol,
				TakeProfit: &price,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-position-trading-stop.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InverseFuture().CreateFuturesOrder(bybit.CreateFuturesOrderParam{
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
		_, err := client.Future().InverseFuture().FuturesTradingStop(bybit.FuturesTradingStopParam{
			Symbol: bybit.SymbolFutureBTCUSDT,
		})
		require.Error(t, err)
	})
}

func TestFuturesSaveLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			res, err := client.Future().InverseFuture().FuturesSaveLeverage(bybit.FuturesSaveLeverageParam{
				Symbol:       bybit.SymbolFutureBTCUSDH23,
				BuyLeverage:  10.0,
				SellLeverage: 10.0,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/futures-private-position-leverage-save.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InverseFuture().FuturesSaveLeverage(bybit.FuturesSaveLeverageParam{
			Symbol:       bybit.SymbolFutureBTCUSD,
			BuyLeverage:  10.0,
			SellLeverage: 10.0,
		})
		require.Error(t, err)
	})
}
