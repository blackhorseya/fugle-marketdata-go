package fugle_marketdata

import (
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

const defaultWebSocketClientEndpoint = "wss://api.fugle.tw/marketdata/v1.0/stock/streaming"

// WebSocketClientOption is a struct that represents the websocket client option.
type WebSocketClientOption struct {
	Endpoint string `json:"endpoint"`
	APIKey   string `json:"apiKey"`
}

// WebSocketClient is a struct that represents the websocket client.
type WebSocketClient struct {
	conn *websocket.Conn
}

// Dial is a function used to create a new websocket client.
func Dial(option *WebSocketClientOption) (*WebSocketClient, error) {
	return DialWithContext(context.Background(), option)
}

// DialWithContext is a function used to create a new websocket client.
func DialWithContext(ctx context.Context, option *WebSocketClientOption) (*WebSocketClient, error) {
	if option.Endpoint == "" {
		option.Endpoint = defaultWebSocketClientEndpoint
	}

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, option.Endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &WebSocketClient{
		conn: conn,
	}, nil
}
