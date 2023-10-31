package bybit

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func GetDepositRecords(t *testing.T) {
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

func GetSubDepositRecords(t *testing.T) {
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

func GetInternalDepositRecords(t *testing.T) {
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

func TestGetWithdrawalRecords(t *testing.T) {
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

func TestGetCoinInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		param := V5GetCoinInfoParam{}

		path := "/v5/asset/coin/query-info"
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

		_, err = client.V5().Asset().GetCoinInfo(param)
		assert.Error(t, err)
	})
}
