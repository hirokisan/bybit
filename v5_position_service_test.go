package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5Position_GetPositionInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		symbol := SymbolV5BTCUSDT
		param := V5GetPositionInfoParam{
			Category: CategoryV5Linear,
			Symbol:   &symbol,
		}

		path := "/v5/position/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "",
				"list": []map[string]interface{}{
					{
						"symbol":         "BTCUSDT",
						"leverage":       "2",
						"avgPrice":       "22714",
						"liqPrice":       "0.10",
						"riskLimitValue": "2000000",
						"takeProfit":     "0.00",
						"positionValue":  "22.714",
						"tpslMode":       "Full",
						"riskId":         1,
						"trailingStop":   "0.00",
						"unrealisedPnl":  "1.90344",
						"markPrice":      "24617.44",
						"cumRealisedPnl": "-592.751693",
						"positionMM":     "0.0000005",
						"createdTime":    "1637760408825",
						"positionIdx":    1,
						"positionIM":     "0.22714",
						"updatedTime":    "1676678400081",
						"side":           "Buy",
						"bustPrice":      "0.10",
						"size":           "0.001",
						"positionStatus": "Normal",
						"stopLoss":       "0.00",
						"tradeMode":      0,
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

		resp, err := client.V5().Position().GetPositionInfo(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		symbol := SymbolV5BTCUSDT
		param := V5GetPositionInfoParam{
			Category: CategoryV5Linear,
			Symbol:   &symbol,
		}

		path := "/v5/position/list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "",
				"list": []map[string]interface{}{
					{
						"symbol":         "BTCUSDT",
						"leverage":       "2",
						"avgPrice":       "22714",
						"liqPrice":       "0.10",
						"riskLimitValue": "2000000",
						"takeProfit":     "0.00",
						"positionValue":  "22.714",
						"tpslMode":       "Full",
						"riskId":         1,
						"trailingStop":   "0.00",
						"unrealisedPnl":  "1.90344",
						"markPrice":      "24617.44",
						"cumRealisedPnl": "-592.751693",
						"positionMM":     "0.0000005",
						"createdTime":    "1637760408825",
						"positionIdx":    1,
						"positionIM":     "0.22714",
						"updatedTime":    "1676678400081",
						"side":           "Buy",
						"bustPrice":      "0.10",
						"size":           "0.001",
						"positionStatus": "Normal",
						"stopLoss":       "0.00",
						"tradeMode":      0,
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

		_, err = client.V5().Position().GetPositionInfo(param)
		assert.Error(t, err)
	})
}
