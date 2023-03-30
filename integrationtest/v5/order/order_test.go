//go:build integrationtestv5order

package integrationtestv5order

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	price := "10000.0"
	res, err := client.V5().Order().CreateOrder(bybit.V5CreateOrderParam{
		Category:  bybit.CategoryV5Spot,
		Symbol:    bybit.SymbolV5BTCUSDT,
		Side:      bybit.SideBuy,
		OrderType: bybit.OrderTypeLimit,
		Qty:       "0.01",
		Price:     &price,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-create-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestAmendOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	category := bybit.CategoryV5Linear
	symbol := bybit.SymbolV5BTCUSDT
	var orderID string
	{
		{
			price := "10000.0"
			res, err := client.V5().Order().CreateOrder(bybit.V5CreateOrderParam{
				Category:  category,
				Symbol:    symbol,
				Side:      bybit.SideBuy,
				OrderType: bybit.OrderTypeLimit,
				Qty:       "0.01",
				Price:     &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}
	}

	var amendedOrderID string
	{
		price := "9000.0"
		res, err := client.V5().Order().AmendOrder(bybit.V5AmendOrderParam{
			Category: category,
			Symbol:   symbol,
			OrderID:  &orderID,
			Price:    &price,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v5-amend-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
		amendedOrderID = res.Result.OrderID
	}

	{
		_, err := client.V5().Order().CancelOrder(bybit.V5CancelOrderParam{
			Category: category,
			Symbol:   symbol,
			OrderID:  &amendedOrderID,
		})
		require.NoError(t, err)
	}
}

func TestCancelOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	var orderID string
	category := bybit.CategoryV5Spot
	symbol := bybit.SymbolV5BTCUSDT
	{
		price := "10000.0"
		res, err := client.V5().Order().CreateOrder(bybit.V5CreateOrderParam{
			Category:  category,
			Symbol:    symbol,
			Side:      bybit.SideBuy,
			OrderType: bybit.OrderTypeLimit,
			Qty:       "0.01",
			Price:     &price,
		})
		require.NoError(t, err)
		orderID = res.Result.OrderID
	}

	res, err := client.V5().Order().CancelOrder(bybit.V5CancelOrderParam{
		Category: category,
		Symbol:   symbol,
		OrderID:  &orderID,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-cancel-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetOpenOrders(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	var orderID string
	category := bybit.CategoryV5Spot
	symbol := bybit.SymbolV5BTCUSDT
	{
		price := "10000.0"
		res, err := client.V5().Order().CreateOrder(bybit.V5CreateOrderParam{
			Category:  category,
			Symbol:    symbol,
			Side:      bybit.SideBuy,
			OrderType: bybit.OrderTypeLimit,
			Qty:       "0.01",
			Price:     &price,
		})
		require.NoError(t, err)
		orderID = res.Result.OrderID
	}

	res, err := client.V5().Order().GetOpenOrders(bybit.V5GetOpenOrdersParam{
		Category: category,
		Symbol:   &symbol,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-get-open-orders.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}

	{
		_, err := client.V5().Order().CancelOrder(bybit.V5CancelOrderParam{
			Category: category,
			Symbol:   symbol,
			OrderID:  &orderID,
		})
		require.NoError(t, err)
	}
}

func TestCancelAllOrders(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	category := bybit.CategoryV5Spot
	symbol := bybit.SymbolV5BTCUSDT
	{
		price := "10000.0"
		_, err := client.V5().Order().CreateOrder(bybit.V5CreateOrderParam{
			Category:  category,
			Symbol:    symbol,
			Side:      bybit.SideBuy,
			OrderType: bybit.OrderTypeLimit,
			Qty:       "0.01",
			Price:     &price,
		})
		require.NoError(t, err)
	}

	res, err := client.V5().Order().CancelAllOrders(bybit.V5CancelAllOrdersParam{
		Category: category,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-cancel-all-orders.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
