package bybit

import (
	"testing"

	"github.com/hirokisan/bybit/testhelper"
	"github.com/stretchr/testify/require"
)

func TestSpotSymbols(t *testing.T) {
	client := NewTestClient()
	res, err := client.Market().SpotSymbols()
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteDepth(t *testing.T) {
	client := NewTestClient()
	res, err := client.Market().SpotQuoteDepth(SpotQuoteDepthParam{
		Symbol: SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-depth.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteDepthMerged(t *testing.T) {
	client := NewTestClient()
	res, err := client.Market().SpotQuoteDepthMerged(SpotQuoteDepthMergedParam{
		Symbol: SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-depth-merged.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotTradesMerged(t *testing.T) {
	client := NewTestClient()
	res, err := client.Market().SpotQuoteTrades(SpotQuoteTradesParam{
		Symbol: SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-trades.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
