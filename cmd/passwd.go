package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/passwd"
)

// passwdCmd represents the passwd command.
var (
	passwdCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
		Use:   "passwd",
		Short: "Generate random password",
		Long:  `Simple password generations`,
		Run: func(_ *cobra.Command, _ []string) {
			cfg := passwd.Config{
				Length:         passwdLength,
				IncludeDigits:  passwdDigits,
				IncludeSymbols: passwdSpecials,
			}

			p, err := passwd.GetPasswd(cfg)
			if err != nil {
				_, _ = fmt.Println("password generation err: ", err) //nolint:errcheck,nolintlint // Ignore errors.
				return
			}
			_, _ = fmt.Println(p) //nolint:errcheck,nolintlint // Ignore errors.
		},
	}
	passwdLength   int  //nolint:gochecknoglobals,nolintlint // This is normal.
	passwdDigits   bool //nolint:gochecknoglobals,nolintlint // This is normal.
	passwdSpecials bool //nolint:gochecknoglobals,nolintlint // This is normal.
)

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(passwdCmd)
	passwdCmd.Flags().IntVarP(&passwdLength, "length", "l", 24, "Length of password")
	passwdCmd.Flags().BoolVarP(&passwdDigits, "digits", "d", false, "Present digits in password")
	passwdCmd.Flags().BoolVarP(&passwdSpecials, "specials", "s", false, "Present special symbols in password")
}
