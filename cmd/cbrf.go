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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/cbrf"
)

var (
	// cbrfCmd represents the cbrf command
	cbrfCmd = &cobra.Command{
		Use:   "cbrf",
		Short: "Get currency exchange rates for RUB/USD and RUB/EUR",
		Long:  `Get currency exchange rates for RUB/USD and RUB/EUR`,
		Run: func(cmd *cobra.Command, args []string) {
			exchangeRates, err := cbrf.GetExchangeRates()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if !showAll {
				fmt.Printf("RUB/USD: %.2f\n", exchangeRates.Valute["USD"].Value)
				fmt.Printf("RUB/EUR: %.2f\n", exchangeRates.Valute["EUR"].Value)
			} else {
				_ = exchangeRates.OutputAllRates()
			}
		},
	}
	showAll bool
)

func init() {
	rootCmd.AddCommand(cbrfCmd)

	cbrfCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all rates")
}
