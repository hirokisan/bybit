//go:build integrationtest

package integrationtest

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		price := 28383.5
		res, err := client.Account().CreateLinearOrder(bybit.CreateLinearOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolUSDTBTC,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
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
		_, err := client.Account().CreateLinearOrder(bybit.CreateLinearOrderParam{
			Side:        bybit.SideBuy,
			Symbol:      bybit.SymbolUSDTBTC,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: bybit.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListLinearPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Account().ListLinearPosition(bybit.SymbolUSDTBTC)
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/private-linear-position-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Account().ListLinearPosition(bybit.SymbolUSDTBTC)
		require.Error(t, err)
	})
}

func TestListLinearPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Account().ListLinearPositions()
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/private-linear-position-lists.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Account().ListLinearPositions()
		require.Error(t, err)
	})
}

func TestCancelLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()

		var orderID string
		{
			price := 47000.0
			res, err := client.Account().CreateLinearOrder(bybit.CreateLinearOrderParam{
				Side:        bybit.SideBuy,
				Symbol:      bybit.SymbolUSDTBTC,
				OrderType:   bybit.OrderTypeLimit,
				Qty:         0.001,
				TimeInForce: bybit.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
			orderID = res.Result.OrderID
		}
		res, err := client.Account().CancelLinearOrder(bybit.CancelLinearOrderParam{
			Symbol:  bybit.SymbolUSDTBTC,
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/private-linear-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Account().CancelLinearOrder(bybit.CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestSaveLinearLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Account().SaveLinearLeverage(bybit.SaveLinearLeverageParam{
			Symbol:       bybit.SymbolUSDTBTC,
			BuyLeverage:  2.0,
			SellLeverage: 2.0,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "leverage not modified", res.RetMsg)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Account().CancelLinearOrder(bybit.CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestLinearExecutionList(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Account().LinearExecutionList(bybit.LinearExecutionListParam{
			Symbol: bybit.SymbolUSDTBTC,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/private-linear-trade-execution-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Account().LinearExecutionList(bybit.LinearExecutionListParam{})
		require.Error(t, err)
	})
}

func TestAccountService_LinearCancelAllOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Account().LinearCancelAllOrder(bybit.LinearCancelAllParam{
			Symbol: bybit.SymbolUSDTBTC,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/private-linear-cancel-all-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Account().LinearCancelAllOrder(bybit.LinearCancelAllParam{})
		require.Error(t, err)
	})
}
