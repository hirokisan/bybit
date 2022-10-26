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
