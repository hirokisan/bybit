package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5Order_CreateOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := "10000.0"
		param := V5CreateOrderParam{
			Category:  CategoryV5Spot,
			Symbol:    SymbolV5BTCUSDT,
			Side:      SideBuy,
			OrderType: OrderTypeLimit,
			Qty:       "0.01",
			Price:     &price,
		}

		path := "/v5/order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"orderId":     "1358868270414852352",
				"orderLinkId": "1676725721103693",
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

		resp, err := client.V5().Order().CreateOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := "10000.0"
		param := V5CreateOrderParam{
			Category:  CategoryV5Spot,
			Symbol:    SymbolV5BTCUSDT,
			Side:      SideBuy,
			OrderType: OrderTypeLimit,
			Qty:       "0.01",
			Price:     &price,
		}

		path := "/v5/order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"orderId":     "1358868270414852352",
				"orderLinkId": "1676725721103693",
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

		_, err = client.V5().Order().CreateOrder(param)
		assert.Error(t, err)
	})
}

func TestV5Order_CancelOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderId := "1358868270414852352"
		param := V5CancelOrderParam{
			Category: CategoryV5Spot,
			Symbol:   SymbolV5BTCUSDT,
			OrderID:  &orderId,
		}

		path := "/v5/order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"orderId":     orderId,
				"orderLinkId": "1676725721103693",
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

		resp, err := client.V5().Order().CancelOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		orderId := "1358868270414852352"
		param := V5CancelOrderParam{
			Category: CategoryV5Spot,
			Symbol:   SymbolV5BTCUSDT,
			OrderID:  &orderId,
		}

		path := "/v5/order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"orderId":     orderId,
				"orderLinkId": "1676725721103693",
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

		_, err = client.V5().Order().CancelOrder(param)
		assert.Error(t, err)
	})
}

func TestV5Order_GetOpenOrders(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		symbol := SymbolV5BTCUSDT
		param := V5GetOpenOrdersParam{
			Category: CategoryV5Linear,
			Symbol:   &symbol,
		}

		path := "/v5/order/realtime"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "",
				"list": []map[string]interface{}{
					{
						"symbol":             "ETHUSDT",
						"orderType":          "Limit",
						"orderLinkId":        "1672217748277652",
						"orderId":            "1321052653536515584",
						"cancelType":         "UNKNOWN",
						"avgPrice":           "",
						"stopOrderType":      "tpslOrder",
						"lastPriceOnCreated": "",
						"orderStatus":        "Cancelled",
						"takeProfit":         "",
						"cumExecValue":       "0",
						"triggerDirection":   0,
						"isLeverage":         "0",
						"rejectReason":       "",
						"price":              "1000",
						"orderIv":            "",
						"createdTime":        "1672217748287",
						"tpTriggerBy":        "",
						"positionIdx":        0,
						"timeInForce":        "GTC",
						"leavesValue":        "500",
						"updatedTime":        "1672217748287",
						"side":               "Buy",
						"triggerPrice":       "1500",
						"cumExecFee":         "0",
						"leavesQty":          "0",
						"slTriggerBy":        "",
						"closeOnTrigger":     false,
						"cumExecQty":         "0",
						"reduceOnly":         false,
						"qty":                "0.5",
						"stopLoss":           "",
						"triggerBy":          "1192.5",
					},
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

		resp, err := client.V5().Order().GetOpenOrders(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		symbol := SymbolV5BTCUSDT
		param := V5GetOpenOrdersParam{
			Category: CategoryV5Linear,
			Symbol:   &symbol,
		}

		path := "/v5/order/realtime"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "",
				"list": []map[string]interface{}{
					{
						"symbol":             "ETHUSDT",
						"orderType":          "Limit",
						"orderLinkId":        "1672217748277652",
						"orderId":            "1321052653536515584",
						"cancelType":         "UNKNOWN",
						"avgPrice":           "",
						"stopOrderType":      "tpslOrder",
						"lastPriceOnCreated": "",
						"orderStatus":        "Cancelled",
						"takeProfit":         "",
						"cumExecValue":       "0",
						"triggerDirection":   0,
						"isLeverage":         "0",
						"rejectReason":       "",
						"price":              "1000",
						"orderIv":            "",
						"createdTime":        "1672217748287",
						"tpTriggerBy":        "",
						"positionIdx":        0,
						"timeInForce":        "GTC",
						"leavesValue":        "500",
						"updatedTime":        "1672217748287",
						"side":               "Buy",
						"triggerPrice":       "1500",
						"cumExecFee":         "0",
						"leavesQty":          "0",
						"slTriggerBy":        "",
						"closeOnTrigger":     false,
						"cumExecQty":         "0",
						"reduceOnly":         false,
						"qty":                "0.5",
						"stopLoss":           "",
						"triggerBy":          "1192.5",
					},
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

		_, err = client.V5().Order().GetOpenOrders(param)
		assert.Error(t, err)
	})
}
