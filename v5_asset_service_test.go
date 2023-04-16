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
