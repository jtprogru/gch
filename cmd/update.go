package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jtprogru/gch/internal/gch"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command.
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gch",
	Long:  `Check available update and upgrade gch.`,
	Run: func(cmd *cobra.Command, args []string) {
		if gch.ShouldCheckForUpdate() {
			fmt.Println("current Version: ", cmd.Version)
			if update {
				ctx := context.Background()
				res, err := gch.CheckForUpdate(ctx, &http.Client{}, "/tmp/gch", "jtprogru/gch", cmd.Version)
				if err != nil {
					_ = fmt.Errorf("gch.CheckForUpdate err: %v", err)
					os.Exit(1)
				}
				fmt.Printf("%+v", res.Version)
				return
			}

		}
	},
}

var update bool

func init() {
	rootCmd.AddCommand(updateCmd)
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().BoolVarP(&update, "upgrade", "u", false, "Upgrade gch")
}
