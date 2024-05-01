package fugle_marketdata

import (
	"testing"
)

func TestNewConnection(t *testing.T) {
	client, err := NewWebSocketClient(WebSocketClientOption{})
	if err != nil {
		t.Fatalf("Failed to create a new websocket client: %v", err)
	}

	err = client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to the websocket server: %v", err)
	}
	defer client.Close()

	if client.Conn == nil {
		t.Fatalf("The websocket connection is nil.")
	}
}
