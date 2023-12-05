package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5Account_GetWalletBalance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := "/v5/account/wallet-balance"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"list": []map[string]interface{}{
					{
						"totalEquity":            "18070.32797922",
						"accountIMRate":          "0.0101",
						"totalMarginBalance":     "18070.32797922",
						"totalInitialMargin":     "182.60183684",
						"accountType":            "UNIFIED",
						"totalAvailableBalance":  "17887.72614237",
						"accountMMRate":          "0",
						"totalPerpUPL":           "-0.11001349",
						"totalWalletBalance":     "18070.43799271",
						"totalMaintenanceMargin": "0.38106773",
						"coin": []map[string]interface{}{
							{
								"availableToBorrow":   "2.5",
								"accruedInterest":     "0",
								"availableToWithdraw": "0.805994",
								"totalOrderIM":        "0",
								"equity":              "0.805994",
								"totalPositionMM":     "0",
								"usdValue":            "12920.95352538",
								"unrealisedPnl":       "0",
								"borrowAmount":        "0",
								"totalPositionIM":     "0",
								"walletBalance":       "0.805994",
								"free":                "",
								"locked":              "",
								"cumRealisedPnl":      "0",
								"coin":                "BTC",
							},
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
			WithBaseURL(server.URL).
			WithAuth("test", "test")

		resp, err := client.V5().Account().GetWalletBalance(AccountTypeUnified, nil)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}

func TestV5Account_GetCollateralInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		currency := "BTC"
		param := V5GetCollateralInfoParam{Currency: &currency}

		path := "/v5/account/collateral-info"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"list": []map[string]interface{}{
					{
						"availableToBorrow":   "3",
						"freeBorrowingAmount": "",
						"freeBorrowAmount":    "0",
						"maxBorrowingAmount":  "3",
						"hourlyBorrowRate":    "0.00000147",
						"borrowUsageRate":     "0",
						"collateralSwitch":    true,
						"borrowAmount":        "0",
						"borrowable":          true,
						"currency":            "BTC",
						"marginCollateral":    true,
						"freeBorrowingLimit":  "0",
						"collateralRatio":     "0.95",
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

		resp, err := client.V5().Account().GetCollateralInfo(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}

func TestV5Account_GetAccountInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := "/v5/account/info"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"marginMode":          "REGULAR_MARGIN",
				"updatedTime":         "1672106576000",
				"unifiedMarginStatus": 3,
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

		resp, err := client.V5().Account().GetAccountInfo()
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}

func TestV5Account_GetTransactionLog(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetTransactionLogParam{}

		path := "/v5/account/transaction-log"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"nextPageCursor": "133%3A1%2C133%3A1",
				"list": []map[string]interface{}{
					{
						"symbol":          "BTCUSDT",
						"category":        "linear",
						"side":            "Sell",
						"transactionTime": "1680525485078",
						"type":            "TRADE",
						"qty":             "0.01",
						"size":            "0",
						"currency":        "USDT",
						"tradePrice":      "28149.9",
						"funding":         "",
						"fee":             "0.16889940",
						"cashFlow":        "0.052",
						"change":          "-0.1168994",
						"cashBalance":     "1149.11399896",
						"feeRate":         "0.0006",
						"bonusChange":     "0",
						"tradeId":         "259f8703-26ff-5e31-b9c6-edf3f2869f9a",
						"orderId":         "bb9020dc-92e2-4216-becf-cf4be4dcc81a",
						"orderLinkId":     "",
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

		resp, err := client.V5().Account().GetTransactionLog(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetTransactionLogParam{}

		path := "/v5/account/transaction-log"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"nextPageCursor": "133%3A1%2C133%3A1",
				"list": []map[string]interface{}{
					{
						"symbol":          "BTCUSDT",
						"category":        "linear",
						"side":            "Sell",
						"transactionTime": "1680525485078",
						"type":            "TRADE",
						"qty":             "0.01",
						"size":            "0",
						"currency":        "USDT",
						"tradePrice":      "28149.9",
						"funding":         "",
						"fee":             "0.16889940",
						"cashFlow":        "0.052",
						"change":          "-0.1168994",
						"cashBalance":     "1149.11399896",
						"feeRate":         "0.0006",
						"bonusChange":     "0",
						"tradeId":         "259f8703-26ff-5e31-b9c6-edf3f2869f9a",
						"orderId":         "bb9020dc-92e2-4216-becf-cf4be4dcc81a",
						"orderLinkId":     "",
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

		_, err = client.V5().Account().GetTransactionLog(param)
		assert.Error(t, err)
	})
}
