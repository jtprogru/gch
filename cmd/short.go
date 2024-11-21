package cmd

import (
	"fmt"

	"github.com/jtprogru/gch/internal/clckru"
	"github.com/spf13/cobra"
)

const (
	defaultUtmTags = "utm_source=gch&utm_medium=console"
)

// shortCmd represents the short command
var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "Make short link from URL",
	Long: `Create short link from URL.
	As backend now used https://clck.ru

	Usage: gch short <url> [utm_tag=utm_value]`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if a URL is provided as an argument
		if len(args) < 1 {
			fmt.Println("Usage: gch short <url>")
			return
		}

		longUrl := args[0]
		res, err := clckru.Shorten(longUrl, defaultUtmTags)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(shortCmd)
}
