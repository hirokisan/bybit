package bybit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dimkus/bybit/v2/testhelper"
)

func TestV5WebsocketPrivate_Wallet(t *testing.T) {
	t.Run("for Unifield", func(t *testing.T) {
		respBody := map[string]interface{}{
			"topic":        "wallet",
			"id":           "75d86e42f18b23b9ad2c1f10eaffa8bb:18483ff242aca593:0:01",
			"creationTime": 1677226839837,
			"data": []map[string]interface{}{
				{
					"accountIMRate":          "0.016",
					"accountMMRate":          "0.003",
					"totalEquity":            "12837.78330098",
					"totalWalletBalance":     "12840.4045924",
					"totalMarginBalance":     "12837.78330188",
					"totalAvailableBalance":  "12632.05767702",
					"totalPerpUPL":           "-2.62129051",
					"totalInitialMargin":     "205.72562486",
					"totalMaintenanceMargin": "39.42876721",
					"accountType":            "UNIFIED",
					"coin": []map[string]interface{}{
						{
							"coin":                "USDC",
							"equity":              "200.62572554",
							"usdValue":            "200.62572554",
							"walletBalance":       "201.34882644",
							"availableToWithdraw": "0",
							"availableToBorrow":   "1500000",
							"borrowAmount":        "0",
							"accruedInterest":     "0",
							"totalOrderIM":        "0",
							"totalPositionIM":     "202.99874213",
							"totalPositionMM":     "39.14289747",
							"unrealisedPnl":       "74.2768991",
							"cumRealisedPnl":      "-209.1544627",
						},
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewWebsocketServer(
			testhelper.WithWebsocketHandlerOption(V5WebsocketPrivatePath, bytesBody),
		)
		defer teardown()

		wsClient := NewTestWebsocketClient().
			WithBaseURL(server.URL).
			WithAuth("test", "test")

		svc, err := wsClient.V5().Private(CategoryV5All)
		require.NoError(t, err)

		require.NoError(t, svc.Subscribe())

		{
			_, err := svc.SubscribeWallet(func(response V5WebsocketPrivateWalletResponse) error {
				testhelper.Compare(t, respBody, response)
				return nil
			})
			require.NoError(t, err)
		}

		assert.NoError(t, svc.Run())
		assert.NoError(t, svc.Ping())
		assert.NoError(t, svc.Close())
	})
}
