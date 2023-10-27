package lotrsdk

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test_api_key")
	if client.APIKey != "test_api_key" {
		t.Errorf("Expected API Key to be 'test_api_key', got %s", client.APIKey)
	}
	if client.HTTPClient == nil {
		t.Error("Expected HTTPClient not to be nil")
	}
}
