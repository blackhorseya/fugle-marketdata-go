package fugle_marketdata

import (
	"github.com/gorilla/websocket"
)

const defaultWebSocketClientEndpoint = "wss://api.fugle.tw/marketdata/v1.0/stock/streaming"

// WebSocketClient is a struct that represents the websocket client.
type WebSocketClient struct {
	conn *websocket.Conn
}
