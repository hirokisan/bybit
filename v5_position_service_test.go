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
						"symbol":                 "BTCUSDT",
						"leverage":               "1",
						"avgPrice":               "0",
						"liqPrice":               "",
						"riskLimitValue":         "6000000",
						"takeProfit":             "",
						"positionValue":          "",
						"tpslMode":               "Full",
						"riskId":                 3,
						"trailingStop":           "0",
						"unrealisedPnl":          "",
						"markPrice":              "43577.97",
						"cumRealisedPnl":         "-1618.71655122",
						"adlRankIndicator":       0,
						"isReduceOnly":           false,
						"mmrSysUpdatedTime":      "",
						"leverageSysUpdatedTime": "",
						"positionMM":             "0",
						"positionBalance":        "0",
						"createdTime":            "1677287853003",
						"updatedTime":            "1691735404177",
						"seq":                    8093271558,
						"positionIdx":            0,
						"positionIM":             "0",
						"side":                   "",
						"bustPrice":              "",
						"size":                   "0",
						"positionStatus":         "Normal",
						"stopLoss":               "",
						"tradeMode":              0,
						"autoAddMargin":          0,
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

func TestV5Position_GetClosedPnL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetClosedPnLParam{
			Category: CategoryV5Linear,
		}

		path := "/v5/position/closed-pnl"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "bb9020dc-92e2-4216-becf-cf4be4dcc81a%3A1680525485073724079%2Cbb9020dc-92e2-4216-becf-cf4be4dcc81a%3A1680525485073724079",
				"list": []map[string]interface{}{
					{
						"symbol":        "BTCUSDT",
						"orderId":       "bb9020dc-92e2-4216-becf-cf4be4dcc81a",
						"side":          "Sell",
						"qty":           "0.01",
						"orderPrice":    "26736.7",
						"orderType":     "Market",
						"execType":      "Trade",
						"closedSize":    "0.01",
						"cumEntryValue": "281.447",
						"avgEntryPrice": "28144.7",
						"cumExitValue":  "281.499",
						"avgExitPrice":  "28149.9",
						"closedPnl":     "-0.2857676",
						"fillCount":     "1",
						"leverage":      "1",
						"createdTime":   "1680525485073",
						"updatedTime":   "1680525485078",
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

		resp, err := client.V5().Position().GetClosedPnL(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetClosedPnLParam{
			Category: CategoryV5Linear,
		}

		path := "/v5/position/closed-pnl"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"nextPageCursor": "bb9020dc-92e2-4216-becf-cf4be4dcc81a%3A1680525485073724079%2Cbb9020dc-92e2-4216-becf-cf4be4dcc81a%3A1680525485073724079",
				"list": []map[string]interface{}{
					{
						"symbol":        "BTCUSDT",
						"orderId":       "bb9020dc-92e2-4216-becf-cf4be4dcc81a",
						"side":          "Sell",
						"qty":           "0.01",
						"orderPrice":    "26736.7",
						"orderType":     "Market",
						"execType":      "Trade",
						"closedSize":    "0.01",
						"cumEntryValue": "281.447",
						"avgEntryPrice": "28144.7",
						"cumExitValue":  "281.499",
						"avgExitPrice":  "28149.9",
						"closedPnl":     "-0.2857676",
						"fillCount":     "1",
						"leverage":      "1",
						"createdTime":   "1680525485073",
						"updatedTime":   "1680525485078",
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

		_, err = client.V5().Position().GetClosedPnL(param)
		assert.Error(t, err)
	})
}

func TestV5Position_SwitchPositionMarginMode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5SwitchPositionMarginModeParam{
			Category:     CategoryV5Linear,
			TradeMode:    PositionMarginIsolated,
			Symbol:       SymbolV5BTCUSDT,
			BuyLeverage:  "1",
			SellLeverage: "1",
		}

		path := "/v5/position/switch-isolated"
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

		resp, err := client.V5().Position().SwitchPositionMarginMode(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5SwitchPositionMarginModeParam{
			Category:     CategoryV5Linear,
			TradeMode:    PositionMarginIsolated,
			Symbol:       SymbolV5BTCUSDT,
			BuyLeverage:  "1",
			SellLeverage: "1",
		}

		path := "/v5/position/switch-isolated"
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

		_, err = client.V5().Position().SwitchPositionMarginMode(param)
		assert.Error(t, err)
	})
}

func TestV5Position_SetRiskLimit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5SetRiskLimitParam{
			Category: CategoryV5Linear,
			Symbol:   SymbolV5BTCUSDT,
			RiskID:   3,
		}

		path := "/v5/position/set-risk-limit"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"riskId":         3,
				"riskLimitValue": "6000000",
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

		resp, err := client.V5().Position().SetRiskLimit(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5SetRiskLimitParam{
			Category: CategoryV5Linear,
			Symbol:   SymbolV5BTCUSDT,
			RiskID:   3,
		}

		path := "/v5/position/set-risk-limit"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"category":       "linear",
				"riskId":         3,
				"riskLimitValue": "6000000",
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

		_, err = client.V5().Position().SetRiskLimit(param)
		assert.Error(t, err)
	})
}
