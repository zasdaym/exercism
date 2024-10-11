package piglatin

import "strings"

func Sentence(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = translate(words[i])
	}
	return strings.Join(words, " ")
}

func translate(s string) string {
	if s == "" {
		return ""
	}

	if isVowel(s[0]) {
		return s + "ay"
	}

	if s[:2] == "xr" || s[:2] == "yt" {
		return s + "ay"
	}

	var quPosition int
	for i := range s[:len(s)-1] {
		if s[i:i+2] == "qu" {
			quPosition = i
			return s[quPosition+2:] + s[:quPosition+2] + "ay"
		}
	}

	var vowelPosition int
	for i := range s {
		if isVowel(s[i]) || s[i] == 'y' && i > 0 {
			vowelPosition = i
			break
		}
		continue
	}

	return s[vowelPosition:] + s[:vowelPosition] + "ay"
}

func isVowel(c byte) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		if c == byte(vowel) {
			return true
		}
	}
	return false
}
