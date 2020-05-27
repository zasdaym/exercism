// Package summultiples is solution for problem Sum Of Multiples.
package summultiples

// SumMultiples returns the sum of all the unique multiples of given numbers. up to but not including that number.
func SumMultiples(limit int, divisors ...int) int {
	var result int
	checker := make(map[int]bool)
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}

		x := divisor
		for x < limit {
			if !checker[x] {
				result += x
				checker[x] = true
			}
			x += divisor
		}
	}

	return result
}
