// Package wordcount provides Word Count solution.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is word frequency in a phrase.
type Frequency map[string]int

var wordRegex = regexp.MustCompile(`\b[\w']+\b`)

// WordCount returns the word frequency of s.
func WordCount(s string) Frequency {
	words := wordRegex.FindAllString(s, -1)
	result := make(Frequency)
	for _, word := range words {
		result[strings.ToLower(word)]++
	}
	return result
}
