// Package hamming is solutin for problem Hamming.
package hamming

import (
	"errors"
)

// Distance will returns the distance between to DNA strands.
func Distance(a, b string) (int, error) {
	newA := []rune(a)
	newB := []rune(b)
	if len(a) != len(b) {
		return 0, errors.New("length not same")
	}

	count := 0
	for i := range newA {
		if newA[i] != newB[i] {
			count++
		}
	}
	return count, nil
}
