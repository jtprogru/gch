package cmd

import (
	"fmt"

	"github.com/jtprogru/gch/internal/cbrf"
	"github.com/spf13/cobra"
)

var (
	// cbrfCmd represents the cbrf command
	cbrfCmd = &cobra.Command{
		Use:   "cbrf",
		Short: "Get currency exchange rates for RUB/USD and RUB/EUR",
		Long:  `Get currency exchange rates for all currency from CBRF`,
		Run: func(cmd *cobra.Command, args []string) {
			er := cbrf.CbrfValutes{}
			if !showAll {
				if err := er.ShortRates(); err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				if err := er.FullRates(); err != nil {
					fmt.Println("Error:", err)
				}
			}
		},
	}
	showAll bool
)

func init() {
	rootCmd.AddCommand(cbrfCmd)

	cbrfCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all rates")
}
