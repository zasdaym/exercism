// Package allyourbase is solution for problem All Your Base.
package allyourbase

import (
	"fmt"
	"math"
)

var (
	errInvalidInputBase  = fmt.Errorf("input base must be >= 2")
	errInvalidOutputBase = fmt.Errorf("output base must be >= 2")
	errInvalidDigit      = fmt.Errorf("all digits must satisfy 0 <= d < input base")
)

// ConvertToBase converts digits from a base to another base.
func ConvertToBase(inputBase int, digits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errInvalidInputBase
	}

	if outputBase < 2 {
		return nil, errInvalidOutputBase
	}

	if len(digits) == 0 {
		return []int{0}, nil
	}

	n := len(digits)
	var result []int
	var decimal int
	for _, digit := range digits {
		if digit < 0 || digit >= inputBase {
			return nil, errInvalidDigit
		}

		n--
		decimal += digit * intPow(inputBase, n)
	}

	pow := findOptimalPower(outputBase, decimal)
	fmt.Println(inputBase, digits, outputBase, decimal, pow)
	for pow >= 0 {
		divider := intPow(outputBase, pow)
		x := decimal / divider
		result = append(result, x)
		decimal -= divider * x
		pow--
	}

	return result, nil
}

func intPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func findOptimalPower(base, target int) int {
	if base > target {
		return 0
	}

	n, pow := 1, 0
	for n < target {
		n *= base
		pow++
	}

	return pow - 1
}
