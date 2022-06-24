package bybit

import (
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/hirokisan/bybit/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpotPostOrder(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()
	price := 28383.5
	res, err := client.Account().SpotPostOrder(SpotPostOrderParam{
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
	// clean order
	orderID := res.Result.OrderID
	{
		res, err := client.Account().SpotDeleteOrder(SpotDeleteOrderParam{
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "", res.RetMsg)
		}
	}
}

func TestSpotGetOrder(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()

	var orderID string
	{
		price := 28383.5
		res, err := client.Account().SpotPostOrder(SpotPostOrderParam{
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

	res, err := client.Account().SpotGetOrder(SpotGetOrderParam{
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
	// clean order
	{
		res, err := client.Account().SpotDeleteOrder(SpotDeleteOrderParam{
			OrderID: &orderID,
		})
		{
			require.NoError(t, err)
			require.Equal(t, "", res.RetMsg)
		}
	}
}

func TestSpotDeleteOrder(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()

	var orderID string
	{
		price := 28383.5
		res, err := client.Account().SpotPostOrder(SpotPostOrderParam{
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

	res, err := client.Account().SpotDeleteOrder(SpotDeleteOrderParam{
		OrderID: &orderID,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-delete-order.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotDeleteFastOrder(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()

	var orderID string
	var symbol SymbolSpot
	{
		price := 28383.5
		res, err := client.Account().SpotPostOrder(SpotPostOrderParam{
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
		symbol = SymbolSpot(res.Result.Symbol)
	}

	res, err := client.Account().SpotDeleteOrderFast(SpotDeleteOrderFastParam{
		Symbol:  symbol,
		OrderID: &orderID,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-delete-order-fast.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotOrderBatchCancel(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()

	var symbol SymbolSpot
	{
		price := 28383.5
		res, err := client.Account().SpotPostOrder(SpotPostOrderParam{
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
		symbol = SymbolSpot(res.Result.Symbol)
	}

	res, err := client.Account().SpotOrderBatchCancel(SpotOrderBatchCancelParam{
		Symbol: symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-order-batch-cancel.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotOrderBatchFastCancel(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()

	var symbol SymbolSpot
	{
		price := 28383.5
		res, err := client.Account().SpotPostOrder(SpotPostOrderParam{
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
		symbol = SymbolSpot(res.Result.Symbol)
	}

	res, err := client.Account().SpotOrderBatchFastCancel(SpotOrderBatchFastCancelParam{
		Symbol: symbol,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-order-batch-fast-cancel.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotOrderBatchCancelByIDs(t *testing.T) {
	client := NewTestClient().WithAuthFromEnv()

	var orderID string
	{
		price := 28383.5
		res, err := client.Account().SpotPostOrder(SpotPostOrderParam{
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

	res, err := client.Account().SpotOrderBatchCancelByIDs([]string{orderID})
	{
		require.NoError(t, err)
		require.Equal(t, "", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/spot-v1-order-batch-cancel-by-ids.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSpotOrderBatchCancelParam(t *testing.T) {
	param := SpotOrderBatchCancelParam{
		Symbol: SymbolSpotBTCUSDT,
		Types:  []OrderTypeSpot{OrderTypeSpotLimit, OrderTypeSpotMarket},
	}
	queryString, err := query.Values(param)
	require.NoError(t, err)
	want := url.Values{}
	want.Add("symbolId", string(param.Symbol))
	var types []string
	for _, t := range param.Types {
		types = append(types, string(t))
	}
	want.Add("orderTypes", strings.Join(types, ","))

	assert.Equal(t, want, queryString)
}

func TestSpotOrderBatchFastCancelParam(t *testing.T) {
	param := SpotOrderBatchFastCancelParam{
		Symbol: SymbolSpotBTCUSDT,
		Types:  []OrderTypeSpot{OrderTypeSpotLimit, OrderTypeSpotMarket},
	}
	queryString, err := query.Values(param)
	require.NoError(t, err)
	want := url.Values{}
	want.Add("symbolId", string(param.Symbol))
	var types []string
	for _, t := range param.Types {
		types = append(types, string(t))
	}
	want.Add("orderTypes", strings.Join(types, ","))

	assert.Equal(t, want, queryString)
}
