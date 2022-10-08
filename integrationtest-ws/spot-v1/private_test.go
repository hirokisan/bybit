//go:build integrationtestwsspotv1

package integrationtestwsspotv1

import (
	"errors"
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest-ws/testhelper"
	"github.com/stretchr/testify/require"
)

func TestPrivateOutboundAccountInfo(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.Spot().V1().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	svc.RegisterFuncOutboundAccountInfo(func(response bybit.SpotWebsocketV1PrivateOutboundAccountInfoResponse) error {
		goldenFilename := "./testdata/private-outbound-account-info.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
		return errors.New("finish")
	})

	require.NoError(t, svc.Run())

	go func() {
		for {
			if err := svc.Run(); err != nil {
				return
			}
		}
	}()

	client := bybit.NewTestClient().WithAuthFromEnv()
	_, err = client.Spot().V1().SpotPostOrder(bybit.SpotPostOrderParam{
		Symbol: bybit.SymbolSpotBTCUSDT,
		Qty:    10,
		Side:   bybit.SideBuy,
		Type:   bybit.OrderTypeSpotMarket,
	})
	require.NoError(t, err)
}
