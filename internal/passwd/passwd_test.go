package passwd

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/magiconair/properties/assert"
)

func TestGetPasswdLength(t *testing.T) {
	passwdLen := 10
	password := GetPasswd(passwdLen, false, false)

	assert.Equal(t, len(password), passwdLen, fmt.Sprintf("expected password length %d, got %d", passwdLen, len(password)))
}

func TestGetPasswdWithDigits(t *testing.T) {
	passwdLen := 10
	password := GetPasswd(passwdLen, true, false)
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	assert.Equal(t, hasDigit, true, fmt.Sprintf("expected password to contain digits, got %s", password))
}

func TestGetPasswdWithSymbols(t *testing.T) {
	passwdLen := 10
	password := GetPasswd(passwdLen, false, true)
	hasSymbol := false
	for _, char := range password {
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			hasSymbol = true
			break
		}
	}

	assert.Equal(t, hasSymbol, true, fmt.Sprintf("expected password to contain symbols, got %s", password))
}

func TestGetPasswdWithoutDigitsAndSymbols(t *testing.T) {
	passwdLen := 10
	password := GetPasswd(passwdLen, false, false)
	hasInvalidChar := false
	for _, char := range password {
		if unicode.IsDigit(char) || unicode.IsSymbol(char) || unicode.IsPunct(char) {
			hasInvalidChar = true
			break
		}
	}
	assert.Equal(t, hasInvalidChar, false, fmt.Sprintf("expected password to contain only letters, got %s", password))
}
