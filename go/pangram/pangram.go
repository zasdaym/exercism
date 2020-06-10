// Package pangram is solution for problem Pangram.
package pangram

import (
	"strings"
	"unicode"
)

// IsPangram checks if given string is a pangram.
func IsPangram(s string) bool {
	if len(s) < 26 {
		return false
	}

	checker := make(map[rune]bool)
	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) {
			checker[c] = true
		}
	}

	return len(checker) == 26
}
