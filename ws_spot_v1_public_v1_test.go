package bybit

import (
	"encoding/json"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpotWebsocketV1PublicV1Trade(t *testing.T) {
	respBody := SpotWebsocketV1PublicV1TradeResponse{
		Symbol:         SymbolSpotBTCUSDT,
		SymbolName:     string(SymbolSpotBTCUSDT),
		Topic:          SpotWebsocketV1PublicV1TopicTrade,
		SendTime:       1664284020685,
		IsFirstMessage: true,
		Params: SpotWebsocketV1PublicV1TradeResponseParams{
			RealtimeInterval: "24h",
			Binary:           "false",
		},
		Data: []SpotWebsocketV1PublicV1TradeContent{
			{
				TradeID:        "2100000000002571479",
				Timestamp:      1664283342503,
				Price:          "20191.69",
				Quantity:       "0.000495",
				IsBuySideTaker: true,
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewWebsocketServer(
		testhelper.WithWebsocketHandlerOption(SpotWebsocketV1PublicV1Path, bytesBody),
	)
	defer teardown()

	wsClient := NewTestWebsocketClient().
		WithBaseURL(server.URL)

	svc, err := wsClient.Spot().V1().PublicV1()
	require.NoError(t, err)

	unsubscribe, err := svc.SubscribeTrade(SymbolSpotBTCUSDT, func(response SpotWebsocketV1PublicV1TradeResponse) error {
		assert.Equal(t, respBody, response)
		return nil
	})
	require.NoError(t, err)

	assert.NoError(t, svc.Run())
	assert.NoError(t, unsubscribe())
	assert.NoError(t, svc.Ping())
	assert.NoError(t, svc.Close())
}
