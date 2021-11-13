// Package anagram provides Anagram solution.
package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

// Detect returns the anagram of subject from candidates.
func Detect(subject string, candidates []string) []string {
	subjectCount := countLetters(subject)
	var matches []string

	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}

		count := countLetters(candidate)
		if reflect.DeepEqual(subjectCount, count) {
			matches = append(matches, candidate)
		}
	}

	return matches
}

// countLetters counts letters in a string.
func countLetters(s string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range s {
		count[unicode.ToLower(c)]++
	}
	return count
}
