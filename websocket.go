package fugle_marketdata

import (
	"github.com/gorilla/websocket"
)

// WebSocketClient is a struct that represents the websocket client.
type WebSocketClient struct {
	conn *websocket.Conn
}
