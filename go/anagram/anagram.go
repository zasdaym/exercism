// Package anagram is solution for problem Anagram.
package anagram

import (
	"strings"
	"unicode"
)

type charCount [26]int

// Detect returns sublist of anagrams of the given string.
func Detect(subject string, candidates []string) []string {
	subjectLower := strings.ToLower(subject)
	subjectCount := countChar(subjectLower)
	result := []string{}

	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}

		candidateLower := strings.ToLower(candidate)
		if subjectLower == candidateLower {
			continue
		}

		candidateCount := countChar(candidateLower)
		if subjectCount == candidateCount {
			result = append(result, candidate)
		}
	}

	return result
}

func countChar(s string) charCount {
	var result charCount
	for _, c := range s {
		if unicode.IsLetter(c) {
			result[c-'a']++
		}
	}

	return result
}
