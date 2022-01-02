// Package allyourbase provides All Your Base solution.
package allyourbase

import (
	"errors"
)

// ConvertToBase converts number from one base to another base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	decimal, multiplier := 0, 1
	for i := range inputDigits {
		digit := inputDigits[len(inputDigits)-1-i]
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decimal += multiplier * digit
		multiplier *= inputBase
	}

	divisor, resultLen := 1, 1
	for divisor*outputBase < decimal {
		divisor *= outputBase
		resultLen++
	}
	result := make([]int, resultLen)
	for i := range result {
		result[i] = decimal / divisor
		decimal %= divisor
		divisor /= outputBase
	}
	return result, nil
}
