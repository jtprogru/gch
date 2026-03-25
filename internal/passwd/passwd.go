package passwd

import (
	"crypto/rand"
	"math/big"
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]:<>?,./"
)

// Config struct for configuration of GetPasswd.
type Config struct {
	Length         int
	IncludeDigits  bool
	IncludeSymbols bool
}

// GetPasswd generates a secure password.
func GetPasswd(cfg Config) string {
	alphabet := LowerLetters + UpperLetters

	if cfg.IncludeDigits {
		alphabet += Digits
	}
	if cfg.IncludeSymbols {
		alphabet += Symbols
	}

	password := make([]byte, cfg.Length)
	for i := range cfg.Length {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		password[i] = alphabet[idx.Int64()]
	}

	return string(password)
}
