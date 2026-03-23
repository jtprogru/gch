package utils

import "unicode"

func CheckDigit(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func CheckPunct(s string) bool {
	for _, char := range s {
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			return true
		}
	}
	return false
}

func CheckDigitAndPunt(s string) bool {
	return CheckDigit(s) && CheckPunct(s)
}
