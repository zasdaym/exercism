// Package proverb is solution for problem Proverb.
package proverb

import "fmt"

// Proverb will returns a relevant proverb given a list of words.
func Proverb(rhyme []string) []string {
	result := []string{}
	if len(rhyme) == 0 {
		return result
	}
	for i, _ := range rhyme {
		if i < len(rhyme)-1 {
			s := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
			result = append(result, s)
		}
	}
	s := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return append(result, s)
}
