package bybit

import (
	"encoding/json"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebsocketV5Public_Kline(t *testing.T) {
	respBody := map[string]interface{}{
		"topic": "kline.5.BTCUSDT",
		"type":  "snapshot",
		"ts":    1672324988882,
		"data": []map[string]interface{}{
			{
				"start":     1672324800000,
				"end":       1672325099999,
				"interval":  "5",
				"open":      "16649.5",
				"close":     "16677",
				"high":      "16677",
				"low":       "16608",
				"volume":    "2.081",
				"turnover":  "34666.4005",
				"confirm":   false,
				"timestamp": 1672324988882,
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
		_, err := svc.SubscribeKline(
			V5WebsocketPublicKlineParamKey{
				Interval: Interval5,
				Symbol:   SymbolV5BTCUSDT,
			},
			func(response V5WebsocketPublicKlineResponse) error {
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
