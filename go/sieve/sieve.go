// Package sieve provides Sieve solution.
package sieve

// Sieve returns all prime numbers between two and the given limit.
func Sieve(limit int) []int {
	if limit < 2 {
		return nil
	}
	composites := make([]bool, limit+1)
	var primes []int
	for i := 2; i <= limit; i++ {
		if composites[i] {
			continue
		}
		primes = append(primes, i)
		for j := i; j <= limit; j += i {
			composites[j] = true
		}
	}
	return primes
}
