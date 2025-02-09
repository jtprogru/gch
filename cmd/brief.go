package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/yandexgpt"
)

// briefCmd represents the brief command.
var briefCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
	Use:   "brief",
	Short: "Generate a short description for your long URL",
	Long: `Create short link with brief.
As backend now used https://300.ya.ru
For work with Yandex GPT please read official docs:
https://300.ya.ru/fwv2g5yd/#
And set TOKEN_300_YA_RU environment variable with token.

Usage: gch brief <url>`,
	Run: func(_ *cobra.Command, args []string) {
		if len(args) < 1 {
			_, _ = fmt.Println("Usage: gch brief <url>") //nolint:errcheck,nolintlint // Ignore errors.
			return
		}

		myURL := args[0]
		res, err := yandexgpt.Brief(myURL)
		if err != nil {
			_, _ = fmt.Printf("error: %v\n", err) //nolint:errcheck,nolintlint // Ignore errors.
		}
		_, _ = fmt.Println(res) //nolint:errcheck,nolintlint // Ignore errors.
	},
}

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(briefCmd)
}
