package testhelper

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebsocketServer(t *testing.T) {
	t.Run("path exist", func(t *testing.T) {
		path := "/test"
		respBody := struct {
			Message string `json:"message"`
		}{
			Message: "ok",
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)
		server, teardown := NewWebsocketServer(WithWebsocketHandlerOption(path, bytesBody))
		defer teardown()

		c, _, err := websocket.DefaultDialer.Dial(server.URL+path, nil)
		require.NoError(t, err)

		assert.NoError(t, c.WriteMessage(websocket.TextMessage, nil))

		_, message, err := c.ReadMessage()
		require.NoError(t, err)
		assert.Equal(t, bytesBody, message)
	})
	t.Run("path not exist", func(t *testing.T) {
		path := "/test"
		server, teardown := NewWebsocketServer()
		defer teardown()

		_, _, err := websocket.DefaultDialer.Dial(server.URL+path, nil)
		assert.ErrorIs(t, err, websocket.ErrBadHandshake)
	})
}

func TestWebsocketServerWithCustomDialer(t *testing.T) {
	customDialer := &websocket.Dialer{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return nil, nil
		},
		HandshakeTimeout: 5 * time.Second,
	}

	t.Run("custom dialer success", func(t *testing.T) {
		path := "/custom"
		respBody := struct {
			Message string `json:"message"`
		}{
			Message: "custom dialer ok",
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)
		server, teardown := NewWebsocketServer(WithWebsocketHandlerOption(path, bytesBody))
		defer teardown()

		c, _, err := customDialer.Dial(server.URL+path, nil)
		require.NoError(t, err)

		assert.NoError(t, c.WriteMessage(websocket.TextMessage, nil))

		_, message, err := c.ReadMessage()
		require.NoError(t, err)
		assert.Equal(t, bytesBody, message)
	})

	t.Run("custom dialer failure", func(t *testing.T) {
		path := "/custom"
		server, teardown := NewWebsocketServer()
		defer teardown()

		_, _, err := customDialer.Dial(server.URL+path, nil)
		assert.ErrorIs(t, err, websocket.ErrBadHandshake)
	})
}
