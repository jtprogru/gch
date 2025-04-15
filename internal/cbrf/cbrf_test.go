package cbrf

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const exampleBaseURL = "https://example.com"

func TestGetShortRates(t *testing.T) {
	// Mock HTTP server response for short rates.
	mockResponse := `{
		"Date": "2023-10-01T00:00:00Z",
		"PreviousDate": "2023-09-30T00:00:00Z",
		"PreviousURL": "https://www.cbr-xml-daily.ru/archive/2023/09/30/daily_json.js",
		"Timestamp": "2023-10-01T00:00:00Z",
		"Valute": {
			"USD": {"Value": 100.0},
			"EUR": {"Value": 110.10}
		}
	}`

	wantValutes := map[string]CurrencyDetail{
		"USD": {
			Value: 100.0,
		},
		"EUR": {
			Value: 110.10,
		},
	}

	// Set up mock HTTP server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockResponse)) //nolint:errcheck,nolintlint // Ignore write errors for test.
	}))
	defer server.Close()

	// Configure the client to use the mock server.
	client := NewClient()
	client.BaseURL = server.URL

	// Fetch exchange rates.
	rates, err := client.GetExchangeRates()
	if err != nil {
		t.Fatalf("GetExchangeRates() error: %v", err)
	}

	// Validate short rates (USD and EUR only).
	if usd, ok := rates.Valute["USD"]; !ok || usd.Value != wantValutes["USD"].Value {
		t.Errorf("Invalid USD rate. Got: %v, Expected: %v", usd.Value, wantValutes["USD"].Value)
	}
	if eur, ok := rates.Valute["EUR"]; !ok || eur.Value != wantValutes["EUR"].Value {
		t.Errorf("Invalid EUR rate. Got: %v, Expected: %v", eur.Value, wantValutes["EUR"].Value)
	}
}

func TestGetFullRates(t *testing.T) {
	// Mock HTTP server response for full rates.
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

	// Set up mock HTTP server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockResponse)) //nolint:errcheck,nolintlint // Ignore write errors for test.
	}))
	defer server.Close()

	// Configure the client to use the mock server.
	client := NewClient()
	client.BaseURL = server.URL

	// Fetch exchange rates.
	rates, err := client.GetExchangeRates()
	if err != nil {
		t.Fatalf("GetExchangeRates() error: %v", err)
	}

	// Validate all rates (USD, EUR, and GBP).
	if usd, ok := rates.Valute["USD"]; !ok || usd.Value != 74.32 {
		t.Errorf("Invalid USD rate. Got: %v, Expected: 74.32", usd.Value)
	}
	if eur, ok := rates.Valute["EUR"]; !ok || eur.Value != 87.45 {
		t.Errorf("Invalid EUR rate. Got: %v, Expected: 87.45", eur.Value)
	}
	if gbp, ok := rates.Valute["GBP"]; !ok || gbp.Value != 102.34 {
		t.Errorf("Invalid GBP rate. Got: %v, Expected: 102.34", gbp.Value)
	}
}

func TestShortRatesSuccess(t *testing.T) {
	// Создаем тестовые данные
	valutes := &Valutes{
		Valute: map[string]CurrencyDetail{
			"USD": {Value: 75.50},
			"EUR": {Value: 85.30},
		},
		fetched: true,
	}
	client := &Client{
		BaseURL: exampleBaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	// Выполняем тест
	err := valutes.ShortRates(context.Background(), client)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestShortRatesMissingCurrencies(t *testing.T) {
	// Создаем тестовые данные без USD и EUR
	valutes := &Valutes{
		Valute: map[string]CurrencyDetail{
			"GBP": {Value: 100.00},
		},
		fetched: true,
	}

	client := &Client{
		BaseURL: exampleBaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	// Выполняем тест
	err := valutes.ShortRates(context.Background(), client)
	if err == nil || err.Error() != "USD or EUR not found" {
		t.Fatalf("expected error 'USD or EUR not found', got %v", err)
	}
}

func TestShortRatesFetchError(t *testing.T) {
	// Создаем тестовые данные, чтобы вызвать ошибку при fetch
	valutes := &Valutes{
		Valute:  map[string]CurrencyDetail{},
		fetched: false,
	}

	client := &Client{
		BaseURL: "https://invalid-url", // Некорректный URL для вызова ошибки
		HTTPClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	// Выполняем тест
	err := valutes.ShortRates(context.Background(), client)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestShortRatesPrintError(t *testing.T) {
	// Создаем тестовые данные с некорректными значениями
	valutes := &Valutes{
		Valute: map[string]CurrencyDetail{
			"USD": {Value: 0},
			"EUR": {Value: 0},
		},
		fetched: true,
	}

	client := &Client{
		BaseURL: exampleBaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	// Перехватываем вывод, чтобы проверить ошибки
	err := valutes.ShortRates(context.Background(), client)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
