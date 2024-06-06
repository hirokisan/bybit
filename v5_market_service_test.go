package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
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
					"1713744000000",
					"38515.14",
					"51400",
					"29496.53",
					"43890.18",
					"202399.805532",
					"8017306629.63402469",
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
							"maxMktOrderQty":      "100.000",
							"minNotionalValue":    "5",
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

func TestV5Market_GetOrderbook(t *testing.T) {
	param := V5GetOrderbookParam{
		Category: CategoryV5Spot,
		Symbol:   SymbolV5BTCUSDT,
	}

	path := "/v5/market/orderbook"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"s": "BTCUSDT",
			"a": [][]interface{}{
				{
					"16638.64",
					"0.008479",
				},
			},
			"b": [][]interface{}{
				{
					"16638.27",
					"0.305749",
				},
			},
			"ts": 1672765737733,
			"u":  5277055,
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

	resp, err := client.V5().Market().GetOrderbook(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["s"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["a"].([][]interface{})[0][0], string(resp.Result.Asks[0].Price))
	assert.Equal(t, respBody["result"].(map[string]interface{})["b"].([][]interface{})[0][0], string(resp.Result.Bids[0].Price))
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

func TestV5Market_GetFundingRateHistory(t *testing.T) {
	param := V5GetFundingRateHistoryParam{
		Category: CategoryV5Linear,
		Symbol:   SymbolV5BTCUSDT,
	}

	path := "/v5/market/funding/history"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"list": []map[string]interface{}{
				{
					"symbol":               "BTCUSDT",
					"fundingRate":          "0.0001",
					"fundingRateTimestamp": "1679702400000",
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

	resp, err := client.V5().Market().GetFundingRateHistory(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestV5Market_GetPublicTradingHistory(t *testing.T) {
	param := V5GetPublicTradingHistoryParam{
		Category: CategoryV5Linear,
		Symbol:   SymbolV5BTCUSDT,
	}

	path := "/v5/market/recent-trade"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"list": []map[string]interface{}{
				{
					"execId":       "ab5c8ff6-79ae-508c-9f38-b6540b4ede73",
					"symbol":       "BTCUSDT",
					"price":        "27858.10",
					"size":         "0.001",
					"side":         "Buy",
					"time":         "1679907414058",
					"isBlockTrade": false,
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

	resp, err := client.V5().Market().GetPublicTradingHistory(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestV5Market_GetOpenInterest(t *testing.T) {
	param := V5GetOpenInterestParam{
		Category:     CategoryV5Linear,
		Symbol:       SymbolV5BTCUSDT,
		IntervalTime: Period1h,
	}

	path := "/v5/market/open-interest"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": []map[string]interface{}{
				{
					"openInterest": "60928.34400000",
					"timestamp":    "1679907600000",
				},
			},
			"nextPageCursor": "lastid%3D48598748",
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

	resp, err := client.V5().Market().GetOpenInterest(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestV5Market_GetHistoricalVolatility(t *testing.T) {
	param := V5GetHistoricalVolatilityParam{
		Category: CategoryV5Option,
	}

	path := "/v5/market/historical-volatility"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"category": "option",
		"result": []map[string]interface{}{
			{
				"period": 7,
				"value":  "0.67736876",
				"time":   "1679976000000",
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

	resp, err := client.V5().Market().GetHistoricalVolatility(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["category"], string(resp.Result.Category))
	assert.Equal(t, respBody["result"].([]map[string]interface{})[0]["period"], resp.Result.List[0].Period)
}

func TestV5Market_GetInsurance(t *testing.T) {
	param := V5GetInsuranceParam{}

	path := "/v5/market/insurance"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"updatedTime": "1682467200000",
			"list": []map[string]interface{}{
				{
					"coin":    "USDT",
					"balance": "472565175.36148953",
					"value":   "472629325.5808998",
				},
				{
					"coin":    "USDC",
					"balance": "996049.97239122",
					"value":   "996044.9942383142",
				},
				{
					"coin":    "EOS",
					"balance": "767102.58626542",
					"value":   "833072.7720000001",
				},
				{
					"coin":    "BTC",
					"balance": "47533.89021999",
					"value":   "1418307241.21",
				},
				{
					"coin":    "ADA",
					"balance": "895412.56208743",
					"value":   "372849.5568",
				},
				{
					"coin":    "BIT",
					"balance": "101256.65607912",
					"value":   "51903.8256",
				},
				{
					"coin":    "SOL",
					"balance": "7760.59832427",
					"value":   "101888.8",
				},
				{
					"coin":    "LTC",
					"balance": "909951.79684841",
					"value":   "85062219.48",
				},
				{
					"coin":    "DOT",
					"balance": "1103850.50091492",
					"value":   "6832831.5",
				},
				{
					"coin":    "ETH",
					"balance": "981392.36669191",
					"value":   "1922772648.16",
				},
				{
					"coin":    "XRP",
					"balance": "212730.43859841",
					"value":   "102642.22499999999",
				},
				{
					"coin":    "MANA",
					"balance": "4996881.117703",
					"value":   "2875705.0155",
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

	resp, err := client.V5().Market().GetInsurance(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestV5Market_GetRiskLimit(t *testing.T) {
	param := V5GetRiskLimitParam{
		Category: CategoryV5Linear,
	}

	path := "/v5/market/risk-limit"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"list": []map[string]interface{}{
				{
					"id":                1,
					"symbol":            "BTCUSDT",
					"riskLimitValue":    "2000000",
					"maintenanceMargin": "0.005",
					"initialMargin":     "0.01",
					"isLowestRisk":      1,
					"maxLeverage":       "100.00",
				},
				{
					"id":                2,
					"symbol":            "BTCUSDT",
					"riskLimitValue":    "4000000",
					"maintenanceMargin": "0.01",
					"initialMargin":     "0.0175",
					"isLowestRisk":      0,
					"maxLeverage":       "57.14",
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

	resp, err := client.V5().Market().GetRiskLimit(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}
