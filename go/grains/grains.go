// Package grains provides Grains solution.
package grains

import "errors"

// Square returns the number of grains on n-th square.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n must be between 1 and 64")
	}
	return 1 << (n - 1), nil
}

// Total returns total number of grains on all squares.
func Total() uint64 {
	return ^uint64(0)
}
