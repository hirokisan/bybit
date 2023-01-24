package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := CreateOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolFutureBTCUSD,
			OrderType:   OrderTypeLimit,
			Qty:         10,
			TimeInForce: TimeInForceGoodTillCancel,
		}

		path := "/v2/private/order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateOrderResponse{
			Result: CreateOrderResult{
				CreateOrder: CreateOrder{
					UserID:        146940,
					OrderID:       "0fd3194e-14b3-4050-9e80-941ec5d169c5",
					Symbol:        param.Symbol,
					OrderType:     param.OrderType,
					Price:         0.0,
					Qty:           float64(param.Qty),
					TimeInForce:   param.TimeInForce,
					OrderStatus:   OrderStatusCreated,
					LastExecTime:  0,
					LastExecPrice: 0,
					LeavesQty:     0,
					CumExecQty:    0,
					CumExecValue:  0,
					CumExecFee:    0,
					RejectReason:  "EC_NoError",
					OrderLinkID:   "0",
					CreatedAt:     "2021-11-26T15:08:07Z",
					UpdatedAt:     "2021-11-26T15:08:07Z",
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

		resp, err := client.Future().InversePerpetual().CreateOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := CreateOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolFutureBTCUSD,
			OrderType:   OrderTypeLimit,
			Qty:         10,
			TimeInForce: TimeInForceGoodTillCancel,
		}

		path := "/v2/private/order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateOrderResponse{
			Result: CreateOrderResult{
				CreateOrder: CreateOrder{
					UserID:        146940,
					OrderID:       "0fd3194e-14b3-4050-9e80-941ec5d169c5",
					Symbol:        param.Symbol,
					OrderType:     param.OrderType,
					Price:         0.0,
					Qty:           float64(param.Qty),
					TimeInForce:   param.TimeInForce,
					OrderStatus:   OrderStatusCreated,
					LastExecTime:  0,
					LastExecPrice: 0,
					LeavesQty:     0,
					CumExecQty:    0,
					CumExecValue:  0,
					CumExecFee:    0,
					RejectReason:  "EC_NoError",
					OrderLinkID:   "0",
					CreatedAt:     "2021-11-26T15:08:07Z",
					UpdatedAt:     "2021-11-26T15:08:07Z",
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

		_, err = client.Future().InversePerpetual().CreateOrder(param)
		assert.Error(t, err)
	})
}

func TestListOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := ListOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListOrderResponse{
			Result: ListOrderResult{
				ListOrders: []ListOrder{
					{
						UserID:       146940,
						Symbol:       param.Symbol,
						Side:         SideBuy,
						OrderType:    OrderTypeLimit,
						Price:        "10000",
						Qty:          "1",
						TimeInForce:  TimeInForceGoodTillCancel,
						OrderStatus:  OrderStatusCreated,
						LeavesQty:    "0",
						LeavesValue:  "0",
						CumExecQty:   "0",
						CumExecValue: "0",
						CumExecFee:   "0",
						RejectReason: "EC_PerCancelRequest",
						OrderLinkID:  "",
						CreatedAt:    "2022-06-20T13:33:36.105Z",
						OrderID:      "04e633e6-92a9-4718-a83e-de92a72ce20a",
						TakeProfit:   "0.0000",
						StopLoss:     "0.0000",
						TpTriggerBy:  "UNKNOWN",
						SlTriggerBy:  "UNKNOWN",
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

		resp, err := client.Future().InversePerpetual().ListOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := ListOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListOrderResponse{
			Result: ListOrderResult{
				ListOrders: []ListOrder{
					{
						UserID:       146940,
						Symbol:       param.Symbol,
						Side:         SideBuy,
						OrderType:    OrderTypeLimit,
						Price:        "10000",
						Qty:          "1",
						TimeInForce:  TimeInForceGoodTillCancel,
						OrderStatus:  OrderStatusCreated,
						LeavesQty:    "0",
						LeavesValue:  "0",
						CumExecQty:   "0",
						CumExecValue: "0",
						CumExecFee:   "0",
						RejectReason: "EC_PerCancelRequest",
						OrderLinkID:  "",
						CreatedAt:    "2022-06-20T13:33:36.105Z",
						OrderID:      "04e633e6-92a9-4718-a83e-de92a72ce20a",
						TakeProfit:   "0.0000",
						StopLoss:     "0.0000",
						TpTriggerBy:  "UNKNOWN",
						SlTriggerBy:  "UNKNOWN",
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

		_, err = client.Future().InversePerpetual().ListOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelAllOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := CancelAllOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllOrderResponse{
			Result: []CancelAllOrderResult{
				{
					ClOrdID:     "d1a86997-d6c7-45b5-bd58-1b3caf17387f",
					OrderLinkID: "",
					UserID:      146940,
					Symbol:      "BTCUSD",
					Side:        "Buy",
					OrderType:   "Limit",
					Price:       "10000",
					Qty:         1,
					TimeInForce: "GoodTillCancel",
					CreateType:  "CreateByUser",
					CancelType:  "CancelByUser",
					OrderStatus: "",
					LeavesQty:   1,
					LeavesValue: "0",
					CreatedAt:   "2022-10-09T12:04:22.633327064Z",
					UpdatedAt:   "2022-10-09T12:04:22.72531757Z",
					CrossStatus: "PendingCancel",
					CrossSeq:    5542233074,
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

		resp, err := client.Future().InversePerpetual().CancelAllOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := CancelAllOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllOrderResponse{
			Result: []CancelAllOrderResult{
				{
					ClOrdID:     "d1a86997-d6c7-45b5-bd58-1b3caf17387f",
					OrderLinkID: "",
					UserID:      146940,
					Symbol:      "BTCUSD",
					Side:        "Buy",
					OrderType:   "Limit",
					Price:       "10000",
					Qty:         1,
					TimeInForce: "GoodTillCancel",
					CreateType:  "CreateByUser",
					CancelType:  "CancelByUser",
					OrderStatus: "",
					LeavesQty:   1,
					LeavesValue: "0",
					CreatedAt:   "2022-10-09T12:04:22.633327064Z",
					UpdatedAt:   "2022-10-09T12:04:22.72531757Z",
					CrossStatus: "PendingCancel",
					CrossSeq:    5542233074,
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

		_, err = client.Future().InversePerpetual().CancelAllOrder(param)
		assert.Error(t, err)
	})
}

func TestQueryOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := QueryOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/order"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryOrderResponse{
			Result: []QueryOrderResult{
				{
					UserID:      146940,
					PositionIdx: 0,
					Symbol:      "BTCUSD",
					Side:        "Buy",
					OrderType:   "Limit",
					Price:       "10000",
					Qty:         1,
					TimeInForce: "GoodTillCancel",
					OrderStatus: "New",
					ExtFields: map[string]interface{}{
						"o_req_num": float64(33922395),
					},
					LastExecTime: "1665320279.792137",
					LeavesQty:    1,
					LeavesValue:  "0.0001",
					CumExecQty:   0,
					CumExecValue: "",
					CumExecFee:   "",
					RejectReason: "EC_NoError",
					CancelType:   "UNKNOWN",
					OrderLinkID:  "",
					CreatedAt:    "2022-10-09T12:57:59.791972857Z",
					UpdatedAt:    "2022-10-09T12:57:59.794831698Z",
					OrderID:      "afbbe97c-4515-4a5f-8e0f-bc3e17766db3",
					TakeProfit:   "0.00",
					StopLoss:     "0.00",
					TpTriggerBy:  "UNKNOWN",
					SlTriggerBy:  "UNKNOWN",
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

		resp, err := client.Future().InversePerpetual().QueryOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := QueryOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/order"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryOrderResponse{
			Result: []QueryOrderResult{
				{
					UserID:      146940,
					PositionIdx: 0,
					Symbol:      "BTCUSD",
					Side:        "Buy",
					OrderType:   "Limit",
					Price:       "10000",
					Qty:         1,
					TimeInForce: "GoodTillCancel",
					OrderStatus: "New",
					ExtFields: map[string]interface{}{
						"o_req_num": float64(33922395),
					},
					LastExecTime: "1665320279.792137",
					LeavesQty:    1,
					LeavesValue:  "0.0001",
					CumExecQty:   0,
					CumExecValue: "",
					CumExecFee:   "",
					RejectReason: "EC_NoError",
					CancelType:   "UNKNOWN",
					OrderLinkID:  "",
					CreatedAt:    "2022-10-09T12:57:59.791972857Z",
					UpdatedAt:    "2022-10-09T12:57:59.794831698Z",
					OrderID:      "afbbe97c-4515-4a5f-8e0f-bc3e17766db3",
					TakeProfit:   "0.00",
					StopLoss:     "0.00",
					TpTriggerBy:  "UNKNOWN",
					SlTriggerBy:  "UNKNOWN",
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

		_, err = client.Future().InversePerpetual().QueryOrder(param)
		assert.Error(t, err)
	})
}

func TestCreateStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := 19400.5
		param := CreateStopOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolFutureBTCUSD,
			OrderType:   OrderTypeLimit,
			Qty:         1,
			BasePrice:   price,
			StopPx:      price + 200,
			TimeInForce: TimeInForceGoodTillCancel,
		}

		path := "/v2/private/stop-order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateStopOrderResponse{
			Result: CreateStopOrderResult{
				UserID:       146940,
				Symbol:       "BTCUSD",
				Side:         "Buy",
				OrderType:    "Market",
				Price:        "0.00",
				Qty:          "1",
				TimeInForce:  "ImmediateOrCancel",
				Remark:       "221.112.162.57",
				LeavesQty:    "1",
				LeavesValue:  "0.00000000",
				StopPx:       "19600.50",
				RejectReason: "EC_NoError",
				StopOrderID:  "0519f1e6-1188-4519-9a4c-34fe9a611169",
				OrderLinkID:  "",
				TriggerBy:    "LastPrice",
				BasePrice:    "19400.50",
				CreatedAt:    "2022-10-10T04:35:47.925Z",
				UpdatedAt:    "2022-10-10T04:35:47.925Z",
				TpTriggerBy:  "UNKNOWN",
				SlTriggerBy:  "UNKNOWN",
				TakeProfit:   "0.00",
				StopLoss:     "0.00",
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

		resp, err := client.Future().InversePerpetual().CreateStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := 19400.5
		param := CreateStopOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolFutureBTCUSD,
			OrderType:   OrderTypeLimit,
			Qty:         1,
			BasePrice:   price,
			StopPx:      price + 200,
			TimeInForce: TimeInForceGoodTillCancel,
		}

		path := "/v2/private/stop-order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateStopOrderResponse{
			Result: CreateStopOrderResult{
				UserID:       146940,
				Symbol:       "BTCUSD",
				Side:         "Buy",
				OrderType:    "Market",
				Price:        "0.00",
				Qty:          "1",
				TimeInForce:  "ImmediateOrCancel",
				Remark:       "221.112.162.57",
				LeavesQty:    "1",
				LeavesValue:  "0.00000000",
				StopPx:       "19600.50",
				RejectReason: "EC_NoError",
				StopOrderID:  "0519f1e6-1188-4519-9a4c-34fe9a611169",
				OrderLinkID:  "",
				TriggerBy:    "LastPrice",
				BasePrice:    "19400.50",
				CreatedAt:    "2022-10-10T04:35:47.925Z",
				UpdatedAt:    "2022-10-10T04:35:47.925Z",
				TpTriggerBy:  "UNKNOWN",
				SlTriggerBy:  "UNKNOWN",
				TakeProfit:   "0.00",
				StopLoss:     "0.00",
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

		_, err = client.Future().InversePerpetual().CreateStopOrder(param)
		assert.Error(t, err)
	})
}

func TestListStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := ListStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/stop-order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListStopOrderResponse{
			Result: ListStopOrderResult{
				ListStopOrders: []ListStopOrder{
					{
						UserID:          146940,
						PositionIdx:     0,
						StopOrderStatus: "Untriggered",
						Symbol:          "BTCUSD",
						Side:            "Buy",
						OrderType:       "Market",
						Price:           "0",
						Qty:             "1",
						TimeInForce:     "ImmediateOrCancel",
						StopOrderType:   "Stop",
						TriggerBy:       "LastPrice",
						BasePrice:       "19400.50",
						OrderLinkID:     "",
						CreatedAt:       "2022-10-10T09:51:46.81Z",
						UpdatedAt:       "2022-10-10T09:51:46.81Z",
						StopPx:          "19600.50",
						StopOrderID:     "647479dd-5679-48e8-88d9-0664d5ce4d2a",
						TakeProfit:      "0.00",
						StopLoss:        "0.00",
						TpTriggerBy:     "UNKNOWN",
						SlTriggerBy:     "UNKNOWN",
						Cursor:          "",
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

		resp, err := client.Future().InversePerpetual().ListStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := ListStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/stop-order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListStopOrderResponse{
			Result: ListStopOrderResult{
				ListStopOrders: []ListStopOrder{
					{
						UserID:          146940,
						PositionIdx:     0,
						StopOrderStatus: "Untriggered",
						Symbol:          "BTCUSD",
						Side:            "Buy",
						OrderType:       "Market",
						Price:           "0",
						Qty:             "1",
						TimeInForce:     "ImmediateOrCancel",
						StopOrderType:   "Stop",
						TriggerBy:       "LastPrice",
						BasePrice:       "19400.50",
						OrderLinkID:     "",
						CreatedAt:       "2022-10-10T09:51:46.81Z",
						UpdatedAt:       "2022-10-10T09:51:46.81Z",
						StopPx:          "19600.50",
						StopOrderID:     "647479dd-5679-48e8-88d9-0664d5ce4d2a",
						TakeProfit:      "0.00",
						StopLoss:        "0.00",
						TpTriggerBy:     "UNKNOWN",
						SlTriggerBy:     "UNKNOWN",
						Cursor:          "",
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

		_, err = client.Future().InversePerpetual().ListStopOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stopOrderID := "f13236dd-4988-41e8-8558-011d875d4282"
		param := CancelStopOrderParam{
			Symbol:      SymbolFutureBTCUSD,
			StopOrderID: &stopOrderID,
		}

		path := "/v2/private/stop-order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelStopOrderResponse{
			Result: CancelStopOrderResult{
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

		resp, err := client.Future().InversePerpetual().CancelStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := CancelStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/stop-order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelStopOrderResponse{
			Result: CancelStopOrderResult{
				StopOrderID: "f13236dd-4988-41e8-8558-011d875d4282",
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

		_, err = client.Future().InversePerpetual().CancelStopOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelAllStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := CancelAllStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/stop-order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllStopOrderResponse{
			Result: []CancelAllStopOrderResult{
				{
					ClOrdID:           "64b6bc83-cfa0-4d77-9443-98109942704c",
					OrderLinkID:       "",
					UserID:            146940,
					Symbol:            "BTCUSD",
					Side:              "Buy",
					OrderType:         "Market",
					Price:             "0",
					Qty:               1,
					TimeInForce:       "ImmediateOrCancel",
					CreateType:        "CreateByStopOrder",
					CancelType:        "CancelByUser",
					OrderStatus:       "",
					LeavesQty:         0,
					LeavesValue:       "0",
					CreatedAt:         "2022-10-11T09:42:45.847223316Z",
					UpdatedAt:         "2022-10-11T09:42:45.930763726Z",
					CrossStatus:       "Deactivated",
					CrossSeq:          -1,
					StopOrderType:     "Stop",
					TriggerBy:         "LastPrice",
					BasePrice:         "19400.5",
					ExpectedDirection: "Rising",
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

		resp, err := client.Future().InversePerpetual().CancelAllStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := CancelAllStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/stop-order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllStopOrderResponse{
			Result: []CancelAllStopOrderResult{
				{
					ClOrdID:           "64b6bc83-cfa0-4d77-9443-98109942704c",
					OrderLinkID:       "",
					UserID:            146940,
					Symbol:            "BTCUSD",
					Side:              "Buy",
					OrderType:         "Market",
					Price:             "0",
					Qty:               1,
					TimeInForce:       "ImmediateOrCancel",
					CreateType:        "CreateByStopOrder",
					CancelType:        "CancelByUser",
					OrderStatus:       "",
					LeavesQty:         0,
					LeavesValue:       "0",
					CreatedAt:         "2022-10-11T09:42:45.847223316Z",
					UpdatedAt:         "2022-10-11T09:42:45.930763726Z",
					CrossStatus:       "Deactivated",
					CrossSeq:          -1,
					StopOrderType:     "Stop",
					TriggerBy:         "LastPrice",
					BasePrice:         "19400.5",
					ExpectedDirection: "Rising",
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

		_, err = client.Future().InversePerpetual().CancelAllStopOrder(param)
		assert.Error(t, err)
	})
}

func TestQueryStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := QueryStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/stop-order"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryStopOrderResponse{
			Result: []QueryStopOrderResult{
				{
					UserID:          146940,
					PositionIdx:     0,
					Symbol:          "BTCUSD",
					Side:            "Buy",
					OrderType:       "Market",
					Price:           "0",
					Qty:             1,
					StopPx:          "19600.50",
					BasePrice:       "19400.50",
					TimeInForce:     "ImmediateOrCancel",
					StopOrderStatus: "",
					ExtFields: map[string]interface{}{
						"o_req_num": float64(43061877),
					},
					LeavesQty:    1,
					LeavesValue:  "0",
					CumExecQty:   0,
					CumExecValue: "",
					CumExecFee:   "",
					RejectReason: "EC_NoError",
					OrderLinkID:  "",
					CreatedAt:    "2022-10-12T07:00:10.006460855Z",
					UpdatedAt:    "2022-10-12T07:00:10.006460855Z",
					OrderID:      "8891fda6-eb85-4d22-af1b-27833bbfe276",
					TriggerBy:    "LastPrice",
					TakeProfit:   "0.00",
					StopLoss:     "0.00",
					TpTriggerBy:  "UNKNOWN",
					SlTriggerBy:  "UNKNOWN",
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

		resp, err := client.Future().InversePerpetual().QueryStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := QueryStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/v2/private/stop-order"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryStopOrderResponse{
			Result: []QueryStopOrderResult{
				{
					UserID:          146940,
					PositionIdx:     0,
					Symbol:          "BTCUSD",
					Side:            "Buy",
					OrderType:       "Market",
					Price:           "0",
					Qty:             1,
					StopPx:          "19600.50",
					BasePrice:       "19400.50",
					TimeInForce:     "ImmediateOrCancel",
					StopOrderStatus: "",
					ExtFields: map[string]interface{}{
						"o_req_num": float64(43061877),
					},
					LeavesQty:    1,
					LeavesValue:  "0",
					CumExecQty:   0,
					CumExecValue: "",
					CumExecFee:   "",
					RejectReason: "EC_NoError",
					OrderLinkID:  "",
					CreatedAt:    "2022-10-12T07:00:10.006460855Z",
					UpdatedAt:    "2022-10-12T07:00:10.006460855Z",
					OrderID:      "8891fda6-eb85-4d22-af1b-27833bbfe276",
					TriggerBy:    "LastPrice",
					TakeProfit:   "0.00",
					StopLoss:     "0.00",
					TpTriggerBy:  "UNKNOWN",
					SlTriggerBy:  "UNKNOWN",
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

		_, err = client.Future().InversePerpetual().QueryStopOrder(param)
		assert.Error(t, err)
	})
}

func TestTradingStop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := 20000.0
		param := TradingStopParam{
			Symbol:     SymbolFutureBTCUSD,
			TakeProfit: &price,
		}

		path := "/v2/private/position/trading-stop"
		method := http.MethodPost
		status := http.StatusOK
		respBody := TradingStopResponse{
			Result: TradingStopResult{
				ID:                  0,
				UserID:              146940,
				Symbol:              "BTCUSD",
				Side:                "Buy",
				Size:                3,
				PositionValue:       0.00015564,
				EntryPrice:          19275.25057826,
				RiskID:              1,
				AutoAddMargin:       0,
				Leverage:            2,
				PositionMargin:      0.00007782,
				LiqPrice:            12893.5,
				BustPrice:           12850.5,
				OccClosingFee:       1.5e-7,
				OccFundingFee:       0,
				TakeProfit:          20000,
				StopLoss:            0,
				TrailingStop:        0,
				PositionStatus:      "Normal",
				DeleverageIndicator: 2,
				OcCalcData:          "{\"blq\":0,\"slq\":0,\"bmp\":0,\"smp\":0,\"fq\":-3,\"bv2c\":0.5015,\"sv2c\":0.5009}",
				OrderMargin:         0,
				WalletBalance:       0.07516902,
				RealisedPnl:         -3.7e-7,
				CumRealisedPnl:      -0.00064488,
				CumCommission:       0,
				CrossSeq:            5561803978,
				PositionSeq:         0,
				CreatedAt:           "2021-03-06T09:44:57.204724626Z",
				UpdatedAt:           "2022-10-24T09:22:53.911081802Z",
				ExtFields: map[string]interface{}{
					"mm":              float64(0),
					"tp_trigger_by":   "LastPrice",
					"trailing_active": "0",
					"v":               float64(799708),
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

		resp, err := client.Future().InversePerpetual().TradingStop(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := 20000.0
		param := TradingStopParam{
			Symbol:     SymbolFutureBTCUSD,
			TakeProfit: &price,
		}

		path := "/v2/private/position/trading-stop"
		method := http.MethodPost
		status := http.StatusOK
		respBody := TradingStopResponse{
			Result: TradingStopResult{
				ID:                  0,
				UserID:              146940,
				Symbol:              "BTCUSD",
				Side:                "Buy",
				Size:                3,
				PositionValue:       0.00015564,
				EntryPrice:          19275.25057826,
				RiskID:              1,
				AutoAddMargin:       0,
				Leverage:            2,
				PositionMargin:      0.00007782,
				LiqPrice:            12893.5,
				BustPrice:           12850.5,
				OccClosingFee:       1.5e-7,
				OccFundingFee:       0,
				TakeProfit:          20000,
				StopLoss:            0,
				TrailingStop:        0,
				PositionStatus:      "Normal",
				DeleverageIndicator: 2,
				OcCalcData:          "{\"blq\":0,\"slq\":0,\"bmp\":0,\"smp\":0,\"fq\":-3,\"bv2c\":0.5015,\"sv2c\":0.5009}",
				OrderMargin:         0,
				WalletBalance:       0.07516902,
				RealisedPnl:         -3.7e-7,
				CumRealisedPnl:      -0.00064488,
				CumCommission:       0,
				CrossSeq:            5561803978,
				PositionSeq:         0,
				CreatedAt:           "2021-03-06T09:44:57.204724626Z",
				UpdatedAt:           "2022-10-24T09:22:53.911081802Z",
				ExtFields: map[string]interface{}{
					"mm":              float64(0),
					"tp_trigger_by":   "LastPrice",
					"trailing_active": "0",
					"v":               float64(799708),
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

		_, err = client.Future().InversePerpetual().TradingStop(param)
		assert.Error(t, err)
	})
}

func TestAPIKeyInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := "/v2/private/account/api-key"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": []map[string]interface{}{
				{
					"api_key":    "TrswBXTLDVhxqkfDEI",
					"type":       "personal",
					"user_id":    146940,
					"inviter_id": 0,
					"ips": []string{
						"*",
					},
					"note": "full3",
					"permissions": []string{
						"OptionsTrade",
						"CopyTrading",
						"ExchangeHistory",
						"SpotTrade",
						"AccountTransfer",
						"SubMemberTransfer",
						"Order",
						"Position",
						"DerivativesTrade",
					},
					"created_at":      "2022-06-20T13:26:50Z",
					"expired_at":      "2023-04-23T05:52:24Z",
					"read_only":       false,
					"vip_level":       "No VIP",
					"mkt_maker_level": "0",
					"affiliate_id":    0,
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

		resp, err := client.Future().InversePerpetual().APIKeyInfo()
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		path := "/v2/private/account/api-key"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": []map[string]interface{}{
				{
					"api_key":    "TrswBXTLDVhxqkfDEI",
					"type":       "personal",
					"user_id":    146940,
					"inviter_id": 0,
					"ips": []string{
						"*",
					},
					"note": "full3",
					"permissions": []string{
						"OptionsTrade",
						"CopyTrading",
						"ExchangeHistory",
						"SpotTrade",
						"AccountTransfer",
						"SubMemberTransfer",
						"Order",
						"Position",
						"DerivativesTrade",
					},
					"created_at":      "2022-06-20T13:26:50Z",
					"expired_at":      "2023-04-23T05:52:24Z",
					"read_only":       false,
					"vip_level":       "No VIP",
					"mkt_maker_level": "0",
					"affiliate_id":    0,
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

		_, err = client.Future().InversePerpetual().APIKeyInfo()
		assert.Error(t, err)
	})
}
