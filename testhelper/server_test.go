package testhelper_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/hirokisan/bybit/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	t.Run("handler for the path exists", func(t *testing.T) {
		path := "/test"
		method := http.MethodGet
		status := http.StatusOK
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

		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		gotBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.Equal(t, string(bytesBody), string(gotBody))
		assert.Equal(t, status, resp.StatusCode)
	})
	t.Run("handler for the path not exists", func(t *testing.T) {
		path := "/test"
		method := http.MethodGet

		server, teardown := testhelper.NewServer()
		defer teardown()

		req, err := http.NewRequest(method, server.URL+path, nil)
		require.NoError(t, err)

		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
