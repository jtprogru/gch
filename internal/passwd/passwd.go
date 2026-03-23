package passwd

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted Digits.
	Digits = "0123456789"

	// Symbols is the list of Symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]:<>?,./"
)

// Config struct for configuration of GetPasswd.
type Config struct {
	Length         int
	IncludeDigits  bool
	IncludeSymbols bool
}

// GetPasswd generates a secure passwordBuilder.
func GetPasswd(cfg Config) string {
	var passwordBuilder strings.Builder
	alphabet := LowerLetters + UpperLetters

	// Add the required character sets to the alphabet.
	if cfg.IncludeDigits {
		alphabet += Digits
	}
	if cfg.IncludeSymbols {
		alphabet += Symbols
	}

	// Generate password.
	for i := 0; i < cfg.Length; i++ {
		// Get a secure random index.
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))

		// Append the character at the generated index to the password.
		passwordBuilder.WriteString(string(alphabet[idx.Int64()]))
	}

	return passwordBuilder.String()
}
