//Package isogram is solution for problem Isogram
package isogram

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile("[^a-zA-Z]+")

// IsIsogram will check if given string is an isogram
func IsIsogram(s string) bool {
	letters := regex.ReplaceAllString(strings.ToLower(s), "")
	checked := map[rune]bool{}
	for _, letter := range letters {
		if _, ok := checked[letter]; ok {
			return false
		}

		checked[letter] = true
	}

	return true
}
