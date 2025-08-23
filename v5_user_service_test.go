package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/require"
)

func TestV5User_CreateSubUID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := "/v5/user/create-sub-member"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"uid":        "53888000",
				"username":   "test0001",
				"memberType": 1,
				"status":     1,
				"remark":     "test account",
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

		password := "Password123"
		switchVal := 1
		note := "test account"
		param := V5CreateSubUIDParam{
			Username:   "test0001",
			MemberType: 1,
			Password:   &password,
			Switch:     &switchVal,
			Note:       &note,
		}

		resp, err := client.V5().User().CreateSubUID(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}

func TestV5User_GetSubUIDList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := "/v5/user/query-sub-members"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"subMembers": []map[string]interface{}{
					{
						"uid":         "53888000",
						"username":    "test0001",
						"memberType":  1,
						"status":      1,
						"accountMode": 3,
						"remark":      "test account",
					},
					{
						"uid":         "53888001",
						"username":    "test0002",
						"memberType":  1,
						"status":      1,
						"accountMode": 3,
						"remark":      "",
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

		resp, err := client.V5().User().GetSubUIDList()
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}

func TestV5User_CreateSubUIDAPIKey(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := "/v5/user/create-sub-api"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"id":       "16651283",
				"note":     "test api key",
				"apiKey":   "XXXXXXXXXX",
				"readOnly": 1,
				"secret":   "YYYYYYYYYY",
				"permissions": map[string]interface{}{
					"ContractTrade": []interface{}{"Order", "Position"},
					"Spot":          []interface{}{"SpotTrade"},
					"Wallet":        []interface{}{"AccountTransfer"},
					"Options":       []interface{}{},
					"Derivatives":   []interface{}{},
					"CopyTrading":   []interface{}{},
					"BlockTrade":    []interface{}{},
					"Exchange":      []interface{}{},
					"NFT":           []interface{}{},
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

		note := "test api key"
		ips := "192.168.0.1,192.168.0.2"
		param := V5CreateSubUIDAPIKeyParam{
			Subuid:   53888000,
			ReadOnly: 1,
			Note:     &note,
			Ips:      &ips,
			Permissions: V5APIKeyPermissionsParam{
				ContractTrade: []string{"Order", "Position"},
				Spot:          []string{"SpotTrade"},
				Wallet:        []string{"AccountTransfer"},
			},
		}

		resp, err := client.V5().User().CreateSubUIDAPIKey(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
}

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
