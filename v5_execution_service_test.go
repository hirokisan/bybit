package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5ExecutionService_GetExecutionList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetExecutionParam{
			Category: CategoryV5Linear,
		}

		path := "/v5/execution/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "",
				"list": []map[string]interface{}{
					{
						"symbol":          "BTCUSDT",
						"orderId":         "orderId",
						"orderLinkId":     "orderLinkId",
						"side":            "Buy",
						"orderPrice":      "24617.44",
						"orderQty":        "0.01",
						"leavesQty":       "0.5",
						"orderType":       "Limit",
						"stopOrderType":   "tpslOrder",
						"execFee":         "0.01385448",
						"execId":          "execId",
						"execPrice":       "23089.30",
						"execQty":         "0.001",
						"execType":        "Trade",
						"execValue":       "34.5",
						"execTime":        "1693555728000",
						"feeCurrency":     "BTC",
						"isMaker":         true,
						"feeRate":         "0.03455",
						"tradeIv":         "0.0035",
						"markIv":          "0.0352",
						"markPrice":       "23239.10",
						"indexPrice":      "20334.55",
						"underlyingPrice": "21243.21",
						"blockTradeId":    "blockTradeId",
						"closedSize":      "233",
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
			WithBaseURL(server.URL).
			WithAuth("test", "test")

		resp, err := client.V5().Execution().GetExecutionList(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetExecutionParam{
			Category: CategoryV5Linear,
		}

		path := "/v5/execution/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "",
				"list": []map[string]interface{}{
					{
						"symbol":          "BTCUSDT",
						"orderId":         "orderId",
						"orderLinkId":     "orderLinkId",
						"side":            "Buy",
						"orderPrice":      "24617.44",
						"orderQty":        "0.01",
						"leavesQty":       "0.5",
						"orderType":       "Limit",
						"stopOrderType":   "tpslOrder",
						"execFee":         "0.01385448",
						"execId":          "execId",
						"execPrice":       "23089.30",
						"execQty":         "0.001",
						"execType":        "Trade",
						"execValue":       "34.5",
						"execTime":        "1693555728000",
						"feeCurrency":     "BTC",
						"isMaker":         true,
						"feeRate":         "0.03455",
						"tradeIv":         "0.0035",
						"markIv":          "0.0352",
						"markPrice":       "23239.10",
						"indexPrice":      "20334.55",
						"underlyingPrice": "21243.21",
						"blockTradeId":    "blockTradeId",
						"closedSize":      "233",
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

		_, err = client.V5().Execution().GetExecutionList(param)
		assert.Error(t, err)
	})
}
