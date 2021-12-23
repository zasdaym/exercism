// Package cryptosquare provides Crypto Square solution.
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode encodes pt to square code.
func Encode(pt string) string {
	var (
		normalized = normalize(pt)
		c          = int(math.Ceil(math.Sqrt(float64(len(normalized)))))
		r          = c - 1
	)
	if c*r < len(normalized) {
		r++
	}
	cols := make([]string, c)
	for i := 0; i < c*r; i++ {
		s := " "
		if i < len(normalized) {
			s = string(normalized[i])
		}
		cols[i%c] += s
	}
	return strings.Join(cols, " ")
}

// normalize converts all letters in s to lowercase and removes all non-letter characters.
func normalize(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			continue
		}
		sb.WriteRune(unicode.ToLower(c))
	}
	return sb.String()
}
