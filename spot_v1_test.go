package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/hirokisan/bybit/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpotPostOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := SpotPostOrderParam{
			Symbol: SymbolSpotBTCUSDT,
			Qty:    1.1,
			Side:   SideBuy,
			Type:   OrderTypeSpotMarket,
		}

		path := "/spot/v1/order"
		method := http.MethodPost
		status := http.StatusOK
		respBody := SpotPostOrderResponse{
			Result: SpotPostOrderResult{
				OrderID:      "1037799004578056704",
				OrderLinkID:  "1638451282020267",
				Symbol:       string(param.Symbol),
				TransactTime: "1638451282090",
				Price:        "28383.5",
				OrigQty:      fmt.Sprintf("%f", param.Qty),
				Type:         OrderTypeSpotMarket,
				Side:         string(param.Side),
				Status:       OrderStatusSpotNew,
				TimeInForce:  TimeInForceSpotGTC,
				AccountID:    "213998",
				SymbolName:   "BTCUSDT",
				ExecutedQty:  "0",
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL).
			WithAuth("test", "test")

		resp, err := client.Spot().V1().SpotPostOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := SpotPostOrderParam{
			Symbol: SymbolSpotBTCUSDT,
			Qty:    1.1,
			Side:   SideBuy,
			Type:   OrderTypeSpotMarket,
		}

		path := "/spot/v1/order"
		method := http.MethodPost
		status := http.StatusOK
		respBody := SpotPostOrderResponse{
			Result: SpotPostOrderResult{
				OrderID:      "1037799004578056704",
				OrderLinkID:  "1638451282020267",
				Symbol:       string(param.Symbol),
				TransactTime: "1638451282090",
				Price:        "28383.5",
				OrigQty:      fmt.Sprintf("%f", param.Qty),
				Type:         OrderTypeSpotMarket,
				Side:         string(param.Side),
				Status:       OrderStatusSpotNew,
				TimeInForce:  TimeInForceSpotGTC,
				AccountID:    "213998",
				SymbolName:   "BTCUSDT",
				ExecutedQty:  "0",
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		_, err = client.Spot().V1().SpotPostOrder(param)
		assert.Error(t, err)
	})
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

func TestSpotOpenOrders(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		symbol := SymbolSpotBTCUSDT
		param := SpotOpenOrdersParam{
			Symbol: &symbol,
		}

		path := "/spot/v1/open-orders"
		method := http.MethodGet
		status := http.StatusOK
		respBody := SpotOpenOrdersResponse{
			Result: []SpotOpenOrdersResult{
				{
					AccountID:           "213998",
					ExchangeID:          "301",
					Symbol:              "BTCUSDT",
					SymbolName:          "BTCUSDT",
					OrderLinkID:         "1664962738850574",
					OrderID:             "1260193223383517952",
					Price:               "1",
					OrigQty:             "0.01",
					ExecutedQty:         "0",
					CummulativeQuoteQty: "0",
					AvgPrice:            "0",
					Status:              "NEW",
					TimeInForce:         "GTC",
					Type:                "LIMIT",
					Side:                "BUY",
					StopPrice:           "0.0",
					IcebergQty:          "0.0",
					Time:                "1664962738856",
					UpdateTime:          "1664962738874",
					IsWorking:           true,
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL).
			WithAuth("test", "test")

		resp, err := client.Spot().V1().SpotOpenOrders(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		symbol := SymbolSpotBTCUSDT
		param := SpotOpenOrdersParam{
			Symbol: &symbol,
		}

		path := "/spot/v1/open-orders"
		method := http.MethodGet
		status := http.StatusOK
		respBody := SpotOpenOrdersResponse{
			Result: []SpotOpenOrdersResult{
				{
					AccountID:           "213998",
					ExchangeID:          "301",
					Symbol:              "BTCUSDT",
					SymbolName:          "BTCUSDT",
					OrderLinkID:         "1664962738850574",
					OrderID:             "1260193223383517952",
					Price:               "1",
					OrigQty:             "0.01",
					ExecutedQty:         "0",
					CummulativeQuoteQty: "0",
					AvgPrice:            "0",
					Status:              "NEW",
					TimeInForce:         "GTC",
					Type:                "LIMIT",
					Side:                "BUY",
					StopPrice:           "0.0",
					IcebergQty:          "0.0",
					Time:                "1664962738856",
					UpdateTime:          "1664962738874",
					IsWorking:           true,
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		_, err = client.Spot().V1().SpotOpenOrders(param)
		assert.Error(t, err)
	})
}
