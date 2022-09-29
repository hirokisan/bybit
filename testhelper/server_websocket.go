package testhelper

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/websocket"
)

// NewWebsocketServer :
func NewWebsocketServer(opts ...func(*http.ServeMux)) (*httptest.Server, func()) {
	mux := http.NewServeMux()
	for _, opt := range opts {
		opt(mux)
	}
	ts := httptest.NewServer(mux)
	ts.URL = makeWsProtocol(ts.URL)

	return ts, ts.Close
}

var upgrader = websocket.Upgrader{}

// WithWebsocketHandlerOption :
func WithWebsocketHandlerOption(
	path string,
	respBody []byte,
) func(*http.ServeMux) {
	return func(mux *http.ServeMux) {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			c, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Print("upgrade:", err)
				return
			}
			defer c.Close()
			for {
				mt, _, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				if err := c.WriteMessage(mt, respBody); err != nil {
					log.Println("write:", err)
					break
				}
			}
		})
	}
}

func makeWsProtocol(u string) string {
	if strings.HasPrefix(u, "https") {
		return "wss" + strings.TrimPrefix(u, "https")
	}
	if strings.HasPrefix(u, "http") {
		return "ws" + strings.TrimPrefix(u, "http")
	}
	return u
}
