// Package acronym is solution for problem Acronym.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate will return the abbreviations of a string.
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
