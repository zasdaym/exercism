// Package prime is solution for problem Prime Factors.
package prime

// Factors returns all prime factors of a number.
func Factors(n int64) []int64 {
	result := make([]int64, 0)
	for n != 1 {
		factor := Factor(n)
		result = append(result, factor)
		n /= factor
	}
	return result
}

// Factor returns smallest prime factor of a number except 1.
func Factor(n int64) int64 {
	i := int64(2)
	for n%i != 0 {
		i++
	}
	return i
}
