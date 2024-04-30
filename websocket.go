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

// ReadMessage reads a message from the Fugle API.
func (c *WebSocketClient) ReadMessage() (ch chan IEvent, err error) {
	retCh := make(chan IEvent)

	go func() {
		for {
			_, message, err2 := c.Conn.ReadMessage()
			if err2 != nil {
				continue
			}

			event, err2 := UnmarshalEvent(message)
			if err2 != nil {
				continue
			}

			retCh <- event
		}
	}()

	return retCh, nil
}
