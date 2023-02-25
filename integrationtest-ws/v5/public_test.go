//go:build integrationtestwsv5

package integrationtestwsv5

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest-ws/testhelper"
	"github.com/stretchr/testify/require"
)

func TestV5Public_OrderBook(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.V5().Public(bybit.CategoryV5Linear)
	require.NoError(t, err)

	_, err = svc.SubscribeOrderBook(
		bybit.V5WebsocketPublicOrderBookParamKey{
			Depth:  1,
			Symbol: bybit.SymbolV5BTCUSDT,
		},
		func(response bybit.V5WebsocketPublicOrderBookResponse) error {
			goldenFilename := "./testdata/public-v5-orderbook.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
			return nil
		},
	)
	require.NoError(t, err)

	require.NoError(t, svc.Run())
}
