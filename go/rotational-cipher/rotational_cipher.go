// Package rotationalcipher provides Rotational Cipher solution.
package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher encrypts plain by shifting each char as many times as the value of shiftKey.
func RotationalCipher(plain string, shiftKey int) string {
	var sb strings.Builder
	for _, c := range plain {
		if !unicode.IsLetter(c) {
			sb.WriteRune(c)
			continue
		}
		offset := 'a'
		if unicode.IsUpper(c) {
			offset = 'A'
		}
		sb.WriteRune((c-offset+rune(shiftKey))%26 + offset)

	}
	return sb.String()
}
