package fugle_marketdata

import (
	"github.com/gorilla/websocket"
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

// NewWebSocketClient is a function used to create a new websocket client.
func NewWebSocketClient(option WebSocketClientOption) (*WebSocketClient, error) {
	if option.Endpoint == "" {
		option.Endpoint = defaultWebSocketClientEndpoint
	}

	return &WebSocketClient{
		Conn:   nil,
		option: option,
	}, nil
}

// Connect is a function used to connect to the websocket server.
func (client *WebSocketClient) Connect() error {
	conn, resp, err := websocket.DefaultDialer.Dial(client.option.Endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	client.Conn = conn
	return nil
}

// Close is a function used to close the websocket connection.
func (client *WebSocketClient) Close() error {
	// deadline := time.Now().Add(time.Minute)
	// err := client.Conn.WriteControl(
	// 	websocket.CloseMessage,
	// 	websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	// 	deadline,
	// )
	// if err != nil {
	// 	return err
	// }

	// todo: 2024/4/30|sean|send close message to server

	return client.Conn.Close()
}
