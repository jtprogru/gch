package cmd

import (
	"fmt"

	"github.com/jtprogru/gch/internal/cbrf"
	"github.com/spf13/cobra"
)

var (
	// cbrfCmd представляет команду cbrf
	cbrfCmd = &cobra.Command{
		Use:   "cbrf",
		Short: "Get currency exchange rates for RUB/USD and RUB/EUR",
		Long:  `Get currency exchange rates for all currencies from CBRF`,
		Run: func(cmd *cobra.Command, args []string) {
			// Инициализируем клиент CBRF
			client := cbrf.NewClient()

			rates, err := client.GetExchangeRates()
			if err != nil {
				fmt.Println("Error fetching exchange rates:", err)
				return
			}

			if showAll {
				printAllRates(rates)
			} else {
				printShortRates(rates)
			}
		},
	}
	showAll bool
)

func init() {
	// Регистрируем команду, подключая к rootCmd
	rootCmd.AddCommand(cbrfCmd)

	// Добавляем опцию (флаг) к команде
	cbrfCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all rates")
}

// printAllRates выводит все курсы валют
func printAllRates(rates cbrf.Valutes) {
	fmt.Println("All exchange rates:")
	for code, valute := range rates.Valute {
		fmt.Printf("%s: %.2f\n", code, valute.Value)
	}
}

// printShortRates выводит короткую сводку для RUB/USD и RUB/EUR
func printShortRates(rates cbrf.Valutes) {
	usd, okUSD := rates.Valute["USD"]
	eur, okEUR := rates.Valute["EUR"]

	// Проверяем, содержатся ли данные по USD и EUR
	if okUSD {
		fmt.Printf("USD/RUB: %.2f\n", usd.Value)
	} else {
		fmt.Println("USD data is not available.")
	}

	if okEUR {
		fmt.Printf("EUR/RUB: %.2f\n", eur.Value)
	} else {
		fmt.Println("EUR data is not available.")
	}
}
