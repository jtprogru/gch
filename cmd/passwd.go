/*
Copyright © 2023 Michael <jtprogru@gmail> Savin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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
