package validators

import "unicode"

func validateHash(input string) bool {
	expectedLength := 64

	// Check length
	if len(input) != expectedLength {
		return false
	}

	// Check character set
	for _, char := range input {
		if !unicode.IsLower(char) && !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}
