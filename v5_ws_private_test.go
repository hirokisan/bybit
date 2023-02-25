package bybit

import (
	"encoding/json"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5WebsocketPrivate_Position(t *testing.T) {
	respBody := V5WebsocketPrivatePositionResponseContent{
		Topic:        "position",
		ID:           "75d86e42f18b23b9ad2c1f10eaffa8bb:18483ff242aca593:0:01",
		CreationTime: 1677226839837,
		Data: []V5WebsocketPrivatePositionResponseData{
			{
				BustPrice:       "0.10",
				Category:        "linear",
				CreatedTime:     "1666729101704",
				CumRealisedPnl:  "-4.77363636",
				EntryPrice:      "23780.1",
				Leverage:        "10",
				LiqPrice:        "0.10",
				MarkPrice:       "23782.27",
				PositionBalance: "2.39085126",
				PositionIdx:     1,
				PositionMM:      "0.0000005",
				PositionIM:      "0.237801",
				PositionStatus:  "Normal",
				PositionValue:   "23.7801",
				RiskID:          1,
				RiskLimitValue:  "2000000",
				Side:            "Buy",
				Size:            "0.001",
				StopLoss:        "0.00",
				Symbol:          "BTCUSDT",
				TakeProfit:      "0.00",
				TpSlMode:        "Full",
				TradeMode:       0,
				AutoAddMargin:   0,
				TrailingStop:    "0.00",
				UnrealisedPnl:   "0.00217",
				UpdatedTime:     "1677226839835",
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

	svc, err := wsClient.V5().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	require.NoError(t, svc.RegisterFuncPosition(func(response V5WebsocketPrivatePositionResponseContent) error {
		assert.Equal(t, respBody, response)
		return nil
	}))

	assert.NoError(t, svc.Run())
	assert.NoError(t, svc.Ping())
	assert.NoError(t, svc.Close())
}
