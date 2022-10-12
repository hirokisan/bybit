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
