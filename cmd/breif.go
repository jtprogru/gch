package cmd

import (
	"fmt"

	"github.com/jtprogru/gch/internal/yandexgpt"
	"github.com/spf13/cobra"
)

// breifCmd represents the breif command
var breifCmd = &cobra.Command{
	Use:   "breif",
	Short: "Generate a short description for your long URL",
	Long: `Create short link with breif.
As backend now used https://300.ya.ru
For work with Yandex GPT please read official docs:
https://300.ya.ru/fwv2g5yd/#
And set TOKEN_300_YA_RU environment variable with token.

Usage: gch breif <url>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: gch breif <url>")
			return
		}

		myUrl := args[0]
		res, err := yandexgpt.Breif(myUrl)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(breifCmd)
}
