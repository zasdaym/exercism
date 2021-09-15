// Package hamming provides Hamming solution.
package hamming

import "errors"

// Distance calculates the Hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands have different length")
	}

	diff := 0
	src, target := []rune(a), []rune(b)
	for i := range src {
		if src[i] != target[i] {
			diff++
		}
	}
	return diff, nil
}
