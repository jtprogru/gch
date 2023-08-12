/*
Copyright © 2023 Michael <jtprogru@gmail> Savin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
