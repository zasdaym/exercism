// Package etl provides ETL solution.
package etl

import "strings"

// Transform transforms given list of letters per score to list of score per letter.
func Transform(in map[int][]string) map[string]int {
	scores := make(map[string]int)
	for score, chars := range in {
		for _, c := range chars {
			scores[strings.ToLower(c)] = score
		}
	}
	return scores
}
