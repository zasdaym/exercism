// Package isbn is solution for problem ISBN Verifier.
package isbn

import (
	"regexp"
)

var isbnRegex = regexp.MustCompile(`[^0-9X]+`)

// IsValidISBN checks if given string is a valid ISBN.
func IsValidISBN(s string) bool {
	numbers := isbnRegex.ReplaceAllString(s, "")
	if len(numbers) != 10 {
		return false
	}

	sum := 0
	for i, r := range numbers {
		checkDigit := int(r - '0')
		if r == 'X' {
			if i != len(numbers)-1 {
				return false
			}
			checkDigit = 10
		}
		sum += (10 - i) * checkDigit
	}
	return sum%11 == 0
}
