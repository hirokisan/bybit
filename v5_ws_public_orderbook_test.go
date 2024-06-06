package bybit

import (
	"encoding/json"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebsocketV5Public_OrderBook(t *testing.T) {
	respBody := map[string]interface{}{
		"topic": "orderbook.1.BTCUSDT",
		"type":  "snapshot",
		"ts":    1677322353682,
		"data": map[string]interface{}{
			"s": "BTCUSDT",
			"b": [][]string{
				{
					"22975.10",
					"261.537",
				},
			},
			"a": [][]string{
				{
					"22975.40",
					"131.388",
				},
			},
			"u":   642570,
			"seq": 7995099758,
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
		_, err := svc.SubscribeOrderBook(
			V5WebsocketPublicOrderBookParamKey{
				Depth:  1,
				Symbol: SymbolV5BTCUSDT,
			},
			func(response V5WebsocketPublicOrderBookResponse) error {
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
