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
		Long:  `Get currency exchange rates for all currency from CBRF`,
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
