// Package prime provides Nth Prime solution.
package prime

import (
	"math"
)

// Nth returns the n-th prime number.
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	count, candidate := 0, 1
	for count < n {
		candidate++
		if isPrime(candidate) {
			count++
		}
	}

	return candidate, true
}

// isPrime determine whether n is a prime number.
func isPrime(n int) bool {
	if n != 2 && n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
