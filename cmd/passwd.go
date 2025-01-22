package cmd

import (
	"fmt"
	"math/rand"

	"github.com/jtprogru/gch/internal/passwd"
	"github.com/spf13/cobra"
)

// passwdCmd represents the passwd command
var (
	passwdCmd = &cobra.Command{
		Use:   "passwd",
		Short: "Generate random password",
		Long:  `Simple password generations`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(passwd.GetPasswd(passwdLength, passwdDigits, passwdSpecials))
		},
	}
	passwdLength   int
	passwdDigits   bool
	passwdSpecials bool
	r              rand.Rand
)

func init() {
	rootCmd.AddCommand(passwdCmd)
	passwdCmd.Flags().IntVarP(&passwdLength, "length", "l", 24, "Length of password")
	passwdCmd.Flags().BoolVarP(&passwdDigits, "digits", "d", false, "Present digits in password")
	passwdCmd.Flags().BoolVarP(&passwdSpecials, "specials", "s", false, "Present special symbols in password")
}
