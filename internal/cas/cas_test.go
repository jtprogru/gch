package cas

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	client := New(10, true)
	if client.verbose != true {
		t.Errorf("Expected verbose to be true, got %v", client.verbose)
	}
	if client.httpClient.Timeout != 10*time.Second {
		t.Errorf("Expected timeout to be 10 seconds, got %v", client.httpClient.Timeout)
	}
	if client.logger == nil {
		t.Errorf("Expected logger to be initialized, got nil")
	}
}

func TestCheck(t *testing.T) {
	// Mock HTTP server
	mockResponse := `{"ok": true}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Override the CasApiUrl with the mock server URL
	CasApiUrl = server.URL

	// Create a new client
	logger := log.New(os.Stderr, "test: ", log.LstdFlags)
	client := &CasClient{
		verbose: true,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		logger: logger,
	}

	// Test Check method
	userId := uint64(12345)
	ok, err := client.Check(userId)
	if err != nil {
		t.Fatalf("Check() error: %v", err)
	}
	if !ok {
		t.Errorf("Expected user to be in the CAS list, got %v", ok)
	}
}

func TestCheckUserNotInCasList(t *testing.T) {
	// Mock HTTP server
	mockResponse := `{"ok": false}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Override the CasApiUrl with the mock server URL
	CasApiUrl = server.URL

	// Create a new client
	logger := log.New(os.Stderr, "test: ", log.LstdFlags)
	client := &CasClient{
		verbose: true,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		logger: logger,
	}

	// Test Check method
	userId := uint64(12345)
	ok, err := client.Check(userId)
	if err != nil {
		t.Fatalf("Check() error: %v", err)
	}
	if ok {
		t.Errorf("Expected user not to be in the CAS list, got %v", ok)
	}
}
