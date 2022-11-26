//go:build integrationtestderivativeunifiedmargin

package integrationtestderivativeunifiedmargin

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestDerivativesOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesOrderBook(bybit.DerivativesOrderBookParam{
		Symbol:   bybit.SymbolDerivativeBTCUSDT,
		Category: bybit.CategoryDerivativeLinear,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestDerivativesKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesKline(bybit.DerivativesKlineParam{
		Symbol:   bybit.SymbolDerivativeBTCUSDT,
		Category: bybit.CategoryDerivativeLinear,
		Interval: bybit.IntervalD,
		Start:    1652112000000,
		End:      1652544000000,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestDerivativesTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesTickers(bybit.DerivativesTickersParam{
		Category: bybit.CategoryDerivativeLinear,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestDerivativesTickersForOption(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesTickersForOption(bybit.DerivativesTickersForOptionParam{
		Symbol: bybit.SymbolDerivativeBTC31MAR23_40000C,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-tickers-for-option.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestDerivativesInstruments(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesInstruments(bybit.DerivativesInstrumentsParam{
		Category: bybit.CategoryDerivativeLinear,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-instruments-info.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestDerivativesInstrumentsForOption(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesInstrumentsForOption(bybit.DerivativesInstrumentsForOptionParam{})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-instruments-info-for-option.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestDerivativesMarkPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().Contract().DerivativesMarkPriceKline(bybit.DerivativesMarkPriceKlineParam{
		Symbol:   bybit.SymbolDerivativeBTCUSDT,
		Category: bybit.CategoryDerivativeLinear,
		Interval: bybit.IntervalD,
		Start:    1652112000000,
		End:      1652544000000,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-mark-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestDerivativesIndexPriveKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesIndexPriceKline(bybit.DerivativesIndexPriceKlineParam{
		Symbol:   bybit.SymbolDerivativeBTCUSDT,
		Category: bybit.CategoryDerivativeLinear,
		Interval: bybit.IntervalD,
		Start:    1652112000000,
		End:      1652544000000,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
