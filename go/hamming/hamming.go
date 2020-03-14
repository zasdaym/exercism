// Package hamming is solutin for problem Hamming.
package hamming

import "errors"

// Distance will returns the distance between to DNA strands.
func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return count, errors.New("length not same")
	}

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
