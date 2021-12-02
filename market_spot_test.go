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

func TestSpotQuoteTrades(t *testing.T) {
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

func TestSpotQuoteKline(t *testing.T) {
	client := NewTestClient()
	res, err := client.Market().SpotQuoteKline(SpotQuoteKlineParam{
		Symbol:   SymbolSpotBTCUSDT,
		Interval: SpotInterval1d,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTicker24hr(t *testing.T) {
	client := NewTestClient()
	symbol := SymbolSpotBTCUSDT
	res, err := client.Market().SpotQuoteTicker24hr(SpotQuoteTicker24hrParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-ticker-24hr.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTickerPrice(t *testing.T) {
	client := NewTestClient()
	symbol := SymbolSpotBTCUSDT
	res, err := client.Market().SpotQuoteTickerPrice(SpotQuoteTickerPriceParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-ticker-price.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTickerBookTicker(t *testing.T) {
	client := NewTestClient()
	symbol := SymbolSpotBTCUSDT
	res, err := client.Market().SpotQuoteTickerBookTicker(SpotQuoteTickerBookTickerParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-ticker-book-ticker.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotPostOrder(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()
	price := 28383.5
	res, err := client.Market().SpotPostOrder(SpotPostOrderParam{
		Symbol: SymbolSpotBTCUSDT,
		Qty:    0.01,
		Side:   SideBuy,
		Type:   OrderTypeSpotLimit,
		Price:  &price,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-post-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
	// TODO : cancel order
}

func TestSpotGetOrder(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()

	var orderID string
	{
		price := 28383.5
		res, err := client.Market().SpotPostOrder(SpotPostOrderParam{
			Symbol: SymbolSpotBTCUSDT,
			Qty:    0.01,
			Side:   SideBuy,
			Type:   OrderTypeSpotLimit,
			Price:  &price,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "", res.RetMsg)
		}
		orderID = res.Result.OrderID
	}

	res, err := client.Market().SpotGetOrder(SpotGetOrderParam{
		OrderID: &orderID,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-get-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
	// TODO : cancel order
}
