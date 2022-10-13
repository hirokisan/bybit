//go:build integrationtestfutureusdtperpetual

package integrationtestfutureusdtperpetual

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
		res, err := client.Future().USDTPerpetual().Balance(bybit.CoinUSDT)
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
		_, err := client.Future().USDTPerpetual().Balance(bybit.CoinBTC)
		require.Error(t, err)
	})
}

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().OrderBook(bybit.SymbolFutureBTCUSD)
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/v2-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestListLinearKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().ListLinearKline(bybit.ListLinearKlineParam{
		Symbol:   bybit.SymbolFutureBTCUSDT,
		Interval: bybit.Interval120,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/public-linear-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().Tickers(bybit.SymbolFutureBTCUSD)
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/v2-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().Symbols()
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/v2-public-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().OpenInterest(bybit.OpenInterestParam{
		Symbol: bybit.SymbolFutureBTCUSD,
		Period: bybit.Period1h,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/v2-public-open-interest.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestBigDeal(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().BigDeal(bybit.BigDealParam{
		Symbol: bybit.SymbolFutureBTCUSD,
	})
	{
		require.NoError(t, err)
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
	res, err := client.Future().USDTPerpetual().AccountRatio(bybit.AccountRatioParam{
		Symbol: bybit.SymbolFutureBTCUSD,
		Period: bybit.Period1h,
		Limit:  &limit,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/v2-public-account-ratio.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestCreateLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		price := 28383.5
		res, err := client.Future().USDTPerpetual().CreateLinearOrder(bybit.CreateLinearOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolFutureBTCUSDT,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-order-create.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 28383.5
		_, err := client.Future().USDTPerpetual().CreateLinearOrder(bybit.CreateLinearOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolFutureBTCUSDT,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListLinearOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	symbol := bybit.SymbolFutureBTCUSDT

	var orderID string
	{
		price := 10000.0
		res, err := client.Future().USDTPerpetual().CreateLinearOrder(bybit.CreateLinearOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      symbol,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		{
			require.NoError(t, err)
		}
		orderID = res.Result.OrderID
		t.Log(res)
	}

	{
		orderStatus := bybit.OrderStatusNew
		res, err := client.Future().USDTPerpetual().ListLinearOrder(bybit.ListLinearOrderParam{
			Symbol:      symbol,
			OrderStatus: &orderStatus,
		})
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-order-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}

	{
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(bybit.CancelLinearOrderParam{
			Symbol:  symbol,
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
		}
	}
}

func TestListLinearPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().ListLinearPosition(bybit.SymbolFutureBTCUSDT)
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-position-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().ListLinearPosition(bybit.SymbolFutureBTCUSDT)
		require.Error(t, err)
	})
}

func TestListLinearPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().ListLinearPositions()
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-position-lists.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().ListLinearPositions()
		require.Error(t, err)
	})
}

func TestCancelLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()

		var orderID string
		{
			price := 47000.0
			res, err := client.Future().USDTPerpetual().CreateLinearOrder(bybit.CreateLinearOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolFutureBTCUSDT,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         0.001,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}
		res, err := client.Future().USDTPerpetual().CancelLinearOrder(bybit.CancelLinearOrderParam{
			Symbol:  bybit.SymbolFutureBTCUSDT,
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(bybit.CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestSaveLinearLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		_, err := client.Future().USDTPerpetual().SaveLinearLeverage(bybit.SaveLinearLeverageParam{
			Symbol:       bybit.SymbolFutureBTCUSDT,
			BuyLeverage:  2.0,
			SellLeverage: 2.0,
		})
		{
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(bybit.CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestLinearExecutionList(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().LinearExecutionList(bybit.LinearExecutionListParam{
			Symbol: bybit.SymbolFutureBTCUSDT,
		})
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-trade-execution-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().LinearExecutionList(bybit.LinearExecutionListParam{})
		require.Error(t, err)
	})
}

func TestAccountService_LinearCancelAllOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().LinearCancelAllOrder(bybit.LinearCancelAllParam{
			Symbol: bybit.SymbolFutureBTCUSDT,
		})
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-cancel-all-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().LinearCancelAllOrder(bybit.LinearCancelAllParam{})
		require.Error(t, err)
	})
}

func TestQueryLinearOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	symbol := bybit.SymbolFutureBTCUSDT

	var orderID string
	{
		price := 10000.0
		res, err := client.Future().USDTPerpetual().CreateLinearOrder(bybit.CreateLinearOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      symbol,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		{
			require.NoError(t, err)
		}
		orderID = res.Result.OrderID
	}

	{
		res, err := client.Future().USDTPerpetual().QueryLinearOrder(bybit.QueryLinearOrderParam{
			Symbol: symbol,
		})
		{
			require.NoError(t, err)
		}
		{
			goldenFilename := "./testdata/private-linear-order-search.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}

	{
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(bybit.CancelLinearOrderParam{
			Symbol:  symbol,
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
		}
	}
}
