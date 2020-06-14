// Package cipher is solution for problem Simple Cipher.
package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

// Caesar represents a caesar cipher.
type Caesar struct {
	shift rune
}

// NewCaesar returns new Caesar with fixed distance.
func NewCaesar() Caesar {
	return Caesar{rune(3)}
}

// NewShift returns new Caesar with given distance.
func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return Caesar{rune(distance)}
}

// Encode encodes given string with simple cipher.
func (c Caesar) Encode(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) {
			continue
		}
		b.WriteRune(mod((r-'a'+c.shift), 26) + 'a')
	}
	return b.String()
}

// Decode decodes given simple cipher string.
func (c Caesar) Decode(s string) string {
	var b strings.Builder
	for _, r := range s {
		b.WriteRune(mod((r-'a'-c.shift), 26) + 'a')
	}
	return b.String()
}

// Vigenere represents a vigenere cipher.
type Vigenere struct {
	key []rune
}

var vigenereKeyRegex = regexp.MustCompile(`^(a+)$`)

// NewVigenere creates new Vigenere.
func NewVigenere(key string) Cipher {
	k := vigenereKeyRegex.ReplaceAllString(key, "")
	if k == "" {
		return nil
	}
	for _, r := range k {
		if !unicode.IsLetter(r) || !unicode.IsLower(r) {
			return nil
		}
	}
	return Vigenere{[]rune(key)}
}

func (v Vigenere) Encode(s string) string {
	var b strings.Builder
	var i int
	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) {
			continue
		}
		pos := i % len(v.key)
		shift := v.key[pos] - 'a'
		d := mod((r-'a'+shift), 26) + 'a'
		b.WriteRune(d)
		i++
	}
	return b.String()
}

func (v Vigenere) Decode(s string) string {
	var b strings.Builder
	for i, r := range strings.ToLower(s) {
		pos := i % len(v.key)
		shift := v.key[pos] - 'a'
		b.WriteRune(mod((r-'a'-shift), 26) + 'a')
	}
	return b.String()
}

func mod(a, b rune) rune {
	return (a%b + b) % b
}
