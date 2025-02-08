package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jtprogru/gch/internal/datescalculator"
	"github.com/spf13/cobra"
)

var (
	flagDaysCount  int
	flagDate       string
	flagJSONOutput bool
)

// dcalcCmd represents the dcalc command
var dcalcCmd = &cobra.Command{
	Use:   "dcalc",
	Short: "Calculate dates",
	Long: `This command calculate "some" days in
  past and future from defained date.
  By default (if date not define as args), calculate
  "some" days from current date.

  Usage: gch dcalc -d 2025-02-07 -c 10`,
	Run: func(cmd *cobra.Command, args []string) {
		var res *datescalculator.DatesResponse
		var err error
		today := time.Now().Format("2006-01-02")

		if flagDate != "" {
			res, err = datescalculator.Calc(flagDate, flagDaysCount)
		} else {
			res, err = datescalculator.Calc(today, flagDaysCount)
		}

		if err != nil {
			fmt.Printf("Calculate err: %s", err)
			os.Exit(1)
		}

		if !flagJSONOutput {
			outputInText(res)
			return
		}
		outputInJSON(res)
	},
}

func init() {
	rootCmd.AddCommand(dcalcCmd)
	dcalcCmd.Flags().StringVarP(&flagDate, "date", "d", "", "Date from")
	dcalcCmd.Flags().IntVarP(&flagDaysCount, "count", "c", 10, "How much days count from date")
	dcalcCmd.Flags().BoolVarP(&flagJSONOutput, "json", "j", false, "Output in JSON format")
}

func outputInJSON(r *datescalculator.DatesResponse) {

	prettyJSON, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(prettyJSON))
}

func outputInText(r *datescalculator.DatesResponse) {
	fmt.Printf("Date %s +/- %d days:\n", r.Today, flagDaysCount)
	fmt.Printf("+%d days = %s\n", flagDaysCount, r.FutureDate)
	fmt.Printf("-%d days = %s\n", flagDaysCount, r.PastDate)
}
