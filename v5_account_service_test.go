package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
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

		resp, err := client.V5().Account().GetWalletBalance(UnifiedAccount, "")
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}
