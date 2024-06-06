package bybit

import (
	"encoding/json"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebsocketV5Public_Liquidation(t *testing.T) {
	respBody := map[string]interface{}{
		"topic": "liquidation.BTCUSDT",
		"type":  "snapshot",
		"ts":    1673251091822,
		"data": []map[string]interface{}{
			{
				"price":       "25844.48",
				"side":        "Buy",
				"size":        "2.8",
				"symbol":      "BTCUSDT",
				"updatedTime": 1673251091822,
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
		_, err := svc.SubscribeLiquidation(
			V5WebsocketPublicLiquidationParamKey{
				Symbol: SymbolV5BTCUSDT,
			},
			func(response V5WebsocketPublicLiquidationResponse) error {
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
