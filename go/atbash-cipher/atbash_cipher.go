// Package atbash is solution for problem Atbash Cipher.
package atbash

import (
	"strings"
	"unicode"
)

// Atbash returns given string encoded with atbash cipher.
func Atbash(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			continue
		}

		if unicode.IsLetter(r) {
			r = 'z' - r + 'a'
		}

		b.WriteRune(r)
		if (b.Len() + 1) % 6 == 0 {
			b.WriteRune(' ')
		}
	}
	return strings.TrimSpace(b.String())
}
