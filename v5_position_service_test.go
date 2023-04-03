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

func TestV5Position_SetLeverage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5SetLeverageParam{
			Category:     CategoryV5Linear,
			Symbol:       SymbolV5BTCUSDT,
			BuyLeverage:  "0",
			SellLeverage: "0",
		}

		path := "/v5/position/set-leverage"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": nil,
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

		resp, err := client.V5().Position().SetLeverage(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5SetLeverageParam{
			Category:     CategoryV5Linear,
			Symbol:       SymbolV5BTCUSDT,
			BuyLeverage:  "0",
			SellLeverage: "0",
		}

		path := "/v5/position/set-leverage"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": nil,
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		_, err = client.V5().Position().SetLeverage(param)
		assert.Error(t, err)
	})
}

func TestV5Position_SetTradingStop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		price := "40000"
		param := V5SetTradingStopParam{
			Category:    CategoryV5Linear,
			Symbol:      SymbolV5BTCUSDT,
			PositionIdx: PositionIdxOneWay,
			TakeProfit:  &price,
		}

		path := "/v5/position/trading-stop"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": nil,
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

		resp, err := client.V5().Position().SetTradingStop(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		price := "40000"
		param := V5SetTradingStopParam{
			Category:    CategoryV5Linear,
			Symbol:      SymbolV5BTCUSDT,
			PositionIdx: PositionIdxOneWay,
			TakeProfit:  &price,
		}

		path := "/v5/position/trading-stop"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": nil,
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		_, err = client.V5().Position().SetTradingStop(param)
		assert.Error(t, err)
	})
}

func TestV5Position_SetTpSlMode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5SetTpSlModeParam{
			Category: CategoryV5Linear,
			Symbol:   SymbolV5BTCUSDT,
			TpSlMode: TpSlModeFull,
		}

		path := "/v5/position/set-tpsl-mode"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"tpSlMode": "Full",
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

		resp, err := client.V5().Position().SetTpSlMode(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5SetTpSlModeParam{
			Category: CategoryV5Linear,
			Symbol:   SymbolV5BTCUSDT,
			TpSlMode: TpSlModeFull,
		}

		path := "/v5/position/set-tpsl-mode"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"tpSlMode": "Full",
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

		_, err = client.V5().Position().SetTpSlMode(param)
		assert.Error(t, err)
	})
}

func TestV5Position_SwitchPositionMode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		coin := CoinBTC
		param := V5SwitchPositionModeParam{
			Category: CategoryV5Inverse,
			Mode:     PositionModeBothSides,
			Coin:     &coin,
		}

		path := "/v5/position/switch-mode"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": nil,
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

		resp, err := client.V5().Position().SwitchPositionMode(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		coin := CoinBTC
		param := V5SwitchPositionModeParam{
			Category: CategoryV5Inverse,
			Mode:     PositionModeBothSides,
			Coin:     &coin,
		}

		path := "/v5/position/switch-mode"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": nil,
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		client := NewTestClient().
			WithBaseURL(server.URL)

		_, err = client.V5().Position().SwitchPositionMode(param)
		assert.Error(t, err)
	})
}
