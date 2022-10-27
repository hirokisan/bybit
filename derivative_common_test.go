package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
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
