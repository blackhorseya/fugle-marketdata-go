package fugle_marketdata

import (
	"sync"

	"github.com/gorilla/websocket"
)

// MessageHandler is a function type that represents the message handler.
type MessageHandler func(string)

// ErrorHandler is a function type that represents the error handler.
type ErrorHandler func(error)

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

	onMessage MessageHandler
	onError   ErrorHandler
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
		onMessage:   nil,
		onError:     nil,
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

	go client.listen()

	return nil
}

// Close is a function used to close the websocket connection.
func (client *WebSocketClient) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()

	if !client.isConnected {
		return nil
	}

	client.isConnected = false

	err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
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

// OnMessage is a function used to set the message handler.
func (client *WebSocketClient) OnMessage(handler MessageHandler) {
	client.mu.Lock()
	defer client.mu.Unlock()

	client.onMessage = handler
}

// OnError is a function used to set the error handler.
func (client *WebSocketClient) OnError(handler ErrorHandler) {
	client.mu.Lock()
	defer client.mu.Unlock()

	client.onError = handler
}

func (client *WebSocketClient) listen() {
	for client.isConnected {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if client.onError != nil {
				client.onError(err)
			}
			return
		}

		if client.onMessage != nil {
			client.onMessage(string(message))
		}
	}
}
