package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5Asset_CreateInternalTransfer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5CreateInternalTransferParam{
			TransferID:      "42c0cfb0-6bca-c242-bc76-4e6df6cbcb16",
			Coin:            CoinBTC,
			Amount:          "0.05",
			FromAccountType: AccountTypeV5UNIFIED,
			ToAccountType:   AccountTypeV5CONTRACT,
		}

		path := "/v5/asset/transfer/inter-transfer"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"transferId": "42c0cfb0-6bca-c242-bc76-4e6df6cbcb16",
				"status":     "SUCCESS",
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

		resp, err := client.V5().Asset().CreateInternalTransfer(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		fmt.Println(resp.Result, respBody["result"])
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5CreateInternalTransferParam{
			TransferID:      "42c0cfb0-6bca-c242-bc76-4e6df6cbcb16",
			Coin:            CoinBTC,
			Amount:          "0.05",
			FromAccountType: AccountTypeV5UNIFIED,
			ToAccountType:   AccountTypeV5CONTRACT,
		}

		path := "/v5/asset/transfer/inter-transfer"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"transferId": "42c0cfb0-6bca-c242-bc76-4e6df6cbcb16",
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

		_, err = client.V5().Asset().CreateInternalTransfer(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetInternalTransferRecords(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetInternalTransferRecordsParam{}

		path := "/v5/asset/transfer/query-inter-transfer-list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"list": []map[string]interface{}{
					{
						"transferId":      "selfTransfer_5ce5b8d9-8477-4bc6-91a4-9a98dad6dc65",
						"coin":            "BTC",
						"amount":          "0.1",
						"fromAccountType": "SPOT",
						"toAccountType":   "CONTRACT",
						"timestamp":       "1637939106000",
						"status":          "SUCCESS",
					},
				},
				"nextPageCursor": "eyJtaW5JRCI6MTYyMjgwLCJtYXhJRCI6MTYyMjgwfQ==",
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

		resp, err := client.V5().Asset().GetInternalTransferRecords(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetInternalTransferRecordsParam{}

		path := "/v5/asset/transfer/query-inter-transfer-list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"list": []map[string]interface{}{
					{
						"transferId":      "selfTransfer_5ce5b8d9-8477-4bc6-91a4-9a98dad6dc65",
						"coin":            "BTC",
						"amount":          "0.1",
						"fromAccountType": "SPOT",
						"toAccountType":   "CONTRACT",
						"timestamp":       "1637939106000",
						"status":          "SUCCESS",
					},
				},
				"nextPageCursor": "eyJtaW5JRCI6MTYyMjgwLCJtYXhJRCI6MTYyMjgwfQ==",
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

		_, err = client.V5().Asset().GetInternalTransferRecords(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_CreateUniversalTransfer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5CreateUniversalTransferParam{
			TransferID:      "be7a2462-1138-4e27-80b1-62653f24925e",
			Coin:            CoinETH,
			Amount:          "0.5",
			FromMemberID:    592334,
			ToMemberID:      691355,
			FromAccountType: AccountTypeV5CONTRACT,
			ToAccountType:   AccountTypeV5UNIFIED,
		}

		path := "/v5/asset/transfer/universal-transfer"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"transferId": "be7a2462-1138-4e27-80b1-62653f24925e",
				"status":     "SUCCESS",
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

		resp, err := client.V5().Asset().CreateUniversalTransfer(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		fmt.Println(resp.Result, respBody["result"])
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5CreateUniversalTransferParam{
			TransferID:      "be7a2462-1138-4e27-80b1-62653f24925e",
			Coin:            CoinETH,
			Amount:          "0.5",
			FromMemberID:    592334,
			ToMemberID:      691355,
			FromAccountType: AccountTypeV5CONTRACT,
			ToAccountType:   AccountTypeV5UNIFIED,
		}

		path := "/v5/asset/transfer/universal-transfer"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"transferId": "42c0cfb0-6bca-c242-bc76-4e6df6cbcb16",
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

		_, err = client.V5().Asset().CreateUniversalTransfer(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetUniversalTransferRecords(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetUniversalTransferRecordsParam{
			Limit:  testhelper.Ptr(1),
			Cursor: testhelper.Ptr("eyJtaW5JRCI6MTc5NjU3OCwibWF4SUQiOjE3OTY1Nzh9"),
		}

		path := "/v5/asset/transfer/query-universal-transfer-list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"list": []map[string]interface{}{
					{
						"transferId":      "universalTransfer_4c3cfe2f-85cb-11ed-ac09-9e37823c81cd_533285",
						"coin":            "USDC",
						"amount":          "1000",
						"timestamp":       "1672134373000",
						"status":          "SUCCESS",
						"fromAccountType": "UNIFIED",
						"toAccountType":   "UNIFIED",
						"fromMemberId":    "533285",
						"toMemberId":      "592324",
					},
				},
				"nextPageCursor": "eyJtaW5JRCI6MTc4OTYwNSwibWF4SUQiOjE3ODk2MDV9",
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

		resp, err := client.V5().Asset().GetUniversalTransferRecords(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetUniversalTransferRecordsParam{
			Limit:  testhelper.Ptr(1),
			Cursor: testhelper.Ptr("eyJtaW5JRCI6MTc5NjU3OCwibWF4SUQiOjE3OTY1Nzh9"),
		}

		path := "/v5/asset/transfer/query-universal-transfer-list"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"list": []map[string]interface{}{
					{
						"transferId":      "universalTransfer_4c3cfe2f-85cb-11ed-ac09-9e37823c81cd_533285",
						"coin":            "USDC",
						"amount":          "1000",
						"timestamp":       "1672134373000",
						"status":          "SUCCESS",
						"fromAccountType": "UNIFIED",
						"toAccountType":   "UNIFIED",
						"fromMemberId":    "533285",
						"toMemberId":      "592324",
					},
				},
				"nextPageCursor": "eyJtaW5JRCI6MTc4OTYwNSwibWF4SUQiOjE3ODk2MDV9",
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

		_, err = client.V5().Asset().GetUniversalTransferRecords(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetDepositRecords(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetDepositRecordsParam{}

		path := "/v5/asset/deposit/query-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows": []map[string]interface{}{
					{
						"coin":          "USDT",
						"chain":         "ETH",
						"amount":        "10000",
						"txID":          "skip-notification-scene-test-amount-202304160106-146940-USDT",
						"status":        3,
						"toAddress":     "test-amount-address",
						"tag":           "",
						"depositFee":    "",
						"successAt":     "1681607166000",
						"confirmations": "10000",
						"txIndex":       "",
						"blockHash":     "",
					},
				},
				"nextPageCursor": "eyJtaW5JRCI6MTEzMzUzNSwibWF4SUQiOjExMzM1MzV9",
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

		resp, err := client.V5().Asset().GetDepositRecords(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetDepositRecordsParam{}

		path := "/v5/asset/deposit/query-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows": []map[string]interface{}{
					{
						"coin":          "USDT",
						"chain":         "ETH",
						"amount":        "10000",
						"txID":          "skip-notification-scene-test-amount-202304160106-146940-USDT",
						"status":        3,
						"toAddress":     "test-amount-address",
						"tag":           "",
						"depositFee":    "",
						"successAt":     "1681607166000",
						"confirmations": "10000",
						"txIndex":       "",
						"blockHash":     "",
					},
				},
				"nextPageCursor": "eyJtaW5JRCI6MTEzMzUzNSwibWF4SUQiOjExMzM1MzV9",
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

		_, err = client.V5().Asset().GetDepositRecords(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetSubDepositRecords(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetSubDepositRecordsParam{}

		path := "/v5/asset/deposit/query-sub-member-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows":           []map[string]interface{}{},
				"nextPageCursor": "",
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

		resp, err := client.V5().Asset().GetSubDepositRecords(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetSubDepositRecordsParam{}

		path := "/v5/asset/deposit/query-sub-member-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows":           []map[string]interface{}{},
				"nextPageCursor": "",
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

		_, err = client.V5().Asset().GetSubDepositRecords(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetInternalDepositRecords(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetInternalDepositRecordsParam{}

		path := "/v5/asset/deposit/query-internal-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows":           []map[string]interface{}{},
				"nextPageCursor": "",
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

		resp, err := client.V5().Asset().GetInternalDepositRecords(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetInternalDepositRecordsParam{}

		path := "/v5/asset/deposit/query-internal-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows":           []map[string]interface{}{},
				"nextPageCursor": "",
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

		_, err = client.V5().Asset().GetInternalDepositRecords(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetMasterDepositAddress(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetMasterDepositAddressParam{}

		path := "/v5/asset/deposit/query-address"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"coin": "eth",
				"chains": []interface{}{
					map[string]interface{}{
						"chainType":         "ERC20",
						"addressDeposit":    "0xd9e1cd77afa0e50b452a62fbb68a3340602286c3",
						"tagDeposit":        "",
						"chain":             "ETH",
						"batchReleaseLimit": "-1",
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

		resp, err := client.V5().Asset().GetMasterDepositAddress(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetMasterDepositAddressParam{}

		path := "/v5/asset/deposit/query-address"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"coin": "eth",
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

		_, err = client.V5().Asset().GetMasterDepositAddress(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetWithdrawalRecords(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetWithdrawalRecordsParam{}

		path := "/v5/asset/withdraw/query-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows":           []map[string]interface{}{},
				"nextPageCursor": "",
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

		resp, err := client.V5().Asset().GetWithdrawalRecords(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetWithdrawalRecordsParam{}

		path := "/v5/asset/withdraw/query-record"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows":           []map[string]interface{}{},
				"nextPageCursor": "",
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

		_, err = client.V5().Asset().GetWithdrawalRecords(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetCoinInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetCoinInfoParam{}

		path := "/v5/asset/coin/query-info"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows": []map[string]interface{}{
					{
						"name":         "BTC",
						"coin":         "BTC",
						"remainAmount": "1500",
						"chains": []map[string]interface{}{
							{
								"chain":                 "BTC",
								"chainType":             "BTC",
								"confirmation":          "1",
								"withdrawFee":           "0.0005",
								"depositMin":            "0.0005",
								"withdrawMin":           "0.001",
								"minAccuracy":           "8",
								"chainDeposit":          "1",
								"chainWithdraw":         "1",
								"withdrawPercentageFee": "0",
								"contractAddress":       "",
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

		resp, err := client.V5().Asset().GetCoinInfo(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetCoinInfoParam{}

		path := "/v5/asset/coin/query-info"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"rows": []map[string]interface{}{
					{
						"name":         "BTC",
						"coin":         "BTC",
						"remainAmount": "1500",
						"chains": []map[string]interface{}{
							{
								"chain":                 "BTC",
								"chainType":             "BTC",
								"confirmation":          "1",
								"withdrawFee":           "0.0005",
								"depositMin":            "0.0005",
								"withdrawMin":           "0.001",
								"minAccuracy":           "8",
								"chainDeposit":          "1",
								"chainWithdraw":         "1",
								"withdrawPercentageFee": "0",
								"contractAddress":       "",
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
			WithBaseURL(server.URL)

		_, err = client.V5().Asset().GetCoinInfo(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_GetAllCoinsBalance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetAllCoinsBalanceParam{}

		path := "/v5/asset/transfer/query-account-coins-balance"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"memberId":    "XXXX",
				"accountType": "FUND",
				"balance": []map[string]interface{}{
					{
						"coin":            "USDC",
						"transferBalance": "0",
						"walletBalance":   "0",
						"bonus":           "",
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

		resp, err := client.V5().Asset().GetAllCoinsBalance(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5GetAllCoinsBalanceParam{}

		path := "/v5/asset/transfer/query-account-coins-balance"
		method := http.MethodGet
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"memberId":    "XXXX",
				"accountType": "FUND",
				"balance": []map[string]interface{}{
					{
						"coin":            "USDC",
						"transferBalance": "0",
						"walletBalance":   "0",
						"bonus":           "",
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

		_, err = client.V5().Asset().GetAllCoinsBalance(param)
		assert.Error(t, err)
	})
}

func TestV5Asset_Withdraw(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5WithdrawParam{
			Coin:        CoinETH,
			Chain:       testhelper.Ptr("ETH"),
			Address:     "0x99ced129603abc771c0dabe935c326ff6c86645d",
			Tag:         nil,
			Amount:      "24",
			Timestamp:   1672196561407,
			ForceChain:  testhelper.Ptr(true),
			AccountType: testhelper.Ptr(AccountTypeV5FUND),
			FeeType:     testhelper.Ptr(0),
		}

		path := "/v5/asset/withdraw/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"id": "42c0cfb0-6bca-c242-bc76-4e6df6cbcb16",
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

		resp, err := client.V5().Asset().Withdraw(param)
		require.NoError(t, err)

		require.NotNil(t, resp)
		fmt.Println(resp.Result, respBody["result"])
		testhelper.Compare(t, respBody["result"], resp.Result)
	})
	t.Run("authentication required", func(t *testing.T) {
		param := V5WithdrawParam{
			Coin:        CoinETH,
			Chain:       testhelper.Ptr("ETH"),
			Address:     "0x99ced129603abc771c0dabe935c326ff6c86645d",
			Tag:         nil,
			Amount:      "24",
			Timestamp:   1672196561407,
			ForceChain:  testhelper.Ptr(true),
			AccountType: testhelper.Ptr(AccountTypeV5FUND),
			FeeType:     testhelper.Ptr(0),
		}

		path := "/v5/asset/withdraw/create"
		method := http.MethodPost
		status := http.StatusOK
		respBody := map[string]interface{}{
			"result": map[string]interface{}{
				"transferId": "42c0cfb0-6bca-c242-bc76-4e6df6cbcb16",
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

		_, err = client.V5().Asset().Withdraw(param)
		assert.Error(t, err)
	})
}
