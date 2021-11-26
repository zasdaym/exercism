// Package pythagorean provides Pythagorean Triplet.
package pythagorean

import (
	"math"
)

// Triplet represents a pythagorean triplet.
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var triplets []Triplet
	for i := min; i <= max; i++ {
		for j := i + 1; j <= max; j++ {
			k := int(math.Sqrt(float64(i*i + j*j)))
			if k > max {
				break
			}
			if k*k != i*i+j*j {
				continue
			}
			triplets = append(triplets, Triplet{i, j, k})
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.
func Sum(p int) []Triplet {
	var result []Triplet
	triplets := Range(1, p)
	for _, triplet := range triplets {
		if triplet[0]+triplet[1]+triplet[2] == p {
			result = append(result, triplet)
		}
	}
	return result
}
