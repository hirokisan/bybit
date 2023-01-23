package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListLinearKline(t *testing.T) {
	param := ListLinearKlineParam{
		Symbol: SymbolFutureBTCUSDT,
	}

	path := "/public/linear/kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := ListLinearKlineResponse{
		Result: []ListLinearKlineResult{
			{
				Symbol:   "BTCUSDT",
				Period:   "120",
				Interval: "120",
				StartAt:  1665489600,
				OpenTime: 1665489600,
				Volume:   3371.005,
				Open:     19095,
				High:     19408,
				Low:      19056.5,
				Close:    19057,
				Turnover: 64668178.383,
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

	resp, err := client.Future().USDTPerpetual().ListLinearKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody, *resp)
}

func TestListLinearOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := ListLinearOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListLinearOrderResponse{
			Result: ListLinearOrderResult{
				CurrentPage: 1,
				Content: []ListLinearOrderResultContent{
					{
						OrderID:        "08f55183-76ce-4269-9039-433ce6881632",
						UserID:         146940,
						Symbol:         SymbolFutureBTCUSDT,
						Side:           SideBuy,
						OrderType:      OrderTypeLimit,
						Price:          10000.0,
						Qty:            0.001,
						TimeInForce:    TimeInForceGoodTillCancel,
						OrderStatus:    OrderStatusNew,
						LastExecPrice:  0,
						CumExecQty:     0,
						CumExecValue:   0,
						CumExecFee:     0,
						ReduceOnly:     false,
						CloseOnTrigger: false,
						OrderLinkID:    "",
						CreatedTime:    "2022-10-07T08:48:50Z",
						UpdatedTime:    "2022-10-07T08:48:50Z",
						TakeProfit:     0,
						StopLoss:       0,
						TpTriggerBy:    "UNKNOWN",
						SlTriggerBy:    "UNKNOWN",
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

		resp, err := client.Future().USDTPerpetual().ListLinearOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})

	t.Run("authentication required", func(t *testing.T) {
		param := ListLinearOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListLinearOrderResponse{
			Result: ListLinearOrderResult{
				CurrentPage: 1,
				Content: []ListLinearOrderResultContent{
					{
						OrderID:        "08f55183-76ce-4269-9039-433ce6881632",
						UserID:         146940,
						Symbol:         SymbolFutureBTCUSDT,
						Side:           SideBuy,
						OrderType:      OrderTypeLimit,
						Price:          10000.0,
						Qty:            0.001,
						TimeInForce:    TimeInForceGoodTillCancel,
						OrderStatus:    OrderStatusNew,
						LastExecPrice:  0,
						CumExecQty:     0,
						CumExecValue:   0,
						CumExecFee:     0,
						ReduceOnly:     false,
						CloseOnTrigger: false,
						OrderLinkID:    "",
						CreatedTime:    "2022-10-07T08:48:50Z",
						UpdatedTime:    "2022-10-07T08:48:50Z",
						TakeProfit:     0,
						StopLoss:       0,
						TpTriggerBy:    "UNKNOWN",
						SlTriggerBy:    "UNKNOWN",
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

		_, err = client.Future().USDTPerpetual().ListLinearOrder(param)
		assert.Error(t, err)
	})
}

func TestQueryLinearOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := QueryLinearOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/order/search"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryLinearOrderResponse{
			Result: []QueryLinearOrderResult{
				{
					OrderID:        "d8bc8319-17b1-41a5-bff6-728458f35248",
					UserID:         146940,
					Symbol:         "BTCUSDT",
					Side:           "Buy",
					OrderType:      "Limit",
					Price:          10000,
					Qty:            0.001,
					TimeInForce:    "GoodTillCancel",
					OrderStatus:    "New",
					LastExecPrice:  0,
					CumExecQty:     0,
					CumExecValue:   0,
					CumExecFee:     0,
					ReduceOnly:     false,
					CloseOnTrigger: false,
					OrderLinkID:    "",
					CreatedTime:    "2022-10-13T02:55:01Z",
					UpdatedTime:    "2022-10-13T02:55:01Z",
					TakeProfit:     0,
					StopLoss:       0,
					TpTriggerBy:    "UNKNOWN",
					SlTriggerBy:    "UNKNOWN",
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

		resp, err := client.Future().USDTPerpetual().QueryLinearOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := QueryLinearOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/order/search"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryLinearOrderResponse{
			Result: []QueryLinearOrderResult{
				{
					OrderID:        "d8bc8319-17b1-41a5-bff6-728458f35248",
					UserID:         146940,
					Symbol:         "BTCUSDT",
					Side:           "Buy",
					OrderType:      "Limit",
					Price:          10000,
					Qty:            0.001,
					TimeInForce:    "GoodTillCancel",
					OrderStatus:    "New",
					LastExecPrice:  0,
					CumExecQty:     0,
					CumExecValue:   0,
					CumExecFee:     0,
					ReduceOnly:     false,
					CloseOnTrigger: false,
					OrderLinkID:    "",
					CreatedTime:    "2022-10-13T02:55:01Z",
					UpdatedTime:    "2022-10-13T02:55:01Z",
					TakeProfit:     0,
					StopLoss:       0,
					TpTriggerBy:    "UNKNOWN",
					SlTriggerBy:    "UNKNOWN",
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

		_, err = client.Future().USDTPerpetual().QueryLinearOrder(param)
		assert.Error(t, err)
	})
}

func TestCreateLinearStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := 19400.5
		param := CreateLinearStopOrderParam{
			Side:           SideBuy,
			Symbol:         SymbolFutureBTCUSDT,
			OrderType:      OrderTypeMarket,
			Qty:            0.001,
			BasePrice:      price,
			StopPx:         price + 200,
			TimeInForce:    TimeInForceGoodTillCancel,
			TriggerBy:      TriggerByFutureLastPrice,
			ReduceOnly:     true,
			CloseOnTrigger: true,
		}

		path := "/private/linear/stop-order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateLinearStopOrderResponse{
			Result: CreateLinearStopOrderResult{
				StopOrderID:    "524c6849-2670-4d3b-8e03-38d728d0c8d4",
				UserID:         146940,
				Symbol:         "BTCUSDT",
				Side:           "Buy",
				OrderType:      "Market",
				Price:          0,
				Qty:            0.001,
				TimeInForce:    "ImmediateOrCancel",
				OrderStatus:    "Untriggered",
				TriggerPrice:   19600.5,
				OrderLinkID:    "",
				CreatedTime:    "2022-10-14T00:12:04Z",
				UpdatedTime:    "2022-10-14T00:12:04Z",
				BasePrice:      "19400.50",
				TriggerBy:      "LastPrice",
				TpTriggerBy:    "UNKNOWN",
				SlTriggerBy:    "UNKNOWN",
				TakeProfit:     0,
				StopLoss:       0,
				ReduceOnly:     true,
				CloseOnTrigger: true,
				PositionIdx:    2,
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

		resp, err := client.Future().USDTPerpetual().CreateLinearStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := 19400.5
		param := CreateLinearStopOrderParam{
			Side:           SideBuy,
			Symbol:         SymbolFutureBTCUSDT,
			OrderType:      OrderTypeMarket,
			Qty:            0.001,
			BasePrice:      price,
			StopPx:         price + 200,
			TimeInForce:    TimeInForceGoodTillCancel,
			TriggerBy:      TriggerByFutureLastPrice,
			ReduceOnly:     true,
			CloseOnTrigger: true,
		}

		path := "/private/linear/stop-order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateLinearStopOrderResponse{
			Result: CreateLinearStopOrderResult{
				StopOrderID:    "524c6849-2670-4d3b-8e03-38d728d0c8d4",
				UserID:         146940,
				Symbol:         "BTCUSDT",
				Side:           "Buy",
				OrderType:      "Market",
				Price:          0,
				Qty:            0.001,
				TimeInForce:    "ImmediateOrCancel",
				OrderStatus:    "Untriggered",
				TriggerPrice:   19600.5,
				OrderLinkID:    "",
				CreatedTime:    "2022-10-14T00:12:04Z",
				UpdatedTime:    "2022-10-14T00:12:04Z",
				BasePrice:      "19400.50",
				TriggerBy:      "LastPrice",
				TpTriggerBy:    "UNKNOWN",
				SlTriggerBy:    "UNKNOWN",
				TakeProfit:     0,
				StopLoss:       0,
				ReduceOnly:     true,
				CloseOnTrigger: true,
				PositionIdx:    2,
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

		_, err = client.Future().USDTPerpetual().CreateLinearStopOrder(param)
		assert.Error(t, err)
	})
}

func TestListLinearStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := ListLinearStopOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/stop-order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListLinearStopOrderResponse{
			Result: ListLinearStopOrderResult{
				CurrentPage: 1,
				LastPage:    0,
				Content: []ListLinearStopOrderResultContent{
					{
						StopOrderID:    "ce6e676e-21b3-4d6c-bfef-25f6996fc2dd",
						UserID:         146940,
						Symbol:         "BTCUSDT",
						Side:           "Buy",
						OrderType:      "Market",
						Price:          0,
						Qty:            0.001,
						TimeInForce:    "ImmediateOrCancel",
						OrderStatus:    "Untriggered",
						TriggerPrice:   20000.5,
						OrderLinkID:    "",
						CreatedTime:    "2022-10-14T04:38:00Z",
						UpdatedTime:    "2022-10-14T04:38:00Z",
						TakeProfit:     0,
						StopLoss:       0,
						TriggerBy:      "LastPrice",
						BasePrice:      "19800.50",
						TpTriggerBy:    "UNKNOWN",
						SlTriggerBy:    "UNKNOWN",
						ReduceOnly:     true,
						CloseOnTrigger: true,
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

		resp, err := client.Future().USDTPerpetual().ListLinearStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})

	t.Run("authentication required", func(t *testing.T) {
		param := ListLinearStopOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/stop-order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListLinearStopOrderResponse{
			Result: ListLinearStopOrderResult{
				CurrentPage: 1,
				LastPage:    0,
				Content: []ListLinearStopOrderResultContent{
					{
						StopOrderID:    "ce6e676e-21b3-4d6c-bfef-25f6996fc2dd",
						UserID:         146940,
						Symbol:         "BTCUSDT",
						Side:           "Buy",
						OrderType:      "Market",
						Price:          0,
						Qty:            0.001,
						TimeInForce:    "ImmediateOrCancel",
						OrderStatus:    "Untriggered",
						TriggerPrice:   20000.5,
						OrderLinkID:    "",
						CreatedTime:    "2022-10-14T04:38:00Z",
						UpdatedTime:    "2022-10-14T04:38:00Z",
						TakeProfit:     0,
						StopLoss:       0,
						TriggerBy:      "LastPrice",
						BasePrice:      "19800.50",
						TpTriggerBy:    "UNKNOWN",
						SlTriggerBy:    "UNKNOWN",
						ReduceOnly:     true,
						CloseOnTrigger: true,
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

		_, err = client.Future().USDTPerpetual().ListLinearStopOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelLinearStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stopOrderID := "1cc84e40-8b72-4191-935a-cc3feafbf02e"
		param := CancelLinearStopOrderParam{
			Symbol:      SymbolFutureBTCUSD,
			StopOrderID: &stopOrderID,
		}

		path := "/private/linear/stop-order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelLinearStopOrderResponse{
			Result: CancelLinearStopOrderResult{
				StopOrderID: stopOrderID,
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

		resp, err := client.Future().USDTPerpetual().CancelLinearStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		stopOrderID := "1cc84e40-8b72-4191-935a-cc3feafbf02e"
		param := CancelLinearStopOrderParam{
			Symbol:      SymbolFutureBTCUSD,
			StopOrderID: &stopOrderID,
		}

		path := "/private/linear/stop-order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelLinearStopOrderResponse{
			Result: CancelLinearStopOrderResult{
				StopOrderID: stopOrderID,
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

		_, err = client.Future().USDTPerpetual().CancelLinearStopOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelAllLinearStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := CancelAllLinearStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/private/linear/stop-order/cancel-all"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllLinearStopOrderResponse{
			Result: CancelAllLinearStopOrderResult{
				"ab39efc4-e15f-4538-92d5-2e02a14f4845",
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

		resp, err := client.Future().USDTPerpetual().CancelAllLinearStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := CancelAllLinearStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/private/linear/stop-order/cancel-all"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllLinearStopOrderResponse{
			Result: CancelAllLinearStopOrderResult{
				"ab39efc4-e15f-4538-92d5-2e02a14f4845",
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

		_, err = client.Future().USDTPerpetual().CancelAllLinearStopOrder(param)
		assert.Error(t, err)
	})
}

func TestReplaceLinearOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderID := "test"
		newPrice := 10000.0
		param := ReplaceLinearOrderParam{
			Symbol:   SymbolFutureBTCUSDT,
			OrderID:  &orderID,
			NewPrice: &newPrice,
		}

		path := "/private/linear/order/replace"
		method := http.MethodPost
		status := http.StatusOK
		respBody := ReplaceLinearOrderResponse{
			Result: ReplaceLinearOrderResult{
				OrderID: "b777573b-026e-48ee-b9ac-28c81e829ae8",
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

		resp, err := client.Future().USDTPerpetual().ReplaceLinearOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		orderID := "test"
		newPrice := 10000.0
		param := ReplaceLinearOrderParam{
			Symbol:   SymbolFutureBTCUSDT,
			OrderID:  &orderID,
			NewPrice: &newPrice,
		}

		path := "/private/linear/order/replace"
		method := http.MethodPost
		status := http.StatusOK
		respBody := ReplaceLinearOrderResponse{
			Result: ReplaceLinearOrderResult{
				OrderID: "b777573b-026e-48ee-b9ac-28c81e829ae8",
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

		_, err = client.Future().USDTPerpetual().ReplaceLinearOrder(param)
		assert.Error(t, err)
	})
}

func TestQueryLinearStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := QueryLinearStopOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/stop-order/search"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryLinearStopOrderResponse{
			Result: []QueryLinearStopOrderResult{
				{
					StopOrderID:    "98674e2e-779b-4c5d-8a9f-bbcd45ccbcf4",
					UserID:         146940,
					Symbol:         "BTCUSDT",
					Side:           "Buy",
					OrderType:      "Market",
					Price:          0,
					Qty:            0.001,
					TimeInForce:    "ImmediateOrCancel",
					OrderStatus:    "Untriggered",
					TriggerPrice:   20000.5,
					BasePrice:      "19800.50",
					OrderLinkID:    "",
					CreatedTime:    "2022-10-15T04:02:36.000Z",
					UpdatedTime:    "2022-10-15T04:02:36.000Z",
					TakeProfit:     0,
					StopLoss:       0,
					TpTriggerBy:    "UNKNOWN",
					SlTriggerBy:    "UNKNOWN",
					TriggerBy:      "LastPrice",
					ReduceOnly:     true,
					CloseOnTrigger: true,
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

		resp, err := client.Future().USDTPerpetual().QueryLinearStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := QueryLinearStopOrderParam{
			Symbol: SymbolFutureBTCUSDT,
		}

		path := "/private/linear/stop-order/search"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryLinearStopOrderResponse{
			Result: []QueryLinearStopOrderResult{
				{
					StopOrderID:    "98674e2e-779b-4c5d-8a9f-bbcd45ccbcf4",
					UserID:         146940,
					Symbol:         "BTCUSDT",
					Side:           "Buy",
					OrderType:      "Market",
					Price:          0,
					Qty:            0.001,
					TimeInForce:    "ImmediateOrCancel",
					OrderStatus:    "Untriggered",
					TriggerPrice:   20000.5,
					BasePrice:      "19800.50",
					OrderLinkID:    "",
					CreatedTime:    "2022-10-15T04:02:36.000Z",
					UpdatedTime:    "2022-10-15T04:02:36.000Z",
					TakeProfit:     0,
					StopLoss:       0,
					TpTriggerBy:    "UNKNOWN",
					SlTriggerBy:    "UNKNOWN",
					TriggerBy:      "LastPrice",
					ReduceOnly:     true,
					CloseOnTrigger: true,
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

		_, err = client.Future().USDTPerpetual().QueryLinearStopOrder(param)
		assert.Error(t, err)
	})
}

func TestListLinearPositions(t *testing.T) {
	t.Run("Permission denied", func(t *testing.T) {
		path := "/private/linear/position/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"ret_code": 10005,
			"ret_msg":  "Permission denied, please check your API key permissions",
			"result":   struct{}{},
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

		_, err = client.Future().USDTPerpetual().ListLinearPositions()
		require.Error(t, err)

		var wantErr *ErrorResponse
		assert.ErrorAs(t, err, &wantErr)
	})
}

func TestLinearTradingStop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := 20000.0
		param := LinearTradingStopParam{
			Symbol:     SymbolFutureBTCUSDT,
			Side:       SideBuy,
			TakeProfit: &price,
		}

		path := "/private/linear/position/trading-stop"
		method := http.MethodPost
		status := http.StatusOK
		respBody := LinearTradingStopResponse{}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL).
			WithAuth("test", "test")

		resp, err := client.Future().USDTPerpetual().LinearTradingStop(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := 20000.0
		param := LinearTradingStopParam{
			Symbol:     SymbolFutureBTCUSDT,
			Side:       SideBuy,
			TakeProfit: &price,
		}

		path := "/private/linear/position/trading-stop"
		method := http.MethodPost
		status := http.StatusOK
		respBody := LinearTradingStopResponse{}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		_, err = client.Future().USDTPerpetual().LinearTradingStop(param)
		assert.Error(t, err)
	})
}
