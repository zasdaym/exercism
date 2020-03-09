// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"unicode"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if isNothing(remark) {
		return "Fine. Be that way!"
	} else if isUpper(remark) {
		if isQuestion(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else if isQuestion(remark) {
		return "Sure."
	} else {
		return "Whatever."
	}
}

func isNothing(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func isUpper(s string) bool {
	foundLetter := false
	for _, x := range s {
		if !foundLetter && unicode.IsLetter(x) {
			foundLetter = true
		}

		if !unicode.IsUpper(x) && unicode.IsLetter(x) {
			return false
		}
	}
	return true && foundLetter
}

func isQuestion(s string) bool {
	str := strings.TrimSpace(s)
	return str[len(str)-1:] == "?"
}
