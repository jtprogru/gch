package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/clckru"
)

const (
	defaultUtmTags = "utm_source=gch&utm_medium=console"
)

// shortCmd represents the short command.
var shortCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
	Use:   "short",
	Short: "Make short link from URL",
	Long: `Create short link from URL.
	As backend now used https://clck.ru

	Usage: gch short <url> [utm_tag=utm_value]`,
	Run: func(_ *cobra.Command, args []string) {
		// Check if a URL is provided as an argument.
		if len(args) < 1 {
			_, _ = fmt.Println("Usage: gch short <url>") //nolint:errcheck,nolintlint // Ignore errors for test.
			return
		}

		longURL := args[0]
		res, err := clckru.Shorten(longURL, defaultUtmTags)
		if err != nil {
			_, _ = fmt.Printf("error: %v\n", err) //nolint:errcheck,nolintlint // Ignore errors for test.
		}
		_, _ = fmt.Println(res) //nolint:errcheck,nolintlint // Ignore errors for test.
	},
}

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(shortCmd)
}
