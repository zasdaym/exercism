// Package lsproduct is solution for problem Largest Series Product.
package lsproduct

import (
	"fmt"
	"unicode"
)

// LargestSeriesProduct returns the largets series product for given length of given digits string.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 {
		return 1, nil
	}

	if span < 0 {
		return -1, fmt.Errorf("span must be greater than zero")
	}

	if len(digits) < span {
		return -1, fmt.Errorf("span must be smaller then string length")
	}

	var result int64
	for i := range digits {
		if i+span > len(digits) {
			break
		}

		var product int64 = 1
		for _, r := range digits[i : i+span] {
			if !unicode.IsNumber(r) {
				return -1, fmt.Errorf("digits must only contains digits")
			}
			product *= int64(r - '0')
		}

		if product > result {
			result = product
		}
	}

	return result, nil
}
