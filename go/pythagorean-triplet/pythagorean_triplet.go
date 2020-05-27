// Package pythagorean is solution for problem Pythagorean Triplet.
package pythagorean

import (
	"math"
)

// Triplet represents a triplet.
type Triplet [3]int

func (t Triplet) Sum() int {
	return t[0] + t[1] + t[2]
}

// Range finds all pythagorean triplets with sides in the range min to max inclusive.
func Range(min, max int) []Triplet {
	var result []Triplet
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			cSquare := a*a + b*b
			if !isSquare(cSquare) {
				continue
			}

			c := intSqrt(cSquare)
			if c > max {
				continue
			}
			result = append(result, Triplet{a, b, c})
		}
	}

	return result
}

// Sum finds all pythagorean triplets where the sum of all sides in a triplets is equal to given number.
func Sum(n int) []Triplet {
	var result []Triplet
	for _, triplet := range Range(1, n) {
		if triplet.Sum() == n {
			result = append(result, triplet)
		}	
	}
	return result
}

func isSquare(n int) bool {
	f := float64(n)
	sqrt := math.Sqrt(f)
	return sqrt-math.Floor(sqrt) == 0
}

func intSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}
