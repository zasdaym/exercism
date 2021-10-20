// Package triangle provides Triangle solution.
package triangle

import "math"

// Kind represents a triangle type.
type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides determines a triangle type given the triangle sides.
func KindFromSides(a, b, c float64) Kind {
	if math.IsInf(a+b+c, 0) || math.IsNaN(a+b+c) || a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
