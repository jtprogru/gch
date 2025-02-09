package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/cbrf"
)

var (
	// cbrfCmd представляет команду cbrf.
	cbrfCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
		Use:   "cbrf",
		Short: "Get currency exchange rates for RUB/USD and RUB/EUR",
		Long:  `Get currency exchange rates for all currencies from CBRF`,
		Run: func(_ *cobra.Command, _ []string) {
			// Инициализируем клиент CBRF.
			client := cbrf.NewClient()

			rates, err := client.GetExchangeRates()
			if err != nil {
				_, _ = fmt.Println("Error fetching exchange rates:", err) //nolint:errcheck,nolintlint // Ignore errors.
				return
			}

			if showAll {
				printAllRates(rates)
			} else {
				printShortRates(rates)
			}
		},
	}
	showAll bool //nolint:gochecknoglobals,nolintlint // This is normal.
)

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	// Регистрируем команду, подключая к rootCmd.
	rootCmd.AddCommand(cbrfCmd)

	// Добавляем опцию (флаг) к команде.
	cbrfCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all rates")
}

// printAllRates выводит все курсы валют.
func printAllRates(rates cbrf.Valutes) {
	_, _ = fmt.Println("All exchange rates:") //nolint:errcheck,nolintlint // Ignore errors.
	for code, valute := range rates.Valute {
		_, _ = fmt.Printf("%s: %.2f\n", code, valute.Value) //nolint:errcheck,nolintlint // Ignore errors.
	}
}

// printShortRates выводит короткую сводку для RUB/USD и RUB/EUR.
func printShortRates(rates cbrf.Valutes) {
	usd, okUSD := rates.Valute["USD"]
	eur, okEUR := rates.Valute["EUR"]

	// Проверяем, содержатся ли данные по USD и EUR.
	if okUSD {
		_, _ = fmt.Printf("USD/RUB: %.2f\n", usd.Value) //nolint:errcheck,nolintlint // Ignore errors.
	} else {
		_, _ = fmt.Println("USD data is not available.") //nolint:errcheck,nolintlint // Ignore errors.
	}

	if okEUR {
		_, _ = fmt.Printf("EUR/RUB: %.2f\n", eur.Value) //nolint:errcheck,nolintlint // Ignore errors.
	} else {
		_, _ = fmt.Println("EUR data is not available.") //nolint:errcheck,nolintlint // Ignore errors.
	}
}
