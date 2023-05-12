package main

import (
	"regexp"
)

var hashRegex = regexp.MustCompile(`^[a-f0-9]+$`)

func validateHash(input string) bool {
	// Check string is valid 256-bit hex
	if len(input) != 64 {
		return false
	}
	if !hashRegex.MatchString(input) {
		return false
	}

	// Reject hash that matches salt (i.e. an empty string)
	if input == "a61f97a8aa10024cc73d4428117396e184926870935054446e84c1c4290da6c3" {
		return false
	}

	return true
}
