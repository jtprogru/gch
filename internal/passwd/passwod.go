package passwd

import (
	"time"

	"math/rand"
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted Digits.
	Digits = "0123456789"

	// Symbols is the list of Symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var r = *rand.New(rand.NewSource(time.Now().UnixMicro()))

func GetPasswd(passwdLen int, digitFlag bool, specFlag bool) string {
	var alphabet string
	var password string = ""
	alphabet = LowerLetters + UpperLetters
	if digitFlag {
		alphabet += Digits
	}
	if specFlag {
		alphabet += Symbols
	}
	for i := 0; i < passwdLen; i++ {
		password += string(alphabet[r.Intn(len(alphabet))])
	}
	return password
}
