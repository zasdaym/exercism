// Package cipher provides Simple Cipher implementation.
package cipher

import (
	"strings"
	"unicode"
)

// normalize removes all non-letter characters from s and convert all letters to lowercase.
func normalize(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		sb.WriteRune(unicode.ToLower(c))
	}
	return sb.String()
}

// shiftLetter shifts c by given distance.
func shiftLetter(c rune, distance int) rune {
	return (c-'a'+rune(distance)+26)%26 + 'a'
}

// shift represents Shift Cipher implementation.
type shift struct {
	distance int
}

// NewShift creates a new shift.
func NewShift(distance int) *shift {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return &shift{
		distance: (distance + 26) % 26,
	}
}

// NewCaesar creates a new shift with fixed distance of 3.
func NewCaesar() *shift {
	return NewShift(3)
}

// Encode encodes input using Shift Cipher.
func (s shift) Encode(input string) string {
	var sb strings.Builder
	for _, c := range normalize(input) {
		sb.WriteRune(shiftLetter(c, s.distance))
	}
	return sb.String()
}

// Decode decodes input using Shift Cipher.
func (s shift) Decode(input string) string {
	var sb strings.Builder
	for _, c := range normalize(input) {
		sb.WriteRune(shiftLetter(c, -s.distance))
	}
	return sb.String()
}

// vigenere represents a Vigener Cipher implementation.
type vigenere struct {
	key string
}

// NewVigenere creates a new vigenere.
func NewVigenere(key string) *vigenere {
	if key == "" {
		return nil
	}
	var complexKey bool
	for _, c := range key {
		if !unicode.IsLetter(c) || unicode.IsUpper(c) {
			return nil
		}
		if c != 'a' {
			complexKey = true
		}
	}
	if !complexKey {
		return nil
	}
	return &vigenere{
		key: key,
	}
}

// Encode encodes input using Vigener Cipher.
func (v vigenere) Encode(input string) string {
	var sb strings.Builder
	normalized := normalize(input)
	for i, c := range normalized {
		shift := v.key[i%len(v.key)] - 'a'
		sb.WriteRune(shiftLetter(c, int(shift)))
	}
	return sb.String()
}

// Decode decodes input using Vigener Cipher.
func (v vigenere) Decode(input string) string {
	var sb strings.Builder
	normalized := normalize(input)
	for i, c := range normalized {
		shift := v.key[i%len(v.key)] - 'a'
		sb.WriteRune(shiftLetter(c, -int(shift)))
	}
	return sb.String()
}
