package fugle_marketdata

import (
	"testing"
)

func TestNewConnection(t *testing.T) {
	client, err := Dial(WebSocketClientOption{})
	if err != nil {
		t.Fatalf("Dial() failed: %v", err)
	}

	if client.Conn == nil {
		t.Fatalf("client.Conn is nil")
	}
}
