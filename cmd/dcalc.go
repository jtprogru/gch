package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/datescalculator"
)

var (
	flagDaysCount  int    //nolint:gochecknoglobals,nolintlint // This is normal.
	flagDate       string //nolint:gochecknoglobals,nolintlint // This is normal.
	flagJSONOutput bool   //nolint:gochecknoglobals,nolintlint // This is normal.
)

// dcalcCmd represents the dcalc command.
var dcalcCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
	Use:   "dcalc",
	Short: "Calculate dates",
	Long: `This command calculate "some" days in
  past and future from defained date.
  By default (if date not define as args), calculate
  "some" days from current date.

  Usage: gch dcalc -d 2025-02-07 -c 10`,
	Run: func(_ *cobra.Command, _ []string) {
		var res *datescalculator.DatesResponse
		var err error
		today := time.Now().Format("2006-01-02")

		if flagDate != "" {
			res, err = datescalculator.Calc(flagDate, flagDaysCount)
		} else {
			res, err = datescalculator.Calc(today, flagDaysCount)
		}

		if err != nil {
			_, _ = fmt.Printf("Calculate err: %s", err) //nolint:errcheck,nolintlint // Ignore errors.
			os.Exit(1)                                  //nolint:revive,nolintlint // This is normal.
		}

		if !flagJSONOutput {
			outputInText(res)
			return
		}
		outputInJSON(res)
	},
}

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(dcalcCmd)
	dcalcCmd.Flags().StringVarP(&flagDate, "date", "d", "", "Date from")
	dcalcCmd.Flags().IntVarP(&flagDaysCount, "count", "c", 10, "How much days count from date")
	dcalcCmd.Flags().BoolVarP(&flagJSONOutput, "json", "j", false, "Output in JSON format")
}

func outputInJSON(r *datescalculator.DatesResponse) {
	prettyJSON, _ := json.MarshalIndent(r, "", "  ") //nolint:errcheck,nolintlint // Ignore errors.
	_, _ = fmt.Println(string(prettyJSON))           //nolint:errcheck,nolintlint // Ignore errors.
}

func outputInText(r *datescalculator.DatesResponse) {
	_, _ = fmt.Printf("Date %s +/- %d days:\n", r.Today, flagDaysCount) //nolint:errcheck,nolintlint // Ignore errors.
	_, _ = fmt.Printf("+%d days = %s\n", flagDaysCount, r.FutureDate)   //nolint:errcheck,nolintlint // Ignore errors.
	_, _ = fmt.Printf("-%d days = %s\n", flagDaysCount, r.PastDate)     //nolint:errcheck,nolintlint // Ignore errors.
}
