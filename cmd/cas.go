package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jtprogru/gch/internal/cas"
	"github.com/spf13/cobra"
)

var defaultTimeout = 15
var verbose bool

// casCmd represents the cas command
var casCmd = &cobra.Command{
	Use:   "cas",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		casClinet := cas.New(defaultTimeout, verbose)
		if len(args) < 1 {
			fmt.Println("Please provide a user id")
			os.Exit(1)
		}
		userId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid user id! Please provide a number")
		}
		status, err := casClinet.Check(userId)
		if err != nil {
			fmt.Println("Error while checking user status")
			os.Exit(1)
		}
		if !status {
			fmt.Println("User is not in the CAS list")
		} else {
			fmt.Println("User is in the CAS list")
		}
	},
}

func init() {
	rootCmd.AddCommand(casCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// casCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	casCmd.Flags().BoolVarP(&verbose, "verbode", "v", false, "Help message for toggle")
}
