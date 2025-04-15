package cbrf

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	defaultTimeout = time.Second * 15
)

// Client basic client for CBRF.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Timeout    time.Duration
}

// NewClient create and configure basic client for CBRF.
func NewClient() *Client {
	return &Client{
		BaseURL: "https://www.cbr-xml-daily.ru/daily_json.js",
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
		Timeout: defaultTimeout,
	}
}

func (c *Client) GetExchangeRates() (Valutes, error) {
	valutes := new(Valutes)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	return *valutes, valutes.getExchangeRates(ctx, c)
}

type Valutes struct {
	Date         time.Time                 `json:"date"`
	PreviousDate time.Time                 `json:"previous_date"` //nolint:tagliatelle // This is predefined by the API
	PreviousURL  string                    `json:"previous_url"`  //nolint:tagliatelle // This is predefined by the API
	Timestamp    time.Time                 `json:"timestamp"`
	Valute       map[string]CurrencyDetail `json:"valute"`
	fetched      bool
}

// CurrencyDetail store information about currency.
type CurrencyDetail struct {
	Value float64 `json:"value"`
}

func (v *Valutes) ShortRates(ctx context.Context, cfg *Client) error {
	if !v.fetched {
		if err := v.getExchangeRates(ctx, cfg); err != nil {
			return err
		}
	}

	_, _ = fmt.Println("Exchange Rates:") //nolint:errcheck,nolintlint // Ignore write errors for test.

	rubUsd, okUsd := v.Valute["USD"]
	rubEur, okEur := v.Valute["EUR"]

	if !okUsd || !okEur {
		return errors.New("USD or EUR not found")
	}

	if _, err := fmt.Printf("RUB/USD: %.2f\n", rubUsd.Value); err != nil {
		return err
	}
	if _, err := fmt.Printf("RUB/EUR: %.2f\n", rubEur.Value); err != nil {
		return err
	}
	return nil
}

func (v *Valutes) FullRates(ctx context.Context, cfg *Client) error {
	if !v.fetched {
		if err := v.getExchangeRates(ctx, cfg); err != nil {
			return err
		}
	}

	_, _ = fmt.Println("Exchange Rates:") //nolint:errcheck,nolintlint // Ignore write errors for test.
	for code, currency := range v.Valute {
		if _, err := fmt.Printf("RUB/%s: %.2f\n", code, currency.Value); err != nil {
			return err
		}
	}

	return nil
}

func (v *Valutes) getExchangeRates(ctx context.Context, cfg *Client) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.BaseURL, http.NoBody)
	if err != nil {
		return err
	}

	response, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(v); err != nil {
		return err
	}

	v.fetched = true
	return nil
}
