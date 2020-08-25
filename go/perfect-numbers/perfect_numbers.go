// Package perfect is solution for problem Perfect Numbers.
package perfect

import (
	"fmt"
	"math"
)

// Classification is number classification based by its factors sum.
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = fmt.Errorf("input must be a positive number")

// Classify classifies given integer with its factors sum.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return ClassificationDeficient, ErrOnlyPositive
	}

	var sum int64
	limit := int64(math.Ceil(float64(n / 2)))
	for i := int64(1); i <= limit; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	var c Classification
	switch {
	case sum == n:
		c = ClassificationPerfect
	case sum > n:
		c = ClassificationAbundant
	case sum < n:
		c = ClassificationDeficient
	}

	return c, nil
}
