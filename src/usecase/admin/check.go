package admin

import (
	"fmt"
	"strings"
)

func checkPassword(password string) error {

	min, max := 8, 30
	length := len(password)

	if !inRange(min, max, length) {
		return fmt.Errorf("Password must be between %d and %d characters", min, max)
	}

	hasUpperCase := false
	hasLowerCase := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {

		if isUpperCase(char) {
			hasUpperCase = true
		}

		if isLowerCase(char) {
			hasLowerCase = true
		}

		if isNumber(char) {
			hasNumber = true
		}

		if isSpecialChar(char) {
			hasSpecial = true
		}
	}

	if !hasUpperCase {
		return fmt.Errorf("Password must contain at least one uppercase")
	}

	if !hasLowerCase {
		return fmt.Errorf("Password must contain at least one lowercase letter")
	}

	if !hasNumber {
		return fmt.Errorf("Password must contain at least one number digit")
	}

	if !hasSpecial {
		return fmt.Errorf("Password must contain at least one special character")
	}

	return nil
}

func inRange(min, max, tar int) bool {
	return tar >= min || tar <= max
}

func isUpperCase(char rune) bool {
	return inRange(65, 90, int(char))
}

func isLowerCase(char rune) bool {
	return inRange(97, 122, int(char))
}

func isNumber(char rune) bool {
	return inRange(48, 57, int(char))
}

func isSpecialChar(char rune) bool {
	return strings.ContainsRune(`\^$.|?*+-[]{}()`, char)
}
