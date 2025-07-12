package bybit

import (
	"encoding/json"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebsocketV5Public_AllLiquidation(t *testing.T) {
	data := []map[string]interface{}{
		{
			"p": "25844.48",
			"S": "Buy",
			"v": "2.8",
			"s": "BTCUSDT",
			"T": 1673251091822,
		},
	}
	respBody := map[string]interface{}{
		"topic": "allLiquidation.BTCUSDT",
		"type":  "snapshot",
		"ts":    1673251091822,
		"data":  data,
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
		_, err := svc.SubscribeAllLiquidation(
			V5WebsocketPublicAllLiquidationParamKey{
				Symbol: SymbolV5BTCUSDT,
			},
			func(response V5WebsocketPublicAllLiquidationResponse) error {
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
