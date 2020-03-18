// Package etl is solution for problem ETL.
package etl

import "strings"

// Transform will accept letter per score list, and returns score per letter list.
func Transform(scoreList map[int][]string) map[string]int {
	result := map[string]int{}
	for score, letters := range scoreList {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}
