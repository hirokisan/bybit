//go:build integrationtestv5market

package integrationtestv5market

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetKline(bybit.V5GetKlineParam{
		Category: bybit.CategoryV5Spot,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Interval: bybit.IntervalD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetMarkPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetMarkPriceKline(bybit.V5GetMarkPriceKlineParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Interval: bybit.IntervalD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-mark-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetIndexPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetIndexPriceKline(bybit.V5GetIndexPriceKlineParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Interval: bybit.IntervalD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
