//go:build integrationtestwsspotv1

package integrationtestwsspotv1

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest-ws/testhelper"
	"github.com/stretchr/testify/require"
)

func TestPublicV2Trade(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.Spot().V1().PublicV2()
	require.NoError(t, err)

	_, err = svc.SubscribeTrade(bybit.SymbolSpotBTCUSDT, func(response bybit.SpotWebsocketV1PublicV2TradeResponse) error {
		goldenFilename := "./testdata/public-v2-trade.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
		return nil
	})
	require.NoError(t, err)

	require.NoError(t, svc.Run()) // ignore first message
	require.NoError(t, svc.Run())
}
