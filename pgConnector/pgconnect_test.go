package pgConnector

import "testing"

func TestGoConnect(t *testing.T) {
	isConnected, err := Connect()
	if isConnected != true {
		t.Error(err)
	}
}
