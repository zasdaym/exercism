// Package pangram provides Pangram solution.
package pangram

import (
	"unicode"
)

// IsPangram determines whether input is a pangram.
func IsPangram(input string) bool {
	count := make(map[rune]int)

	for _, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}

		count[unicode.ToLower(c)]++
	}

	for i := 0; i < 26; i++ {
		if count[rune(i)+'a'] == 0 {
			return false
		}
	}

	return true
}
