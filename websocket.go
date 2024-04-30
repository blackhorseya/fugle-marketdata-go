package fugle_marketdata

import (
	"sync"
	"time"

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
	Conn *websocket.Conn

	mu          sync.Mutex
	option      WebSocketClientOption
	isConnected bool
}

// NewWebSocketClient is a function used to create a new websocket client.
func NewWebSocketClient(option WebSocketClientOption) (*WebSocketClient, error) {
	if option.Endpoint == "" {
		option.Endpoint = defaultWebSocketClientEndpoint
	}

	return &WebSocketClient{
		Conn:        nil,
		mu:          sync.Mutex{},
		option:      option,
		isConnected: false,
	}, nil
}

// Connect is a function used to connect to the websocket server.
func (client *WebSocketClient) Connect() error {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.isConnected {
		return nil
	}

	conn, resp, err := websocket.DefaultDialer.Dial(client.option.Endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	client.Conn = conn
	client.isConnected = true
	return nil
}

// Close is a function used to close the websocket connection.
func (client *WebSocketClient) Close() error {
	deadline := time.Now().Add(time.Minute)
	err := client.Conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		deadline,
	)
	if err != nil {
		return err
	}

	err = client.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return err
	}

	for {
		_, _, err = client.Conn.NextReader()
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			break
		}
		if err != nil {
			break
		}
	}

	err = client.Conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// IsConnected is a function used to check if the websocket client is connected.
func (client *WebSocketClient) IsConnected() bool {
	client.mu.Lock()
	defer client.mu.Unlock()

	return client.isConnected
}
