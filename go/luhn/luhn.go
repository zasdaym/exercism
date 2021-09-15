// Package luhn provides Luhn solution.
package luhn

import (
	"strings"
	"unicode"
)

// Valid checks whether s is valid per the Luhn formula.
func Valid(s string) bool {
	stripped := strings.ReplaceAll(strings.TrimSpace(s), " ", "")
	if len(stripped) < 2 {
		return false
	}

	// Process the digits backward.
	sum := 0
	for i := 0; i < len(stripped); i++ {
		c := rune(stripped[len(stripped)-1-i])

		if !unicode.IsDigit(c) {
			return false
		}

		// Convert char to integer, not its rune value.
		n := int(c - '0')

		if (i+1)%2 == 0 {
			sum += processDigit(n)
		} else {
			sum += n
		}

	}

	return sum%10 == 0
}

// processDigit doubles n and if the result is greater than 9,
// then subtracts 9 from the result.
func processDigit(n int) int {
	doubled := n * 2
	if doubled > 9 {
		return doubled - 9
	}
	return doubled
}
