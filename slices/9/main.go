package main

func isValidPassword(password string) bool {
	// Check length
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	hasUpper := false
	hasDigit := false

	for i := 0; i < len(password); i++ {
		ch := password[i]

		if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		}

		if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}

	return hasUpper && hasDigit
}
