package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/testhelper"
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
