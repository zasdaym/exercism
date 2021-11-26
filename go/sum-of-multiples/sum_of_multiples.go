// Package summultiples provides Sum of Multiples solution.
package summultiples

// SumMultiples returns the sum of all numbers lower than limit that is divisible by any divisors.
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i >= divisor && i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
