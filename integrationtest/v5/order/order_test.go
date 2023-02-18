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
