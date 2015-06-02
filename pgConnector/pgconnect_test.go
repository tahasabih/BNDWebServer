package pgConnector

import "testing"

func TestGoConnect(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Error(err)
	}
}
