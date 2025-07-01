package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
    headers := http.Header{}
    headers.Set("Authorization", "ApiKey abc123def456")
    
    got, err := GetAPIKey(headers)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    expected := "abc123def456"
    if got != expected {
        t.Errorf("Expected %s, got %s", expected, got)
    }

}