package bybit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dimkus/bybit/v2/testhelper"
)

func TestV5WebsocketPrivate_Order(t *testing.T) {
	respBody := V5WebsocketPrivateOrderResponse{
		Topic:        "order",
		ID:           "75d86e42f18b23b9ad2c1f10eaffa8bb:18483ff242aca593:0:01",
		CreationTime: 1677226839837,
		Data: []V5WebsocketPrivateOrderData{
			{
				AvgPrice:           "23090.80",
				BlockTradeID:       "",
				CancelType:         "UNKNOWN",
				Category:           "linear",
				CloseOnTrigger:     false,
				CreatedTime:        "1677375772152",
				CumExecFee:         "0.01385448",
				CumExecQty:         "0.001",
				CumExecValue:       "23.0908",
				LeavesQty:          "0",
				LeavesValue:        "0",
				OrderID:            "cd770a60-549f-433b-8000-aacefec3c7c3",
				OrderIv:            "",
				IsLeverage:         "",
				LastPriceOnCreated: "23089.30",
				OrderStatus:        "Filled",
				OrderLinkID:        "",
				OrderType:          "Market",
				PositionIdx:        1,
				Price:              "24243.70",
				Qty:                "0.001",
				ReduceOnly:         false,
				RejectReason:       "EC_NoError",
				Side:               "Buy",
				SlTriggerBy:        "UNKNOWN",
				StopLoss:           "0.00",
				StopOrderType:      "UNKNOWN",
				Symbol:             "BTCUSDT",
				TakeProfit:         "0.00",
				TimeInForce:        "IOC",
				TpTriggerBy:        "UNKNOWN",
				TriggerBy:          "UNKNOWN",
				TriggerDirection:   0,
				TriggerPrice:       "0.00",
				UpdatedTime:        "1677375772154",
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewWebsocketServer(
		testhelper.WithWebsocketHandlerOption(V5WebsocketPrivatePath, bytesBody),
	)
	defer teardown()

	wsClient := NewTestWebsocketClient().
		WithBaseURL(server.URL).
		WithAuth("test", "test")

	svc, err := wsClient.V5().Private(CategoryV5All)
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	{
		_, err := svc.SubscribeOrder(func(response V5WebsocketPrivateOrderResponse) error {
			assert.Equal(t, respBody, response)
			return nil
		})
		require.NoError(t, err)
	}

	assert.NoError(t, svc.Run())
	assert.NoError(t, svc.Ping())
	assert.NoError(t, svc.Close())
}
