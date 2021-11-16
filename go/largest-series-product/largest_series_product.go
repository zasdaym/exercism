// Package lsproduct provides Largest Series Product solution.
package lsproduct

import (
	"errors"
	"strconv"
)

// LargestSeriesProduct returns the largest product for a series of span from digits.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) == 0 && span == 0 {
		return 1, nil
	}

	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	nums := make([]int64, len(digits))
	for i, c := range digits {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, err
		}
		nums[i] = int64(n)
	}

	var (
		product      int64 = 1
		runningCount int   = 0
		max          int64 = 0
	)

	for i := range nums {
		if nums[i] == 0 {
			product = 1
			runningCount = 0
			continue
		}

		if runningCount >= span {
			product /= nums[i-span]
		}

		product *= nums[i]
		runningCount++
		if runningCount >= span && product > max {
			max = product
		}
	}

	return max, nil
}
