// Package armstrong is solution for problem Armstrong Numbers.
package armstrong

// IsNumber checks if a number is an Armstrong numbers.
func IsNumber(n int) bool {
	digits := intLen(n)
	original, sum := n, 0
	for n > 0 {
		lastDigit := n % 10
		sum += intPow(lastDigit, digits)
		n /= 10
	}
	return sum == original
}

// intLen finds integer digits.
func intLen(n int) int {
	digits := 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits
}

// intPow returns the powers of given base and exponent.
func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
