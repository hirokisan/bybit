//go:build integrationtestspotv1

package integrationtestspotv1

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestSpotSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1().SpotSymbols()
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteDepth(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1().SpotQuoteDepth(bybit.SpotQuoteDepthParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-depth.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteDepthMerged(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1().SpotQuoteDepthMerged(bybit.SpotQuoteDepthMergedParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
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
	client := bybit.NewTestClient()
	res, err := client.Spot().V1().SpotQuoteTrades(bybit.SpotQuoteTradesParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-trades.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Spot().V1().SpotQuoteKline(bybit.SpotQuoteKlineParam{
		Symbol:   bybit.SymbolSpotBTCUSDT,
		Interval: bybit.SpotInterval1d,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTicker24hr(t *testing.T) {
	client := bybit.NewTestClient()
	symbol := bybit.SymbolSpotBTCUSDT
	res, err := client.Spot().V1().SpotQuoteTicker24hr(bybit.SpotQuoteTicker24hrParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-ticker-24hr.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTickerPrice(t *testing.T) {
	client := bybit.NewTestClient()
	symbol := bybit.SymbolSpotBTCUSDT
	res, err := client.Spot().V1().SpotQuoteTickerPrice(bybit.SpotQuoteTickerPriceParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-ticker-price.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotQuoteTickerBookTicker(t *testing.T) {
	client := bybit.NewTestClient()
	symbol := bybit.SymbolSpotBTCUSDT
	res, err := client.Spot().V1().SpotQuoteTickerBookTicker(bybit.SpotQuoteTickerBookTickerParam{
		Symbol: &symbol,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-quote-v1-ticker-book-ticker.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotPostOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	price := 18383.5
	res, err := client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
		Qty:    0.01,
		Side:   bybit.SideBuy,
		Type:   bybit.OrderTypeSpotLimit,
		Price:  &price,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-v1-post-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
	// clean order
	orderID := res.Result.OrderID
	{
		_, err := client.Spot().V1().SpotDeleteOrder(bybit.SpotDeleteOrderParam{
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
		}
	}
}

func TestSpotGetOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	var orderID string
	{
		price := 28383.5
		res, err := client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
			Symbol: bybit.SymbolSpotBTCUSDT,
			Qty:    0.01,
			Side:   bybit.SideBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  &price,
		})
		{
			require.NoError(t, err)
		}
		orderID = res.Result.OrderID
	}

	res, err := client.Spot().V1().SpotGetOrder(bybit.SpotGetOrderParam{
		OrderID: &orderID,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-v1-get-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
	// clean order
	{
		_, err := client.Spot().V1().SpotDeleteOrder(bybit.SpotDeleteOrderParam{
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
		}
	}
}

func TestSpotDeleteOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	var orderID string
	{
		price := 28383.5
		res, err := client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
			Symbol: bybit.SymbolSpotBTCUSDT,
			Qty:    0.01,
			Side:   bybit.SideBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  &price,
		})
		{
			require.NoError(t, err)
		}
		orderID = res.Result.OrderID
	}

	res, err := client.Spot().V1().SpotDeleteOrder(bybit.SpotDeleteOrderParam{
		OrderID: &orderID,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-v1-delete-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotDeleteFastOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	var orderID string
	var symbol bybit.SymbolSpot
	{
		price := 28383.5
		res, err := client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
			Symbol: bybit.SymbolSpotBTCUSDT,
			Qty:    0.01,
			Side:   bybit.SideBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  &price,
		})
		{
			require.NoError(t, err)
		}
		orderID = res.Result.OrderID
		symbol = bybit.SymbolSpot(res.Result.Symbol)
	}

	res, err := client.Spot().V1().SpotDeleteOrderFast(bybit.SpotDeleteOrderFastParam{
		Symbol:  symbol,
		OrderID: &orderID,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-v1-delete-order-fast.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotOrderBatchCancel(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	var symbol bybit.SymbolSpot
	{
		price := 28383.5
		res, err := client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
			Symbol: bybit.SymbolSpotBTCUSDT,
			Qty:    0.01,
			Side:   bybit.SideBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  &price,
		})
		{
			require.NoError(t, err)
		}
		symbol = bybit.SymbolSpot(res.Result.Symbol)
	}

	res, err := client.Spot().V1().SpotOrderBatchCancel(bybit.SpotOrderBatchCancelParam{
		Symbol: symbol,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-v1-order-batch-cancel.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotOrderBatchFastCancel(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	var symbol bybit.SymbolSpot
	{
		price := 28383.5
		res, err := client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
			Symbol: bybit.SymbolSpotBTCUSDT,
			Qty:    0.01,
			Side:   bybit.SideBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  &price,
		})
		{
			require.NoError(t, err)
		}
		symbol = bybit.SymbolSpot(res.Result.Symbol)
	}

	res, err := client.Spot().V1().SpotOrderBatchFastCancel(bybit.SpotOrderBatchFastCancelParam{
		Symbol: symbol,
	})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-v1-order-batch-fast-cancel.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotOrderBatchCancelByIDs(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	var orderID string
	{
		price := 1.0
		res, err := client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
			Symbol: bybit.SymbolSpotBTCUSDT,
			Qty:    0.01,
			Side:   bybit.SideBuy,
			Type:   bybit.OrderTypeSpotLimit,
			Price:  &price,
		})
		{
			require.NoError(t, err)
		}
		orderID = res.Result.OrderID
	}

	res, err := client.Spot().V1().SpotOrderBatchCancelByIDs([]string{orderID})
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/spot-v1-order-batch-cancel-by-ids.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
