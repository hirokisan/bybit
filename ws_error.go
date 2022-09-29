package bybit

import "github.com/gorilla/websocket"

// IsErrWebsocketClosed :
func IsErrWebsocketClosed(err error) bool {
	return websocket.IsCloseError(err, websocket.CloseNormalClosure)
}
