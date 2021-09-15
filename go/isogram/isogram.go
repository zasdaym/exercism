// Package isogram provides Isogram solution.
package isogram

import "unicode"

// IsIsogram checks whether s is an isogram.
func IsIsogram(s string) bool {
	appeared := make(map[rune]bool)
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		normalizedChar := unicode.ToLower(c)
		if appeared[normalizedChar] {
			return false
		}
		appeared[normalizedChar] = true
	}
	return true
}
