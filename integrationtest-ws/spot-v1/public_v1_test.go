//go:build integrationtestwsspotv1

package integrationtestwsspotv1

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest-ws/testhelper"
	"github.com/stretchr/testify/require"
)

func TestPublicV1Trade(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.Spot().V1().PublicV1()
	require.NoError(t, err)

	_, err = svc.SubscribeTrade(bybit.SymbolSpotBTCUSDT, func(response bybit.SpotWebsocketV1PublicV1TradeResponse) error {
		goldenFilename := "./testdata/public-v1-trade.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
		return nil
	})
	require.NoError(t, err)

	require.NoError(t, svc.Run())
}
