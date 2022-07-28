//go:build integrationtest

package integrationtest

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		price := 28383.5
		res, err := client.Account().CreateOrder(bybit.CreateOrderParam{
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
			res, err := client.Account().CancelOrder(bybit.CancelOrderParam{
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
		_, err := client.Account().CreateOrder(bybit.CreateOrderParam{
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
			res, err := client.Account().CreateOrder(bybit.CreateOrderParam{
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
		res, err := client.Account().ListOrder(bybit.ListOrderParam{
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
			res, err := client.Account().CancelOrder(bybit.CancelOrderParam{
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
		_, err := client.Account().CreateOrder(bybit.CreateOrderParam{
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
		res, err := client.Account().ListPosition(bybit.SymbolInverseBTCUSD)
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
		_, err := client.Account().ListPosition(bybit.SymbolInverseBTCUSD)
		require.Error(t, err)
	})
}

func TestListPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Account().ListPositions()
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
		_, err := client.Account().ListPositions()
		require.Error(t, err)
	})
}

func TestCancelOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var orderID string
		{
			price := 28383.5
			res, err := client.Account().CreateOrder(bybit.CreateOrderParam{
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
		res, err := client.Account().CancelOrder(bybit.CancelOrderParam{
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
		_, err := client.Account().CancelOrder(bybit.CancelOrderParam{})
		require.Error(t, err)
	})
}

func TestSaveLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			res, err := client.Account().SaveLeverage(bybit.SaveLeverageParam{
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
		_, err := client.Account().CancelOrder(bybit.CancelOrderParam{})
		require.Error(t, err)
	})
}
