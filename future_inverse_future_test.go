package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFuturesOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := 10000.0
		param := CreateFuturesOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolFutureBTCUSD,
			OrderType:   OrderTypeLimit,
			Qty:         10,
			TimeInForce: TimeInForceGoodTillCancel,
			Price:       &price,
		}

		path := "/futures/private/order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateFuturesOrderResponse{
			Result: CreateFuturesOrderResult{
				UserID:        146940,
				OrderID:       "9192e7cb-2794-4fcb-97e9-42b4a7a4c2bb",
				Symbol:        "BTCUSD",
				Side:          "Buy",
				OrderType:     "Limit",
				Price:         10000,
				Qty:           1,
				TimeInForce:   "GoodTillCancel",
				OrderStatus:   "Created",
				LastExecTime:  0,
				LastExecPrice: 0,
				LeavesQty:     1,
				CumExecQty:    0,
				CumExecValue:  0,
				CumExecFee:    0,
				RejectReason:  "EC_NoError",
				OrderLinkID:   "",
				CreatedAt:     "2022-10-16T00:43:23.105Z",
				UpdatedAt:     "2022-10-16T00:43:23.105Z",
				TakeProfit:    "0.00",
				StopLoss:      "0.00",
				TpTriggerBy:   "UNKNOWN",
				SlTriggerBy:   "UNKNOWN",
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

		resp, err := client.Future().InverseFuture().CreateFuturesOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := 10000.0
		param := CreateFuturesOrderParam{
			Side:        SideBuy,
			Symbol:      SymbolFutureBTCUSD,
			OrderType:   OrderTypeLimit,
			Qty:         10,
			TimeInForce: TimeInForceGoodTillCancel,
			Price:       &price,
		}

		path := "/futures/private/order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateFuturesOrderResponse{
			Result: CreateFuturesOrderResult{
				UserID:        146940,
				OrderID:       "9192e7cb-2794-4fcb-97e9-42b4a7a4c2bb",
				Symbol:        "BTCUSD",
				Side:          "Buy",
				OrderType:     "Limit",
				Price:         10000,
				Qty:           1,
				TimeInForce:   "GoodTillCancel",
				OrderStatus:   "Created",
				LastExecTime:  0,
				LastExecPrice: 0,
				LeavesQty:     1,
				CumExecQty:    0,
				CumExecValue:  0,
				CumExecFee:    0,
				RejectReason:  "EC_NoError",
				OrderLinkID:   "",
				CreatedAt:     "2022-10-16T00:43:23.105Z",
				UpdatedAt:     "2022-10-16T00:43:23.105Z",
				TakeProfit:    "0.00",
				StopLoss:      "0.00",
				TpTriggerBy:   "UNKNOWN",
				SlTriggerBy:   "UNKNOWN",
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

		_, err = client.Future().InverseFuture().CreateFuturesOrder(param)
		assert.Error(t, err)
	})
}

func TestFuturesListOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderStatus := OrderStatusNew
		param := ListFuturesOrderParam{
			Symbol:      SymbolFutureBTCUSD,
			OrderStatus: &orderStatus,
		}

		path := "/futures/private/order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListFuturesOrderResponse{
			Result: ListFuturesOrderResult{
				ListFuturesOrders: []ListFuturesOrder{
					{
						UserID:       146940,
						PositionIdx:  0,
						Symbol:       "BTCUSD",
						Side:         "Buy",
						OrderType:    "Limit",
						Price:        "10000",
						Qty:          "1",
						TimeInForce:  "GoodTillCancel",
						OrderLinkID:  "",
						OrderID:      "04fe50ec-ea81-4735-8046-9d836f9af8c7",
						CreatedAt:    "2022-10-16T08:35:38.852Z",
						UpdatedAt:    "2022-10-16T08:35:38.855Z",
						OrderStatus:  "New",
						LeavesQty:    "1",
						LeavesValue:  "0.0001",
						CumExecQty:   "0",
						CumExecValue: "0",
						CumExecFee:   "0",
						RejectReason: "EC_NoError",
						TakeProfit:   "0.00",
						StopLoss:     "0.00",
						TpTriggerBy:  "UNKNOWN",
						SlTriggerBy:  "UNKNOWN",
						Cursor:       "",
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

		resp, err := client.Future().InverseFuture().ListFuturesOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		orderStatus := OrderStatusNew
		param := ListFuturesOrderParam{
			Symbol:      SymbolFutureBTCUSD,
			OrderStatus: &orderStatus,
		}

		path := "/futures/private/order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListFuturesOrderResponse{
			Result: ListFuturesOrderResult{
				ListFuturesOrders: []ListFuturesOrder{
					{
						UserID:       146940,
						PositionIdx:  0,
						Symbol:       "BTCUSD",
						Side:         "Buy",
						OrderType:    "Limit",
						Price:        "10000",
						Qty:          "1",
						TimeInForce:  "GoodTillCancel",
						OrderLinkID:  "",
						OrderID:      "04fe50ec-ea81-4735-8046-9d836f9af8c7",
						CreatedAt:    "2022-10-16T08:35:38.852Z",
						UpdatedAt:    "2022-10-16T08:35:38.855Z",
						OrderStatus:  "New",
						LeavesQty:    "1",
						LeavesValue:  "0.0001",
						CumExecQty:   "0",
						CumExecValue: "0",
						CumExecFee:   "0",
						RejectReason: "EC_NoError",
						TakeProfit:   "0.00",
						StopLoss:     "0.00",
						TpTriggerBy:  "UNKNOWN",
						SlTriggerBy:  "UNKNOWN",
						Cursor:       "",
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

		_, err = client.Future().InverseFuture().ListFuturesOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelFuturesOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderID := "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787"
		param := CancelFuturesOrderParam{
			Symbol:  SymbolFutureBTCUSD,
			OrderID: &orderID,
		}

		path := "/futures/private/order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelFuturesOrderResponse{
			Result: CancelFuturesOrderResult{
				UserID:        146940,
				OrderID:       "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787",
				Symbol:        "BTCUSD",
				Side:          "Buy",
				OrderType:     "Limit",
				Price:         10000,
				Qty:           1,
				TimeInForce:   "GoodTillCancel",
				OrderStatus:   "New",
				LastExecTime:  1665991158.845237,
				LastExecPrice: 0,
				LeavesQty:     1,
				CumExecQty:    0,
				CumExecValue:  0,
				CumExecFee:    0,
				RejectReason:  "EC_NoError",
				OrderLinkID:   "",
				CreatedAt:     "2022-10-17T07:19:18.845Z",
				UpdatedAt:     "2022-10-17T07:19:19.067Z",
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

		resp, err := client.Future().InverseFuture().CancelFuturesOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		orderID := "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787"
		param := CancelFuturesOrderParam{
			Symbol:  SymbolFutureBTCUSD,
			OrderID: &orderID,
		}

		path := "/futures/private/order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelFuturesOrderResponse{
			Result: CancelFuturesOrderResult{
				UserID:        146940,
				OrderID:       "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787",
				Symbol:        "BTCUSD",
				Side:          "Buy",
				OrderType:     "Limit",
				Price:         10000,
				Qty:           1,
				TimeInForce:   "GoodTillCancel",
				OrderStatus:   "New",
				LastExecTime:  1665991158.845237,
				LastExecPrice: 0,
				LeavesQty:     1,
				CumExecQty:    0,
				CumExecValue:  0,
				CumExecFee:    0,
				RejectReason:  "EC_NoError",
				OrderLinkID:   "",
				CreatedAt:     "2022-10-17T07:19:18.845Z",
				UpdatedAt:     "2022-10-17T07:19:19.067Z",
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

		_, err = client.Future().InverseFuture().CancelFuturesOrder(param)
		assert.Error(t, err)
	})
}

func TestAllCancelFuturesOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := CancelAllFuturesOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/futures/private/order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllFuturesOrderResponse{
			Result: []CancelAllFuturesOrderResult{
				{
					ClOrdID:     "4ad4e9c5-4aa0-40bd-8a4d-53c4b4ca49ef",
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
					CreatedAt:   "2022-10-18T08:10:15.341270694Z",
					UpdatedAt:   "2022-10-18T08:10:15.426826461Z",
					CrossStatus: "PendingCancel",
					CrossSeq:    5554666653,
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

		resp, err := client.Future().InverseFuture().CancelAllFuturesOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := CancelAllFuturesOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/futures/private/order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllFuturesOrderResponse{
			Result: []CancelAllFuturesOrderResult{
				{
					ClOrdID:     "4ad4e9c5-4aa0-40bd-8a4d-53c4b4ca49ef",
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
					CreatedAt:   "2022-10-18T08:10:15.341270694Z",
					UpdatedAt:   "2022-10-18T08:10:15.426826461Z",
					CrossStatus: "PendingCancel",
					CrossSeq:    5554666653,
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

		_, err = client.Future().InverseFuture().CancelAllFuturesOrder(param)
		assert.Error(t, err)
	})
}

func TestQueryFuturesOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderID := "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787"
		param := QueryFuturesOrderParam{
			Symbol:  SymbolFutureBTCUSD,
			OrderID: &orderID,
		}

		path := "/futures/private/order"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryFuturesOrderResponse{
			Result: QueryFuturesOrderResult{
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
					"o_req_num": float64(2492237),
				},
				LastExecTime: "1666138015.344325",
				LeavesQty:    1,
				LeavesValue:  "0.0001",
				CumExecQty:   0,
				CumExecValue: "",
				CumExecFee:   "",
				RejectReason: "EC_NoError",
				CancelType:   "UNKNOWN",
				OrderLinkID:  "",
				CreatedAt:    "2022-10-19T00:06:55.3441555Z",
				UpdatedAt:    "2022-10-19T00:06:55.347065879Z",
				OrderID:      "40a03f42-db73-4395-8ce9-9ac720f9ffdb",
				TakeProfit:   "0.00",
				StopLoss:     "0.00",
				TpTriggerBy:  "UNKNOWN",
				SlTriggerBy:  "UNKNOWN",
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

		resp, err := client.Future().InverseFuture().QueryFuturesOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		orderID := "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787"
		param := QueryFuturesOrderParam{
			Symbol:  SymbolFutureBTCUSD,
			OrderID: &orderID,
		}

		path := "/futures/private/order"
		method := http.MethodGet
		status := http.StatusOK
		respBody := QueryFuturesOrderResponse{
			Result: QueryFuturesOrderResult{
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
					"o_req_num": float64(2492237),
				},
				LastExecTime: "1666138015.344325",
				LeavesQty:    1,
				LeavesValue:  "0.0001",
				CumExecQty:   0,
				CumExecValue: "",
				CumExecFee:   "",
				RejectReason: "EC_NoError",
				CancelType:   "UNKNOWN",
				OrderLinkID:  "",
				CreatedAt:    "2022-10-19T00:06:55.3441555Z",
				UpdatedAt:    "2022-10-19T00:06:55.347065879Z",
				OrderID:      "40a03f42-db73-4395-8ce9-9ac720f9ffdb",
				TakeProfit:   "0.00",
				StopLoss:     "0.00",
				TpTriggerBy:  "UNKNOWN",
				SlTriggerBy:  "UNKNOWN",
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

		_, err = client.Future().InverseFuture().QueryFuturesOrder(param)
		assert.Error(t, err)
	})
}

func TestCreateFuturesStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := 19400.5
		symbol := SymbolFutureBTCUSD
		param := CreateFuturesStopOrderParam{
			Side:        SideBuy,
			Symbol:      symbol,
			OrderType:   OrderTypeMarket,
			Qty:         1,
			BasePrice:   price,
			StopPx:      price + 200,
			TimeInForce: TimeInForceGoodTillCancel,
		}

		path := "/futures/private/stop-order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateFuturesStopOrderResponse{
			Result: CreateFuturesStopOrderResult{
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
				StopOrderID:  "6f30b476-af36-4ccd-88c1-9f49ad598fdf",
				OrderLinkID:  "",
				TriggerBy:    "LastPrice",
				BasePrice:    "19400.50",
				CreatedAt:    "2022-10-19T10:49:47.920Z",
				UpdatedAt:    "2022-10-19T10:49:47.920Z",
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

		resp, err := client.Future().InverseFuture().CreateFuturesStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := 19400.5
		symbol := SymbolFutureBTCUSD
		param := CreateFuturesStopOrderParam{
			Side:        SideBuy,
			Symbol:      symbol,
			OrderType:   OrderTypeMarket,
			Qty:         1,
			BasePrice:   price,
			StopPx:      price + 200,
			TimeInForce: TimeInForceGoodTillCancel,
		}

		path := "/futures/private/stop-order/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CreateFuturesStopOrderResponse{
			Result: CreateFuturesStopOrderResult{
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
				StopOrderID:  "6f30b476-af36-4ccd-88c1-9f49ad598fdf",
				OrderLinkID:  "",
				TriggerBy:    "LastPrice",
				BasePrice:    "19400.50",
				CreatedAt:    "2022-10-19T10:49:47.920Z",
				UpdatedAt:    "2022-10-19T10:49:47.920Z",
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

		_, err = client.Future().InverseFuture().CreateFuturesStopOrder(param)
		assert.Error(t, err)
	})
}

func TestFuturesListStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderStatus := OrderStatusUntriggered
		param := ListFuturesStopOrderParam{
			Symbol:          SymbolFutureBTCUSDH23,
			StopOrderStatus: &orderStatus,
		}

		path := "/futures/private/stop-order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListFuturesStopOrderResponse{
			Result: ListFuturesStopOrderResult{
				ListFuturesStopOrders: []ListFuturesStopOrder{
					{
						UserID:          146940,
						PositionIdx:     0,
						StopOrderStatus: "Untriggered",
						Symbol:          "BTCUSDH23",
						Side:            "Buy",
						OrderType:       "Market",
						StopOrderType:   "Stop",
						Price:           "0",
						Qty:             "1",
						TimeInForce:     "ImmediateOrCancel",
						BasePrice:       "19400.50",
						OrderLinkID:     "",
						CreatedAt:       "2022-10-19T12:49:49.063Z",
						UpdatedAt:       "2022-10-19T12:49:49.063Z",
						StopPx:          "19600.50",
						StopOrderID:     "9df68792-204a-4394-8e92-5a8c7827d99d",
						TriggerBy:       "LastPrice",
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

		resp, err := client.Future().InverseFuture().ListFuturesStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		orderStatus := OrderStatusUntriggered
		param := ListFuturesStopOrderParam{
			Symbol:          SymbolFutureBTCUSDH23,
			StopOrderStatus: &orderStatus,
		}

		path := "/futures/private/stop-order/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := ListFuturesStopOrderResponse{
			Result: ListFuturesStopOrderResult{
				ListFuturesStopOrders: []ListFuturesStopOrder{
					{
						UserID:          146940,
						PositionIdx:     0,
						StopOrderStatus: "Untriggered",
						Symbol:          "BTCUSDH23",
						Side:            "Buy",
						OrderType:       "Market",
						StopOrderType:   "Stop",
						Price:           "0",
						Qty:             "1",
						TimeInForce:     "ImmediateOrCancel",
						BasePrice:       "19400.50",
						OrderLinkID:     "",
						CreatedAt:       "2022-10-19T12:49:49.063Z",
						UpdatedAt:       "2022-10-19T12:49:49.063Z",
						StopPx:          "19600.50",
						StopOrderID:     "9df68792-204a-4394-8e92-5a8c7827d99d",
						TriggerBy:       "LastPrice",
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

		_, err = client.Future().InverseFuture().ListFuturesStopOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelFuturesStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stopOrderID := "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787"
		param := CancelFuturesStopOrderParam{
			Symbol:      SymbolFutureBTCUSD,
			StopOrderID: &stopOrderID,
		}

		path := "/futures/private/stop-order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelFuturesStopOrderResponse{
			Result: CancelFuturesStopOrderResult{
				StopOrderID: "7b9b5a49-8aea-4684-a1ac-4df472dc789c",
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

		resp, err := client.Future().InverseFuture().CancelFuturesStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		stopOrderID := "9dde0c9e-7e10-4d9a-8da9-1ef0f976c787"
		param := CancelFuturesStopOrderParam{
			Symbol:      SymbolFutureBTCUSD,
			StopOrderID: &stopOrderID,
		}

		path := "/futures/private/stop-order/cancel"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelFuturesStopOrderResponse{
			Result: CancelFuturesStopOrderResult{
				StopOrderID: "7b9b5a49-8aea-4684-a1ac-4df472dc789c",
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

		_, err = client.Future().InverseFuture().CancelFuturesStopOrder(param)
		assert.Error(t, err)
	})
}

func TestCancelAllFuturesStopOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := CancelAllFuturesStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/futures/private/stop-order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllFuturesStopOrderResponse{
			Result: []CancelAllFuturesStopOrderResult{
				{
					ClOrdID:           "56d4e0b1-ce57-46d9-8d1d-216f118399d8",
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
					CreatedAt:         "2022-10-21T09:34:28.326445671Z",
					UpdatedAt:         "2022-10-21T09:34:28.411727842Z",
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

		resp, err := client.Future().InverseFuture().CancelAllFuturesStopOrder(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		assert.Equal(t, respBody, *resp)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := CancelAllFuturesStopOrderParam{
			Symbol: SymbolFutureBTCUSD,
		}

		path := "/futures/private/stop-order/cancelAll"
		method := http.MethodPost
		status := http.StatusOK
		respBody := CancelAllFuturesStopOrderResponse{
			Result: []CancelAllFuturesStopOrderResult{
				{
					ClOrdID:           "56d4e0b1-ce57-46d9-8d1d-216f118399d8",
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
					CreatedAt:         "2022-10-21T09:34:28.326445671Z",
					UpdatedAt:         "2022-10-21T09:34:28.411727842Z",
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

		_, err = client.Future().InverseFuture().CancelAllFuturesStopOrder(param)
		assert.Error(t, err)
	})
}
