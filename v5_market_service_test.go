package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5MarketGetKline(t *testing.T) {
	param := V5GetKlineParam{
		Category: CategoryV5Spot,
		Symbol:   SymbolV5BTCUSDT,
		Interval: IntervalD,
	}

	path := "/v5/market/kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": [][]string{
				{
					"1659398400000",
					"21999",
					"21999",
					"18000",
					"19176.24",
					"240.638551",
					"240.638551",
				},
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewServer(
		testhelper.WithHandlerOption(path, method, status, bytesBody),
	)
	defer teardown()

	client := NewTestClient().
		WithBaseURL(server.URL)

	resp, err := client.V5().Market().GetKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["symbol"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][0], resp.Result.List[0].StartTime)
}

func TestV5Market_GetMarkPriceKline(t *testing.T) {
	param := V5GetMarkPriceKlineParam{
		Category: CategoryV5Linear,
		Symbol:   SymbolV5BTCUSDT,
		Interval: IntervalD,
	}

	path := "/v5/market/mark-price-kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": [][]string{
				{
					"1676851200000",
					"24279.66",
					"24978",
					"23851.48",
					"24881.63",
				},
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewServer(
		testhelper.WithHandlerOption(path, method, status, bytesBody),
	)
	defer teardown()

	client := NewTestClient().
		WithBaseURL(server.URL)

	resp, err := client.V5().Market().GetMarkPriceKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["symbol"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][0], resp.Result.List[0].StartTime)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][1], resp.Result.List[0].Open)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][2], resp.Result.List[0].High)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][3], resp.Result.List[0].Low)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][4], resp.Result.List[0].Close)
}

func TestV5Market_GetIndexPriceKline(t *testing.T) {
	param := V5GetIndexPriceKlineParam{
		Category: CategoryV5Linear,
		Symbol:   SymbolV5BTCUSDT,
		Interval: IntervalD,
	}

	path := "/v5/market/index-price-kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": [][]string{
				{
					"1659830400000",
					"22952.39",
					"23399.69",
					"22847.34",
					"23175.46",
				},
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewServer(
		testhelper.WithHandlerOption(path, method, status, bytesBody),
	)
	defer teardown()

	client := NewTestClient().
		WithBaseURL(server.URL)

	resp, err := client.V5().Market().GetIndexPriceKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["symbol"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][0], resp.Result.List[0].StartTime)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][1], resp.Result.List[0].Open)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][2], resp.Result.List[0].High)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][3], resp.Result.List[0].Low)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][4], resp.Result.List[0].Close)
}

func TestV5Market_GetPremiumIndexPriceKline(t *testing.T) {
	param := V5GetPremiumIndexPriceKlineParam{
		Category: CategoryV5Linear,
		Symbol:   SymbolV5BTCUSDT,
		Interval: IntervalD,
	}

	path := "/v5/market/premium-index-price-kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": [][]string{
				{
					"1676246400000",
					"0.000074",
					"0.000508",
					"0.000121",
					"0.000508",
				},
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewServer(
		testhelper.WithHandlerOption(path, method, status, bytesBody),
	)
	defer teardown()

	client := NewTestClient().
		WithBaseURL(server.URL)

	resp, err := client.V5().Market().GetPremiumIndexPriceKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["symbol"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][0], resp.Result.List[0].StartTime)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][1], resp.Result.List[0].Open)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][2], resp.Result.List[0].High)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][3], resp.Result.List[0].Low)
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][4], resp.Result.List[0].Close)
}

func TestV5Market_GetInstrumentsInfo(t *testing.T) {
	t.Run("linear", func(t *testing.T) {
		param := V5GetInstrumentsInfoParam{
			Category: CategoryV5Linear,
		}

		path := "/v5/market/instruments-info"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "",
				"list": []map[string]interface{}{
					{
						"symbol":          "10000NFTUSDT",
						"contractType":    "LinearPerpetual",
						"status":          "Trading",
						"baseCoin":        "10000NFT",
						"quoteCoin":       "USDT",
						"settleCoin":      "USDT",
						"launchTime":      "1643007175000",
						"deliveryTime":    "0",
						"deliveryFeeRate": "",
						"priceScale":      "6",
						"leverageFilter": map[string]interface{}{
							"minLeverage":  "1",
							"maxLeverage":  "12.50",
							"leverageStep": "0.01",
						},
						"priceFilter": map[string]interface{}{
							"minPrice": "0.000005",
							"maxPrice": "9.999990",
							"tickSize": "0.000005",
						},
						"lotSizeFilter": map[string]interface{}{
							"maxOrderQty":         "370000",
							"minOrderQty":         "10",
							"qtyStep":             "10",
							"postOnlyMaxOrderQty": "3700000",
						},
						"unifiedMarginTrade": true,
						"fundingInterval":    480,
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		resp, err := client.V5().Market().GetInstrumentsInfo(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result.LinearInverse)
	})
	t.Run("option", func(t *testing.T) {
		param := V5GetInstrumentsInfoParam{
			Category: CategoryV5Option,
		}

		path := "/v5/market/instruments-info"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "option",
				"nextPageCursor": "0%2C500",
				"list": []map[string]interface{}{
					{
						"symbol":          "BTC-30JUN23-100000-C",
						"optionsType":     "Call",
						"status":          "ONLINE",
						"baseCoin":        "BTC",
						"quoteCoin":       "USD",
						"settleCoin":      "USDC",
						"launchTime":      "1672905600000",
						"deliveryTime":    "1688112000000",
						"deliveryFeeRate": "0.00015",
						"priceFilter": map[string]interface{}{
							"minPrice": "5",
							"maxPrice": "10000000",
							"tickSize": "5",
						},
						"lotSizeFilter": map[string]interface{}{
							"maxOrderQty": "10000",
							"minOrderQty": "0.01",
							"qtyStep":     "0.01",
						},
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		resp, err := client.V5().Market().GetInstrumentsInfo(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result.Option)
	})
	t.Run("spot", func(t *testing.T) {
		param := V5GetInstrumentsInfoParam{
			Category: CategoryV5Spot,
		}

		path := "/v5/market/instruments-info"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category": "spot",
				"list": []map[string]interface{}{
					{
						"symbol":     "BTCUSDT",
						"baseCoin":   "BTC",
						"quoteCoin":  "USDT",
						"innovation": "0",
						"status":     "1",
						"lotSizeFilter": map[string]interface{}{
							"basePrecision":  "0.000001",
							"quotePrecision": "0.00000001",
							"maxOrderQty":    "63.01197227",
							"minOrderQty":    "0.00004",
							"minOrderAmt":    "1",
							"maxOrderAmt":    "100000",
						},
						"priceFilter": map[string]interface{}{
							"tickSize": "0.01",
						},
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		resp, err := client.V5().Market().GetInstrumentsInfo(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result.Spot)
	})
}

func TestV5Market_GetTickers(t *testing.T) {
	t.Run("linear", func(t *testing.T) {
		param := V5GetTickersParam{
			Category: CategoryV5Linear,
		}

		path := "/v5/market/tickers"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category": "linear",
				"list": []map[string]interface{}{
					{
						"symbol":                 "BTCUSD",
						"lastPrice":              "16597.00",
						"indexPrice":             "16598.54",
						"markPrice":              "16596.00",
						"prevPrice24h":           "16464.50",
						"price24hPcnt":           "0.008047",
						"highPrice24h":           "30912.50",
						"lowPrice24h":            "15700.00",
						"prevPrice1h":            "16595.50",
						"openInterest":           "373504107",
						"openInterestValue":      "22505.67",
						"turnover24h":            "2352.94950046",
						"volume24h":              "49337318",
						"fundingRate":            "-0.001034",
						"nextFundingTime":        "1672387200000",
						"predictedDeliveryPrice": "",
						"basisRate":              "",
						"deliveryFeeRate":        "",
						"deliveryTime":           "0",
						"ask1Size":               "1",
						"bid1Price":              "16596.00",
						"ask1Price":              "16597.50",
						"bid1Size":               "1",
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		resp, err := client.V5().Market().GetTickers(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result.LinearInverse)
	})
	t.Run("option", func(t *testing.T) {
		symbol := SymbolV5("BTC-30DEC22-18000-C")
		param := V5GetTickersParam{
			Category: CategoryV5Option,
			Symbol:   &symbol,
		}

		path := "/v5/market/tickers"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category": "option",
				"list": []map[string]interface{}{
					{
						"symbol":                 "BTC-30DEC22-18000-C",
						"bid1Price":              "0",
						"bid1Size":               "0",
						"bid1Iv":                 "0",
						"ask1Price":              "435",
						"ask1Size":               "0.66",
						"ask1Iv":                 "5",
						"lastPrice":              "435",
						"highPrice24h":           "435",
						"lowPrice24h":            "165",
						"markPrice":              "0.00000009",
						"indexPrice":             "16600.55",
						"markIv":                 "0.7567",
						"underlyingPrice":        "16590.42",
						"openInterest":           "6.3",
						"turnover24h":            "2482.73",
						"volume24h":              "0.15",
						"totalVolume":            "99",
						"totalTurnover":          "1967653",
						"delta":                  "0.00000001",
						"gamma":                  "0.00000001",
						"vega":                   "0.00000004",
						"theta":                  "-0.00000152",
						"predictedDeliveryPrice": "0",
						"change24h":              "86",
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		resp, err := client.V5().Market().GetTickers(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result.Option)
	})
	t.Run("spot", func(t *testing.T) {
		param := V5GetTickersParam{
			Category: CategoryV5Spot,
		}

		path := "/v5/market/tickers"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category": "spot",
				"list": []map[string]interface{}{
					{
						"symbol":        "BTCUSDT",
						"bid1Price":     "20517.96",
						"bid1Size":      "2",
						"ask1Price":     "20527.77",
						"ask1Size":      "1.862172",
						"lastPrice":     "20533.13",
						"prevPrice24h":  "20393.48",
						"price24hPcnt":  "0.0068",
						"highPrice24h":  "21128.12",
						"lowPrice24h":   "20318.89",
						"turnover24h":   "243765620.65899866",
						"volume24h":     "11801.27771",
						"usdIndexPrice": "20784.12009279",
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		resp, err := client.V5().Market().GetTickers(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result.Spot)
	})
}
