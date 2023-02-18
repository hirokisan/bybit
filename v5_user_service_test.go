package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/require"
)

func TestV5User_GetAPIKey(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := "/v5/user/query-api"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"id":       "16651472",
				"note":     "testxxx",
				"apiKey":   "opjSlSOzqIeXROT5rq",
				"readOnly": 0,
				"secret":   "",
				"permissions": map[string]interface{}{
					"ContractTrade": []interface{}{},
					"Spot":          []interface{}{},
					"Wallet": []interface{}{
						"AccountTransfer",
					},
					"Options":     []interface{}{},
					"Derivatives": []interface{}{},
					"CopyTrading": []interface{}{},
					"BlockTrade":  []interface{}{},
					"Exchange":    []interface{}{},
					"NFT":         []interface{}{},
				},
				"ips":           []interface{}{"*"},
				"type":          1,
				"deadlineDay":   88,
				"expiredAt":     "2023-05-15T03:13:24Z",
				"createdAt":     "2023-02-15T03:13:24Z",
				"unified":       0,
				"uta":           0,
				"userID":        53888001,
				"inviterID":     0,
				"vipLevel":      "No VIP",
				"mktMakerLevel": "0",
				"affiliateID":   0,
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

		resp, err := client.V5().User().GetAPIKey()
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}
