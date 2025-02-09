package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/cas"
)

var (
	defaultTimeout = 15 //nolint:gochecknoglobals,nolintlint // This is normal.
	verbose        bool //nolint:gochecknoglobals,nolintlint // This is normal.
)

// casCmd represents the cas command.
var casCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
	Use:   "cas",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, args []string) {
		casClinet := cas.New(defaultTimeout, verbose)
		if len(args) < 1 {
			_, _ = fmt.Println("Please provide a user id") //nolint:errcheck,nolintlint // Ignore errors.
			os.Exit(1)                                     //nolint:revive,nolintlint // This is normal.
		}
		userID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			_, _ = fmt.Println("Invalid user id! Please provide a number") //nolint:errcheck,nolintlint // Ignore errors.
		}
		status, err := casClinet.Check(userID)
		if err != nil {
			_, _ = fmt.Println("Error while checking user status") //nolint:errcheck,nolintlint // Ignore errors.
			os.Exit(1)                                             //nolint:revive,nolintlint // This is normal.
		}
		if !status {
			_, _ = fmt.Println("✅ User is not in the CAS list") //nolint:errcheck,nolintlint // Ignore errors.
		} else {
			_, _ = fmt.Println("⛔ User is in the CAS list") //nolint:errcheck,nolintlint // Ignore errors.
		}
	},
}

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(casCmd)
	casCmd.Flags().BoolVarP(&verbose, "verbode", "v", false, "Help message for toggle")
}
