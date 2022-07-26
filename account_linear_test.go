package bybit

import (
	"testing"

	"github.com/hirokisan/bybit/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := NewTestClient().WithAuthFromEnv()
		price := 28383.5
		res, err := client.Account().CreateLinearOrder(CreateLinearOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolUSDTBTC,
			OrderType:   OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: TimeInForceGoodTillCancel,
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
		client := NewTestClient()
		price := 28383.5
		_, err := client.Account().CreateLinearOrder(CreateLinearOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolUSDTBTC,
			OrderType:   OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListLinearPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := NewTestClient().WithAuthFromEnv()
		res, err := client.Account().ListLinearPosition(SymbolUSDTBTC)
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
		client := NewTestClient()
		_, err := client.Account().ListLinearPosition(SymbolUSDTBTC)
		require.Error(t, err)
	})
}

func TestListLinearPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := NewTestClient().WithAuthFromEnv()
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
		client := NewTestClient()
		_, err := client.Account().ListLinearPositions()
		require.Error(t, err)
	})
}

func TestCancelLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := NewTestClient().WithAuthFromEnv()

		var orderID string
		{
			price := 47000.0
			res, err := client.Account().CreateLinearOrder(CreateLinearOrderParam{
				Side:        SideBuy,
				Symbol:      SymbolUSDTBTC,
				OrderType:   OrderTypeLimit,
				Qty:         0.001,
				TimeInForce: TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
			orderID = res.Result.OrderID
		}
		res, err := client.Account().CancelLinearOrder(CancelLinearOrderParam{
			Symbol:  SymbolUSDTBTC,
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
		client := NewTestClient()
		_, err := client.Account().CancelLinearOrder(CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestSaveLinearLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := NewTestClient().WithAuthFromEnv()
		res, err := client.Account().SaveLinearLeverage(SaveLinearLeverageParam{
			Symbol:       SymbolUSDTBTC,
			BuyLeverage:  2.0,
			SellLeverage: 2.0,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "leverage not modified", res.RetMsg)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := NewTestClient()
		_, err := client.Account().CancelLinearOrder(CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestLinearExecutionList(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := NewTestClient().WithAuthFromEnv()
		res, err := client.Account().LinearExecutionList(LinearExecutionListParam{
			Symbol: SymbolUSDTBTC,
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
		client := NewTestClient()
		_, err := client.Account().LinearExecutionList(LinearExecutionListParam{})
		require.Error(t, err)
	})
}

func TestAccountService_LinearCancelAllOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := NewTestClient().WithAuthFromEnv()
		res, err := client.Account().LinearCancelAllOrder(LinearCancelAllParam{
			Symbol: SymbolUSDTBTC,
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
		client := NewTestClient()
		_, err := client.Account().LinearCancelAllOrder(LinearCancelAllParam{})
		require.Error(t, err)
	})
}
