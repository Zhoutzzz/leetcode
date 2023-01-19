package easy

import (
	"strings"
	"unicode"
)

func strongPasswordCheckerII(password string) bool {
	n := len(password)
	if n < 8 {
		return false
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool
	for i, ch := range password {
		if i != n-1 && password[i] == password[i+1] {
			return false
		}
		if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else if strings.ContainsRune("!@#$%^&*()-+", ch) {
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}
