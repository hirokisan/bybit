package bybit

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		WithAuth("test", "test").
		WithDialer(&websocket.Dialer{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return nil, nil
			},
			HandshakeTimeout: 5 * time.Second,
		})

	svc, err := wsClient.V5().Private()
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

func TestV5WebsocketPrivate_Order_ParentOrderLinkID(t *testing.T) {
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
				ParentOrderLinkID:  "entry-123",
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
	bytesBody := []byte(`{
		"topic": "order",
		"id": "75d86e42f18b23b9ad2c1f10eaffa8bb:18483ff242aca593:0:01",
		"creationTime": 1677226839837,
		"data": [{
			"avgPrice": "23090.80",
			"blockTradeId": "",
			"cancelType": "UNKNOWN",
			"category": "linear",
			"closeOnTrigger": false,
			"createdTime": "1677375772152",
			"cumExecFee": "0.01385448",
			"cumExecQty": "0.001",
			"cumExecValue": "23.0908",
			"leavesQty": "0",
			"leavesValue": "0",
			"orderId": "cd770a60-549f-433b-8000-aacefec3c7c3",
			"orderIv": "",
			"isLeverage": "",
			"lastPriceOnCreated": "23089.30",
			"orderStatus": "Filled",
			"orderLinkId": "",
			"parentOrderLinkId": "entry-123",
			"orderType": "Market",
			"positionIdx": 1,
			"price": "24243.70",
			"qty": "0.001",
			"reduceOnly": false,
			"rejectReason": "EC_NoError",
			"side": "Buy",
			"slTriggerBy": "UNKNOWN",
			"stopLoss": "0.00",
			"stopOrderType": "UNKNOWN",
			"symbol": "BTCUSDT",
			"takeProfit": "0.00",
			"timeInForce": "IOC",
			"tpTriggerBy": "UNKNOWN",
			"triggerBy": "UNKNOWN",
			"triggerDirection": 0,
			"triggerPrice": "0.00",
			"updatedTime": "1677375772154"
		}]
	}`)

	var response V5WebsocketPrivateOrderResponse
	err := json.Unmarshal(bytesBody, &response)
	require.NoError(t, err)
	assert.Equal(t, respBody, response)
}
