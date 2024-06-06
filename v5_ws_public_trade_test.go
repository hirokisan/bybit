package bybit

import (
	"encoding/json"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebsocketV5Public_Trade(t *testing.T) {
	respBody := map[string]interface{}{
		"topic": "publicTrade.BTCUSDT",
		"type":  "snapshot",
		"ts":    1672304486868,
		"data": []map[string]interface{}{
			{
				"T":  1672304486865,
				"s":  "BTCUSDT",
				"S":  "Buy",
				"v":  "0.001",
				"p":  "16578.50",
				"L":  "PlusTick",
				"i":  "20f43950-d8dd-5b31-9112-a178eb6023af",
				"BT": false,
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	category := CategoryV5Linear

	server, teardown := testhelper.NewWebsocketServer(
		testhelper.WithWebsocketHandlerOption(V5WebsocketPublicPathFor(category), bytesBody),
	)
	defer teardown()

	wsClient := NewTestWebsocketClient().
		WithBaseURL(server.URL)

	svc, err := wsClient.V5().Public(category)
	require.NoError(t, err)

	{
		_, err := svc.SubscribeTrade(
			V5WebsocketPublicTradeParamKey{
				Symbol: SymbolV5BTCUSDT,
			},
			func(response V5WebsocketPublicTradeResponse) error {
				assert.Equal(t, respBody["topic"], response.Topic)
				return nil
			},
		)
		require.NoError(t, err)
	}

	assert.NoError(t, svc.Run())
	assert.NoError(t, svc.Ping())
	assert.NoError(t, svc.Close())
}
