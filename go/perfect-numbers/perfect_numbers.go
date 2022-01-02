// Package perfect provides Perfect Number solution.
package perfect

import "errors"

// Classification represents a number classification based on their aliquot sum.
type Classification string

const (
	ClassificationPerfect   = Classification("perfect")
	ClassificationAbundant  = Classification("abundant")
	ClassificationDeficient = Classification("deficient")
)

var ErrOnlyPositive = errors.New("n must be a positive number")

// Classify returns the number type based on their aliquot sum.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return "", ErrOnlyPositive
	}
	var aliquotSum int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			aliquotSum += i
		}
	}
	var result Classification
	switch {
	case aliquotSum == n:
		result = ClassificationPerfect
	case aliquotSum > n:
		result = ClassificationAbundant
	case aliquotSum < n:
		result = ClassificationDeficient
	}
	return result, nil
}
