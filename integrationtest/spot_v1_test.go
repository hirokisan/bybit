//go:build integrationtest

package integrationtest

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestSpotSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1.SpotSymbols()
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-v1-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteDepth(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1.SpotQuoteDepth(bybit.SpotQuoteDepthParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-quote-v1-depth.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteDepthMerged(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1.SpotQuoteDepthMerged(bybit.SpotQuoteDepthMergedParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-quote-v1-depth-merged.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTrades(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1.SpotQuoteTrades(bybit.SpotQuoteTradesParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-quote-v1-trades.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1.SpotQuoteKline(bybit.SpotQuoteKlineParam{
		Symbol:   bybit.SymbolSpotBTCUSDT,
		Interval: bybit.SpotInterval1d,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-quote-v1-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTicker24hr(t *testing.T) {
	client := bybit.NewTestClient()
	symbol := bybit.SymbolSpotBTCUSDT
	res, err := client.Spot().V1.SpotQuoteTicker24hr(bybit.SpotQuoteTicker24hrParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-quote-v1-ticker-24hr.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTickerPrice(t *testing.T) {
	client := bybit.NewTestClient()
	symbol := bybit.SymbolSpotBTCUSDT
	res, err := client.Spot().V1.SpotQuoteTickerPrice(bybit.SpotQuoteTickerPriceParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-quote-v1-ticker-price.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTickerBookTicker(t *testing.T) {
	client := bybit.NewTestClient()
	symbol := bybit.SymbolSpotBTCUSDT
	res, err := client.Spot().V1.SpotQuoteTickerBookTicker(bybit.SpotQuoteTickerBookTickerParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1/spot-quote-v1-ticker-book-ticker.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
