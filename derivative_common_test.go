package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/dimkus/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDerivativesOrderBook(t *testing.T) {
	param := DerivativesOrderBookParam{
		Symbol:   SymbolDerivativeBTCUSDT,
		Category: CategoryDerivativeLinear,
	}

	path := "/derivatives/v3/public/order-book/L2"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"s": "BTCUSDT",
			"b": [][]string{
				{
					"20621",
					"0.015",
				},
			},
			"a": [][]string{
				{
					"20621",
					"0.015",
				},
			},
			"ts": 1666776701825,
			"u":  325507,
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesOrderBook(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["s"], string(resp.Result.Symbol))
}

func TestDerivativesKline(t *testing.T) {
	param := DerivativesKlineParam{
		Symbol:   SymbolDerivativeBTCUSDT,
		Category: CategoryDerivativeLinear,
		Interval: IntervalD,
		Start:    1652112000000,
		End:      1652544000000,
	}

	path := "/derivatives/v3/public/kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": [][]string{
				{
					"1652486400000",
					"29284.5",
					"30409.5",
					"28649.5",
					"30115.5",
					"9975.291",
					"292816985.039",
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["symbol"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][0], resp.Result.Lists[0].Start)
}

func TestDerivativesTickers(t *testing.T) {
	param := DerivativesTickersParam{
		Category: CategoryDerivativeLinear,
	}

	path := "/derivatives/v3/public/tickers"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"list": []map[string]string{
				{
					"symbol":                 "ZRXUSDT",
					"bidPrice":               "0.252",
					"askPrice":               "0.2522",
					"lastPrice":              "0.2497",
					"lastTickDirection":      "MinusTick",
					"prevPrice24h":           "0.2638",
					"price24hPcnt":           "-0.053449",
					"highPrice24h":           "0.2645",
					"lowPrice24h":            "0.2497",
					"prevPrice1h":            "0.2530",
					"markPrice":              "0.2522",
					"indexPrice":             "0.2524",
					"openInterest":           "516350",
					"turnover24h":            "1439.724",
					"volume24h":              "5474",
					"fundingRate":            "0.0001",
					"nextFundingTime":        "1666972800000",
					"predictedDeliveryPrice": "",
					"basisRate":              "",
					"deliveryFeeRate":        "",
					"deliveryTime":           "0",
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesTickers(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestDerivativesTickersForOption(t *testing.T) {
	param := DerivativesTickersForOptionParam{
		Symbol: SymbolDerivativeBTC31MAR23_40000C,
	}

	path := "/derivatives/v3/public/tickers"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]string{
			"category":               "option",
			"symbol":                 "BTC-31MAR23-40000-C",
			"bidPrice":               "270",
			"bidSize":                "19.9",
			"bidIv":                  "0.6514",
			"askPrice":               "325",
			"askSize":                "45.9",
			"askIv":                  "0.6778",
			"lastPrice":              "225",
			"highPrice24h":           "0",
			"lowPrice24h":            "0",
			"markPrice":              "277.11257031",
			"indexPrice":             "20193.37",
			"markPriceIv":            "0.655",
			"underlyingPrice":        "20259.98",
			"openInterest":           "4.63",
			"turnover24h":            "0",
			"volume24h":              "0",
			"totalVolume":            "56",
			"totalTurnover":          "1086495",
			"delta":                  "0.08275356",
			"gamma":                  "0.0000177",
			"vega":                   "20.06494071",
			"theta":                  "-4.26877439",
			"predictedDeliveryPrice": "0",
			"change24h":              "0",
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesTickersForOption(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestDerivativesInstruments(t *testing.T) {
	param := DerivativesInstrumentsParam{
		Category: CategoryDerivativeLinear,
	}

	path := "/derivatives/v3/public/instruments-info"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"list": []map[string]interface{}{
				{
					"symbol":       "10000NFTUSDT",
					"contractType": "LinearPerpetual",
					"status":       "Trading",
					"baseCoin":     "10000NFT", "quoteCoin": "USDT",
					"launchTime":      "1643007175000",
					"deliveryTime":    "0",
					"deliveryFeeRate": "",
					"priceScale":      "6",
					"leverageFilter": map[string]interface{}{
						"minLeverage":  "1",
						"maxLeverage":  "12",
						"leverageStep": "0.01",
					},
					"priceFilter": map[string]interface{}{
						"minPrice": "0.000005",
						"maxPrice": "9.999990",
						"tickSize": "0.000005",
					},
					"lotSizeFilter": map[string]interface{}{
						"maxTradingQty": "370000",
						"minTradingQty": "10",
						"qtyStep":       "10",
					},
				},
			},
			"nextPageCursor": "",
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesInstruments(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestDerivativesInstrumentsForOption(t *testing.T) {
	param := DerivativesInstrumentsForOptionParam{}

	path := "/derivatives/v3/public/instruments-info"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"resultTotalSize": 9234,
			"cursor":          "0%2C500",
			"dataList": []map[string]interface{}{
				{
					"category":        "option",
					"symbol":          "BTC-31MAR23-300000-C",
					"status":          "ONLINE",
					"baseCoin":        "BTC",
					"quoteCoin":       "USD",
					"settleCoin":      "USDC",
					"optionsType":     "Call",
					"launchTime":      "1665043200000",
					"deliveryTime":    "1680249600000",
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesInstrumentsForOption(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	testhelper.Compare(t, respBody["result"], resp.Result)
}

func TestDerivativesMarkPriceKline(t *testing.T) {
	param := DerivativesMarkPriceKlineParam{
		Symbol:   SymbolDerivativeBTCUSDT,
		Category: CategoryDerivativeLinear,
		Interval: IntervalD,
		Start:    1652112000000,
		End:      1652544000000,
	}

	path := "/derivatives/v3/public/mark-price-kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": [][]string{
				{
					"1652486400000",
					"29291.4",
					"30327.13",
					"28641.42",
					"30088.12",
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesMarkPriceKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["symbol"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][0], resp.Result.List[0].Start)
}

func TestDerivativesIndexPriceKline(t *testing.T) {
	param := DerivativesIndexPriceKlineParam{
		Symbol:   SymbolDerivativeBTCUSDT,
		Category: CategoryDerivativeLinear,
		Interval: IntervalD,
		Start:    1652112000000,
		End:      1652544000000,
	}

	path := "/derivatives/v3/public/index-price-kline"
	method := http.MethodGet
	status := http.StatusOK
	respBody := map[string]interface{}{
		"result": map[string]interface{}{
			"category": "linear",
			"symbol":   "BTCUSDT",
			"list": [][]string{
				{
					"1652486400000",
					"29291.4",
					"30334.75",
					"28638.38",
					"30088.88",
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

	resp, err := client.Derivative().UnifiedMargin().DerivativesIndexPriceKline(param)
	require.NoError(t, err)

	require.NotNil(t, resp)
	assert.Equal(t, respBody["result"].(map[string]interface{})["symbol"], string(resp.Result.Symbol))
	assert.Equal(t, respBody["result"].(map[string]interface{})["list"].([][]string)[0][0], resp.Result.List[0].Start)
}
