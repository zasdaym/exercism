// Package atbash provides Atbash Cipher solution.
package atbash

import (
	"strings"
	"unicode"
)

// Atbash encrypts s with Atbash Cipher.
func Atbash(s string) string {
	var sb strings.Builder
	var count int
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			continue
		}
		sb.WriteRune(encode(unicode.ToLower(c)))
		count++
		if count%5 == 0 {
			sb.WriteRune(' ')
		}
	}
	return strings.TrimSpace(sb.String())
}

// encode encodes a letter with Atbash Cipher.
func encode(r rune) rune {
	if unicode.IsNumber(r) {
		return r
	}
	return 25 - (r - 'a') + 'a'
}
