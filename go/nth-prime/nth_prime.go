// Package prime is solution for problem Nth Prime.
package prime

import (
	"math"
)

// Nth returns the n-th prime number of given integer.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	for p := 2; ; p++ {
		if !isPrime(p) {
			continue
		}

		n--
		if n == 0 {
			return p, true
		}
	}
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return n == 2
	}

	m := int(math.Sqrt(float64(n))) + 1
	for i := 3; i < m; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
