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
