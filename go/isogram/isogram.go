//Package isogram is solution for problem Isogram
package isogram

import (
	"unicode"
)

// IsIsogram will check if given string is an isogram
func IsIsogram(s string) bool {
	checked := map[rune]bool{}
	for _, c := range s {
		if unicode.IsLetter(c) {
			char := unicode.ToLower(c)
			if checked[char] {
				return false
			}

			checked[char] = true
		}

	}

	return true
}
