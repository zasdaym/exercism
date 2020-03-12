// Package darts is solution for problem Darts.
package darts

import "math"

// Score will returns the score given the float coordinate of a dart.
func Score(x, y float64) int {
	var score int

	radius := math.Sqrt(x*x + y*y)

	switch {
	case radius <= 1:
		score = 10
	case radius <= 5:
		score = 5
	case radius <= 10:
		score = 1
	default:
		score = 0
	}
	return score
}
