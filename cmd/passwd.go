/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

const (
	// lowerLetters is the list of lowercase letters.
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// upperLetters is the list of uppercase letters.
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// digits is the list of permitted digits.
	digits = "0123456789"

	// symbols is the list of symbols.
	symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

// passwdCmd represents the passwd command
var (
	passwdCmd = &cobra.Command{
		Use:   "passwd",
		Short: "Generate random password",
		Long: `Simple password generations
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getPasswd(passwdLength, passwdDigits, passwdSpecials))
		},
	}
	passwdLength   int
	passwdDigits   bool
	passwdSpecials bool
	r              rand.Rand
)

func init() {
	rootCmd.AddCommand(passwdCmd)
	r = *rand.New(rand.NewSource(time.Now().UnixMicro()))
	passwdCmd.Flags().IntVarP(&passwdLength, "length", "l", 24, "Length of password")
	passwdCmd.Flags().BoolVar(&passwdDigits, "digits", true, "Present digits in password")
	passwdCmd.Flags().BoolVar(&passwdSpecials, "specials", true, "Present special symbols in password")
}

func getPasswd(passwdLen int, digitFlag bool, specFlag bool) string {
	var alphabet string
	var password string = ""
	alphabet = lowerLetters + upperLetters
	if digitFlag {
		alphabet += digits
	}
	if specFlag {
		alphabet += symbols
	}
	for i := 0; i < passwdLen; i++ {
		password += string(alphabet[r.Intn(len(alphabet))])
	}
	return password
}
