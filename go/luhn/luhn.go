// Package luhn is solution for problem Luhn.
package luhn

import (
	"strings"
	"strconv"
)

// Valid determines if given numbers is valid per the Luhn formula.
func Valid(s string) bool {
	numbers := strings.ReplaceAll(s, " ", "")
	if len(numbers) <= 1 {
		return false
	}

	total := 0
	toDouble := len(numbers)%2 == 0

	for _, n := range numbers {
		num, err := strconv.Atoi(string(n))
		if err != nil {
			return false
		}

		if toDouble {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		total += num
		toDouble = !toDouble
	}

	return total%10 == 0
}
