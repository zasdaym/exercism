// Package piglatin is solution for problem Pig Latin.
package piglatin

import (
	"regexp"
	"strings"
)

var (
	reBeginWithVowelSound = regexp.MustCompile(`^(xr|yt|[aeiou])`)
	reBeginWithConsSound  = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]{1,3})`)
)

// Sentence converts a string to its Pig Latin version.
func Sentence(s string) string {
	var sb strings.Builder
	for _, word := range strings.Fields(s) {
		switch {
		case reBeginWithVowelSound.MatchString(word):
			sb.WriteString(word + "ay ")
		default:
			cons := reBeginWithConsSound.FindStringSubmatch(word)[1]
			tail := word[len(cons):]

			// if length of consonant is 2 or longer AND ends with y, prepend the y to the tail
			if len := len(cons); len >= 2 && cons[len-1] == 'y' {
				cons = cons[0 : len-1]
				tail = "y" + tail
			}
			sb.WriteString(tail + cons + "ay ")
		}
	}
	return strings.TrimSpace(sb.String())
}
