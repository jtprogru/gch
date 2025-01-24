package cbrf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	cbrfUrl    = "https://www.cbr-xml-daily.ru/daily_json.js"
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

type CbrfValutes struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       map[string]struct {
		Value float64 `json:"Value"`
	} `json:"Valute"`
	fetched bool
}

func (er *CbrfValutes) ShortRates() error {
	if !er.fetched {
		if err := er.getExchangeRates(); err != nil {
			return err
		}
	}

	fmt.Println("Exchange Rates:")

	rubUsd, okUsd := er.Valute["USD"]
	rubEur, okEur := er.Valute["EUR"]

	if !okUsd || !okEur {
		return fmt.Errorf("USD or EUR not found")
	}

	fmt.Printf("RUB/USD: %.2f\n", rubUsd.Value)
	fmt.Printf("RUB/EUR: %.2f\n", rubEur.Value)
	return nil
}

func (er *CbrfValutes) FullRates() error {
	if !er.fetched {
		if err := er.getExchangeRates(); err != nil {
			return err
		}
	}

	fmt.Println("Exchange Rates:")
	for code, currency := range er.Valute {
		fmt.Printf("RUB/%s: %.2f\n", code, currency.Value)
	}

	return nil
}

func (er *CbrfValutes) getExchangeRates() error {
	response, err := httpClient.Get(cbrfUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(er); err != nil {
		return err
	}

	er.fetched = true
	return nil
}
