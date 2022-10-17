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
