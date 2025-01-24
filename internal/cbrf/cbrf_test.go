package cbrf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortRates(t *testing.T) {
	// Mock HTTP server
	mockResponse := `{
        "Date": "2023-10-01T00:00:00Z",
        "PreviousDate": "2023-09-30T00:00:00Z",
        "PreviousURL": "https://www.cbr-xml-daily.ru/archive/2023/09/30/daily_json.js",
        "Timestamp": "2023-10-01T00:00:00Z",
        "Valute": {
            "USD": {"Value": 74.32},
            "EUR": {"Value": 87.45}
        }
    }`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Override the cbrfUrl with the mock server URL
	cbrfUrl = server.URL

	// Test ShortRates
	valutes := &CbrfValutes{}
	err := valutes.ShortRates()
	if err != nil {
		t.Fatalf("ShortRates() error: %v", err)
	}
}

func TestFullRates(t *testing.T) {
	// Mock HTTP server
	mockResponse := `{
        "Date": "2023-10-01T00:00:00Z",
        "PreviousDate": "2023-09-30T00:00:00Z",
        "PreviousURL": "https://www.cbr-xml-daily.ru/archive/2023/09/30/daily_json.js",
        "Timestamp": "2023-10-01T00:00:00Z",
        "Valute": {
            "USD": {"Value": 74.32},
            "EUR": {"Value": 87.45},
            "GBP": {"Value": 102.34}
        }
    }`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Override the cbrfUrl with the mock server URL
	cbrfUrl = server.URL

	// Test FullRates
	valutes := &CbrfValutes{}
	err := valutes.FullRates()
	if err != nil {
		t.Fatalf("FullRates() error: %v", err)
	}
}
