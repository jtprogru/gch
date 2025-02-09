package passwd

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/magiconair/properties/assert"
)

func TestGetPasswdLength(t *testing.T) {
	cfg := Config{
		Length:         64,
		IncludeDigits:  false,
		IncludeSymbols: false,
	}

	password, err := GetPasswd(cfg)
	if err != nil {
		t.Errorf("want no error, got %s", err)
	}

	assert.Equal(t, len(password), cfg.Length, fmt.Sprintf("expected password length %d, got %d", cfg.Length, len(password)))
}

func TestGetPasswdWithDigits(t *testing.T) {
	cfg := Config{
		Length:         32,
		IncludeDigits:  true,
		IncludeSymbols: false,
	}
	password, err := GetPasswd(cfg)
	if err != nil {
		t.Errorf("want no error, got %s", err)
	}

	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}

	assert.Equal(t, hasDigit, true, "expected password to contain digits, got "+password)
}

func TestGetPasswdWithSymbols(t *testing.T) {
	cfg := Config{
		Length:         10,
		IncludeDigits:  false,
		IncludeSymbols: true,
	}
	hasSymbol := false

	password, err := GetPasswd(cfg)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	for _, char := range password {
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			hasSymbol = true
			break
		}
	}

	assert.Equal(t, hasSymbol, true, "expected password to contain symbols, got "+password)
}

func TestGetPasswdWithoutDigitsAndSymbols(t *testing.T) {
	cfg := Config{
		Length:         64,
		IncludeDigits:  false,
		IncludeSymbols: false,
	}

	password, err := GetPasswd(cfg)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	hasInvalidChar := false
	for _, char := range password {
		if unicode.IsDigit(char) || unicode.IsSymbol(char) || unicode.IsPunct(char) {
			hasInvalidChar = true
			break
		}
	}

	assert.Equal(t, hasInvalidChar, false, "expected password to contain only letters, got "+password)
}
