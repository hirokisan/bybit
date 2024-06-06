package bybit

import (
	"encoding/json"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebsocketV5Public_Ticker(t *testing.T) {
	t.Run("linear", func(t *testing.T) {
		respBody := map[string]interface{}{
			"topic": "tickers.BTCUSDT",
			"type":  "snapshot",
			"data": map[string]interface{}{
				"symbol":            "BTCUSDT",
				"tickDirection":     "PlusTick",
				"price24hPcnt":      "0.017103",
				"lastPrice":         "17216.00",
				"prevPrice24h":      "16926.50",
				"highPrice24h":      "17281.50",
				"lowPrice24h":       "16915.00",
				"prevPrice1h":       "17238.00",
				"markPrice":         "17217.33",
				"indexPrice":        "17227.36",
				"openInterest":      "68744.761",
				"openInterestValue": "1183601235.91",
				"turnover24h":       "1570383121.943499",
				"volume24h":         "91705.276",
				"nextFundingTime":   "1673280000000",
				"fundingRate":       "-0.000212",
				"bid1Price":         "17215.50",
				"bid1Size":          "84.489",
				"ask1Price":         "17216.00",
				"ask1Size":          "83.020",
			},
			"cs": 24987956059,
			"ts": 1673272861686,
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
			_, err := svc.SubscribeTicker(
				V5WebsocketPublicTickerParamKey{
					Symbol: SymbolV5BTCUSDT,
				},
				func(response V5WebsocketPublicTickerResponse) error {
					assert.Equal(t, respBody["topic"], response.Topic)
					testhelper.Compare(t, respBody["data"], response.Data.LinearInverse)
					return nil
				},
			)
			require.NoError(t, err)
		}

		assert.NoError(t, svc.Run())
		assert.NoError(t, svc.Ping())
		assert.NoError(t, svc.Close())
	})

	t.Run("option", func(t *testing.T) {
		respBody := map[string]interface{}{
			"topic": "tickers.BTCUSDT",
			"type":  "snapshot",
			"data": map[string]interface{}{
				"symbol":                 "BTC-6JAN23-17500-C",
				"bidPrice":               "0",
				"bidSize":                "0",
				"bidIv":                  "0",
				"askPrice":               "10",
				"askSize":                "5.1",
				"askIv":                  "0.514",
				"lastPrice":              "10",
				"highPrice24h":           "25",
				"lowPrice24h":            "5",
				"markPrice":              "7.86976724",
				"indexPrice":             "16823.73",
				"markPriceIv":            "0.4896",
				"underlyingPrice":        "16815.1",
				"openInterest":           "49.85",
				"turnover24h":            "446802.8473",
				"volume24h":              "26.55",
				"totalVolume":            "86",
				"totalTurnover":          "1437431",
				"delta":                  "0.047831",
				"gamma":                  "0.00021453",
				"vega":                   "0.81351067",
				"theta":                  "-19.9115368",
				"predictedDeliveryPrice": "0",
				"change24h":              "-0.33333334",
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		category := CategoryV5Option

		server, teardown := testhelper.NewWebsocketServer(
			testhelper.WithWebsocketHandlerOption(V5WebsocketPublicPathFor(category), bytesBody),
		)
		defer teardown()

		wsClient := NewTestWebsocketClient().
			WithBaseURL(server.URL)

		svc, err := wsClient.V5().Public(category)
		require.NoError(t, err)

		{
			_, err := svc.SubscribeTicker(
				V5WebsocketPublicTickerParamKey{
					Symbol: SymbolV5BTCUSDT,
				},
				func(response V5WebsocketPublicTickerResponse) error {
					assert.Equal(t, respBody["topic"], response.Topic)
					testhelper.Compare(t, respBody["data"], response.Data.Option)
					return nil
				},
			)
			require.NoError(t, err)
		}

		assert.NoError(t, svc.Run())
		assert.NoError(t, svc.Ping())
		assert.NoError(t, svc.Close())
	})

	t.Run("spot", func(t *testing.T) {
		respBody := map[string]interface{}{
			"topic": "tickers.BTCUSDT",
			"ts":    1673853746003,
			"type":  "snapshot",
			"cs":    2588407389,
			"data": map[string]interface{}{
				"symbol":        "BTCUSDT",
				"lastPrice":     "21109.77",
				"highPrice24h":  "21426.99",
				"lowPrice24h":   "20575",
				"prevPrice24h":  "20704.93",
				"volume24h":     "6780.866843",
				"turnover24h":   "141946527.22907118",
				"price24hPcnt":  "0.0196",
				"usdIndexPrice": "21120.2400136",
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		category := CategoryV5Spot

		server, teardown := testhelper.NewWebsocketServer(
			testhelper.WithWebsocketHandlerOption(V5WebsocketPublicPathFor(category), bytesBody),
		)
		defer teardown()

		wsClient := NewTestWebsocketClient().
			WithBaseURL(server.URL)

		svc, err := wsClient.V5().Public(category)
		require.NoError(t, err)

		{
			_, err := svc.SubscribeTicker(
				V5WebsocketPublicTickerParamKey{
					Symbol: SymbolV5BTCUSDT,
				},
				func(response V5WebsocketPublicTickerResponse) error {
					assert.Equal(t, respBody["topic"], response.Topic)
					testhelper.Compare(t, respBody["data"], response.Data.Spot)
					return nil
				},
			)
			require.NoError(t, err)
		}

		assert.NoError(t, svc.Run())
		assert.NoError(t, svc.Ping())
		assert.NoError(t, svc.Close())
	})
}
