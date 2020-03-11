// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var result []string

	regex := regexp.MustCompile("[-_ ]+")
	str := regex.ReplaceAllString(s, " ")
	words := strings.Split(str, " ")

	for _, word := range words {
		firstChar := strings.ToUpper(word[:1])
		result = append(result, firstChar)
	}

	return strings.Join(result, "")
}
