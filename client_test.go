package bybit_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Run("rate limit error", func(t *testing.T) {
		path := "/test"
		method := http.MethodGet
		status := http.StatusOK
		respBody := struct {
			RetCode          int    `json:"ret_code"`
			RetMsg           string `json:"ret_msg"`
			ExtCode          string `json:"ext_code"`
			ExtInfo          string `json:"ext_info"`
			TimeNow          string `json:"time_now"`
			RateLimitStatus  int    `json:"rate_limit_status"`
			RateLimitResetMs int    `json:"rate_limit_reset_ms"`
			RateLimit        int    `json:"rate_limit"`
		}{
			RetCode:          10006,
			RetMsg:           "Too many visits!",
			ExtCode:          "",
			ExtInfo:          "",
			TimeNow:          "1664610970.291886",
			RateLimitStatus:  0,
			RateLimitResetMs: 1664611023016,
			RateLimit:        75,
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		req, err := http.NewRequest(method, server.URL+path, nil)
		require.NoError(t, err)

		client := bybit.NewTestClient().
			WithBaseURL(server.URL)

		var got interface{}
		var wantErr *bybit.RateLimitError

		gotErr := client.Request(req, &got)
		assert.ErrorAs(t, gotErr, &wantErr)
	})

	t.Run("34036, leverage not modified", func(t *testing.T) {
		path := "/test"
		method := http.MethodGet
		status := http.StatusOK
		respBody := struct {
			RetCode          int    `json:"ret_code"`
			RetMsg           string `json:"ret_msg"`
			ExtCode          string `json:"ext_code"`
			ExtInfo          string `json:"ext_info"`
			TimeNow          string `json:"time_now"`
			RateLimitStatus  int    `json:"rate_limit_status"`
			RateLimitResetMs int    `json:"rate_limit_reset_ms"`
			RateLimit        int    `json:"rate_limit"`
		}{
			RetCode:          34036,
			RetMsg:           "leverage not modified",
			ExtCode:          "",
			ExtInfo:          "",
			TimeNow:          "1664610970.291886",
			RateLimitStatus:  0,
			RateLimitResetMs: 1664611023016,
			RateLimit:        75,
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		req, err := http.NewRequest(method, server.URL+path, nil)
		require.NoError(t, err)

		client := bybit.NewTestClient().
			WithBaseURL(server.URL)

		var got interface{}
		var wantErr *bybit.ErrorResponse

		gotErr := client.Request(req, &got)
		assert.ErrorAs(t, gotErr, &wantErr)
		t.Log(gotErr)
	})

	t.Run("10005, API key is invalid", func(t *testing.T) {
		path := "/test"
		method := http.MethodGet
		status := http.StatusOK
		respBody := struct {
			RetCode          int         `json:"ret_code"`
			RetMsg           string      `json:"ret_msg"`
			ExtCode          string      `json:"ext_code"`
			ExtInfo          string      `json:"ext_info"`
			TimeNow          string      `json:"time_now"`
			RateLimitStatus  int         `json:"rate_limit_status"`
			RateLimitResetMs int         `json:"rate_limit_reset_ms"`
			RateLimit        int         `json:"rate_limit"`
			Result           interface{} `json:"result"`
		}{
			RetCode:          10005,
			RetMsg:           "Permission denied, please check your API key permissions",
			ExtCode:          "",
			ExtInfo:          "",
			TimeNow:          "1674306441.372007",
			RateLimitStatus:  0,
			RateLimitResetMs: 0,
			RateLimit:        0,
			Result:           struct{}{},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		req, err := http.NewRequest(method, server.URL+path, nil)
		require.NoError(t, err)

		client := bybit.NewTestClient().
			WithBaseURL(server.URL)

		var got interface{}
		var wantErr *bybit.ErrorResponse

		gotErr := client.Request(req, &got)
		assert.ErrorAs(t, gotErr, &wantErr)
		t.Log(gotErr)
	})

	t.Run("403, access denied", func(t *testing.T) {
		path := "/test"
		method := http.MethodGet
		status := http.StatusForbidden
		respBody := struct {
			Message string `json:"message"`
		}{
			Message: "ok",
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)
		server, teardown := testhelper.NewServer(
			testhelper.WithHandlerOption(path, method, status, bytesBody),
		)
		defer teardown()

		req, err := http.NewRequest(method, server.URL+path, nil)
		require.NoError(t, err)

		client := bybit.NewTestClient().
			WithBaseURL(server.URL)

		var got interface{}

		gotErr := client.Request(req, &got)
		assert.ErrorIs(t, gotErr, bybit.ErrAccessDenied)
	})
	t.Run("404, path not found", func(t *testing.T) {
		path := "/test"
		method := http.MethodGet

		server, teardown := testhelper.NewServer()
		defer teardown()

		req, err := http.NewRequest(method, server.URL+path, nil)
		require.NoError(t, err)

		client := bybit.NewTestClient().
			WithBaseURL(server.URL)

		var got interface{}

		gotErr := client.Request(req, &got)
		assert.ErrorIs(t, gotErr, bybit.ErrPathNotFound)
	})
}
