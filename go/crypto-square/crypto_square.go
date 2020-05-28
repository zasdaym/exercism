// Package cryptosquare is solution for problem Crypto Square.
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode encodes given string with square encoding.
func Encode(input string) string {
	s := strings.ToLower(input)
	sanitized := strings.Map(sanitize, s)

	row := int(math.Ceil(math.Sqrt(float64(len(sanitized)))))

	columns := make([]string, row)
	for i, r := range sanitized {
		columns[i%row] += string(r)
	}

	for i := range columns {
		if len(columns[i]) < len(columns[0]) {
			columns[i] += strings.Repeat(" ", len(columns[0])-len(columns[i]))
		}
	}

	return strings.Join(columns, " ")
}

func sanitize(r rune) rune {
	if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
		return -1
	}
	return r
}
