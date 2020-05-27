// Package wordcount is solution for problem Word Count.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency represents words frequencies in a sentence.
type Frequency map[string]int

// WordCount returns words frequencies of given string.
func WordCount(s string) Frequency {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})

	result := make(Frequency)
	for _, word := range words {
		result[toLowerWord(word)]++
	}

	return result
}

func toLowerWord(s string) string {
	return strings.ToLower(strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	}))
}
