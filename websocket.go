package fugle_marketdata

import (
	"time"

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
	Conn   *websocket.Conn
	option WebSocketClientOption
}

// Dial is a function used to create a new websocket client.
func Dial(option WebSocketClientOption) (*WebSocketClient, error) {
	return DialWithContext(context.Background(), option)
}

// DialWithContext is a function used to create a new websocket client.
func DialWithContext(ctx context.Context, option WebSocketClientOption) (*WebSocketClient, error) {
	if option.Endpoint == "" {
		option.Endpoint = defaultWebSocketClientEndpoint
	}

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, option.Endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &WebSocketClient{
		Conn:   conn,
		option: option,
	}, nil
}

// Close is a function used to close the websocket connection.
func (c *WebSocketClient) Close() error {
	deadline := time.Now().Add(time.Minute)
	err := c.Conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		deadline,
	)
	if err != nil {
		return err
	}

	return nil
}

// Auth sends an authentication message to the Fugle API.
func (c *WebSocketClient) Auth() error {
	return c.AuthWithKey(c.option.APIKey)
}

// AuthWithKey sends an authentication message to the Fugle API.
func (c *WebSocketClient) AuthWithKey(key string) error {
	return c.Conn.WriteJSON(map[string]any{
		"event": "auth",
		"data": map[string]string{
			"apikey": key,
		},
	})
}
