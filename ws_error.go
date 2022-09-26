package bybit

import "github.com/gorilla/websocket"

// IsErrWebsocketClosed :
func IsErrWebsocketClosed(err error) bool {
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return true
	}
	return false
}
