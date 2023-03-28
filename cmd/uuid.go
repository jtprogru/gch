/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var (
	uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "Generate UUID string",
		Long: `A UUID is a 16 byte (128 bit) array.
UUIDs may be used as keys to maps or compared directly.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getUUID())
		},
	}
	v4UUID    bool
	v5UUID    bool
	nilUUID   bool
	namespace = uuid.NameSpaceURL
	url       = []byte("https://jtprog.ru")
)

func init() {
	rootCmd.AddCommand(uuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	uuidCmd.Flags().BoolVarP(&v4UUID, "v4", "4", true, "Generate UUID4")
	uuidCmd.Flags().BoolVarP(&v5UUID, "v5", "5", false, "Generate UUID5")
	uuidCmd.Flags().BoolVarP(&nilUUID, "v0", "0", false, "Generate UUIDnil")

}

func getUUID() string {
	switch {
	case v4UUID:
		return uuid.New().String()
	case v5UUID:
		return uuid.NewMD5(namespace, url).String()
	case nilUUID:
		return "00000000-0000-0000-0000-000000000000"
	default:
		return "00000000-0000-0000-0000-000000000000"
	}
}