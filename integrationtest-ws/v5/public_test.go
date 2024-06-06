//go:build integrationtestwsv5

package integrationtestwsv5

import (
	"testing"

	"github.com/dimkus/bybit/v2"
	"github.com/dimkus/bybit/v2/integrationtest-ws/testhelper"
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

func TestV5Public_Kline(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.V5().Public(bybit.CategoryV5Linear)
	require.NoError(t, err)

	_, err = svc.SubscribeKline(
		bybit.V5WebsocketPublicKlineParamKey{
			Interval: bybit.Interval5,
			Symbol:   bybit.SymbolV5BTCUSDT,
		},
		func(response bybit.V5WebsocketPublicKlineResponse) error {
			goldenFilename := "./testdata/public-v5-kline.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
			return nil
		},
	)
	require.NoError(t, err)

	require.NoError(t, svc.Run())
}

func TestV5Public_Ticker(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	{
		svc, err := wsClient.V5().Public(bybit.CategoryV5Linear)
		require.NoError(t, err)

		_, err = svc.SubscribeTicker(
			bybit.V5WebsocketPublicTickerParamKey{
				Symbol: bybit.SymbolV5BTCUSDT,
			},
			func(response bybit.V5WebsocketPublicTickerResponse) error {
				goldenFilename := "./testdata/public-v5-ticker-inverse.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
				return nil
			},
		)
		require.NoError(t, err)
		require.NoError(t, svc.Run())
	}
	{
		svc, err := wsClient.V5().Public(bybit.CategoryV5Option)
		require.NoError(t, err)

		_, err = svc.SubscribeTicker(
			bybit.V5WebsocketPublicTickerParamKey{
				Symbol: bybit.SymbolV5BTCUSDT,
			},
			func(response bybit.V5WebsocketPublicTickerResponse) error {
				goldenFilename := "./testdata/public-v5-ticker-option.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
				return nil
			},
		)
		require.NoError(t, err)
		require.NoError(t, svc.Run())
	}
	{
		svc, err := wsClient.V5().Public(bybit.CategoryV5Spot)
		require.NoError(t, err)

		_, err = svc.SubscribeTicker(
			bybit.V5WebsocketPublicTickerParamKey{
				Symbol: bybit.SymbolV5BTCUSDT,
			},
			func(response bybit.V5WebsocketPublicTickerResponse) error {
				goldenFilename := "./testdata/public-v5-ticker-spot.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
				return nil
			},
		)
		require.NoError(t, err)
		require.NoError(t, svc.Run())
	}
}
