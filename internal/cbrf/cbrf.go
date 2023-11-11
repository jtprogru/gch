package cbrf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	cbrfUrl = "https://www.cbr-xml-daily.ru/daily_json.js"
)

type CbrfValutes struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       map[string]struct {
		Value float64 `json:"Value"`
	} `json:"Valute"`
}

func GetExchangeRates() (*CbrfValutes, error) {
	response, err := http.Get(cbrfUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data CbrfValutes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (exchangeRates *CbrfValutes) OutputAllRates() error {

	fmt.Println("Exchange Rates:")
	for code, currency := range exchangeRates.Valute {
		fmt.Printf("RUB/%s: %.2f\n", code, currency.Value)
	}

	return nil
}
