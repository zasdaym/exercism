// Package isbn provides ISBN Verifier solution.
package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN checks whether the given string is a valid ISBN.
func IsValidISBN(isbn string) bool {
	if len(isbn) != 10 && !strings.Contains(isbn, "-") {
		return false
	}

	stripped := strings.ReplaceAll(isbn, "-", "")
	if len(stripped) != 10 {
		return false
	}

	var sum int
	for i, c := range stripped {
		var num int
		if i == 9 && c == 'X' {
			num = 10
		} else {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				return false
			}
			num = n
		}

		sum += (10 - i) * num
	}

	return sum%11 == 0
}
