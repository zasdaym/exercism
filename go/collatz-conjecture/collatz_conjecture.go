// Package collatzconjecture is solution for problem Collatz Conjecture.
package collatzconjecture

import (
	"errors"
)

// CollatzConjecture will return total steps of a collatz conjecture of a given integer
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("initial number must be equal or more than 1")
	}

	count := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		count++
	}
	return count, nil
}
