package bybit

import (
	"encoding/json"
	"testing"

	"github.com/hirokisan/bybit/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpotWebsocketV1PrivateOutboundAccountInfo(t *testing.T) {
	respBody := []SpotWebsocketV1PrivateOutboundAccountInfoResponse{
		{
			Content: SpotWebsocketV1PrivateOutboundAccountInfoResponseContent{
				EventType:     SpotWebsocketV1PrivateEventTypeOutboundAccountInfo,
				Timestamp:     "1664285837492",
				AllowTrade:    true,
				AllowWithdraw: true,
				AllowWDeposit: true,
				WalletBalanceChanges: []SpotWebsocketV1PrivateOutboundAccountInfoResponseWalletBalanceChange{
					{
						SymbolName:       "USDT",
						AvailableBalance: "250.117543",
						ReservedBalance:  "10",
					},
				},
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewWebsocketServer(
		testhelper.WithWebsocketHandlerOption(SpotWebsocketV1PrivatePath, bytesBody),
	)
	defer teardown()

	wsClient := NewTestWebsocketClient().
		WithBaseURL(server.URL).
		WithAuth("test", "test")

	svc, err := wsClient.Spot().V1().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	require.NoError(t, svc.RegisterFuncOutboundAccountInfo(func(response SpotWebsocketV1PrivateOutboundAccountInfoResponse) error {
		assert.Equal(t, respBody[0], response)
		return nil
	}))

	assert.NoError(t, svc.Run())
	assert.NoError(t, svc.Ping())
	assert.NoError(t, svc.Close())
}
