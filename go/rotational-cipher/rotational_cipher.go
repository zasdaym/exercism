// Package rotationalcipher is solution for problem Rotational Cipher.
package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher encodes given string with Caesar Cipher.
func RotationalCipher(plain string, shift int) string {
	var b strings.Builder
	for _, r := range plain {
		if !unicode.IsLetter(r) {
			b.WriteRune(r)
			continue
		}

		var offset rune
		if r >= 'a' && r <= 'z' {
			offset = 'a'
		}

		if r >= 'A' && r <= 'Z' {
			offset = 'A'
		}

		e := (r-offset+rune(shift))%26 + offset
		b.WriteRune(e)
	}
	return b.String()
}
