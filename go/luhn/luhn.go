// Package luhn is solution for problem Luhn.
package luhn

import (
	"strings"
	"strconv"
)

// Valid determines if given numbers is valid per the Luhn formula.
func Valid(numbers string) bool {
	runes := []rune(strings.ReplaceAll(numbers, " ", ""))
	if len(runes) <= 1 {
		return false
	}

	total := 0
	toDouble := len(runes)%2 == 0

	for _, r := range runes {
		num, err := strconv.Atoi(string(r))
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
