// Package triangle is solution for problem Triangle
package triangle

import "sort"

// Kind is the type of a triangle.
type Kind int

const (
	NaT Kind = 0
	Equ Kind = 1
	Iso Kind = 2
	Sca Kind = 3
)

// KindFromSides will returns the type of a triangle given three sides of the triangle.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

func isTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	sides := []float64{a, b, c}
	sort.Float64s(sides)
	return sides[2]-sides[1] <= sides[0]
}
