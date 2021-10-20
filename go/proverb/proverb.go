// Package proverb provides Proverb solution.
package proverb

import "fmt"

// Proverb generates a proverb from given rhyme.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	result := make([]string, len(rhyme))
	for i := range rhyme[:len(rhyme)-1] {
		result[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	result[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return result
}
