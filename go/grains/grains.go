// Package grains is solution for problem Grains.
package grains

import (
	"errors"
)

// Square calculates the number of grains on a given square number.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid square number")
	}

	return 1 << (n - 1), nil
}

// Total returns the total of grains on a chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
