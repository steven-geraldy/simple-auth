package handler

import (
	"strings"
	"unicode"
)

func validateUserPhone(phone string) bool {
	if len(phone) < 10 || len(phone) > 13 {
		return false
	}
	if !strings.HasPrefix(phone, "+62") {
		return false
	}
	return true
}

func validateUserName(name string) bool {
	if len(name) < 3 || len(name) > 60 {
		return false
	}
	return true
}
func validateUserPassword(password string) bool {
	if len(password) < 6 || len(password) > 64 {
		return false
	}

	var hasCapital, hasNumber, hasSymbol bool
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasCapital = true
		}
		if unicode.IsDigit(char) {
			hasNumber = true
		}
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			hasSymbol = true
		}
	}

	if !hasCapital || !hasNumber || !hasSymbol {
		return false
	}
	return true
}
