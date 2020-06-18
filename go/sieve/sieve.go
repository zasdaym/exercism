// Package sieve is solution for problem Sieve.
package sieve

// Sieve returns all prime numbers from 2 up to given limit inclusive.
func Sieve(limit int) []int {
	var result []int
	composites := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		if composites[i] {
			continue
		}

		for j := i; j <= limit; j += i {
			composites[j] = true
		}

		result = append(result, i)
	}
	return result
}
