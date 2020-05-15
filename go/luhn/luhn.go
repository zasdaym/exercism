// Package luhn is solution for problem Luhn.
package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if given numbers is valid per the Luhn formula.
func Valid(numbers string) bool {
	// convert number string to slice of runes
	runes := []rune(strings.ReplaceAll(numbers, " ", ""))
	// edge case
	if len(runes) <= 1 {
		return false
	}

	total := 0
	toDouble := true

	// iterate from back to front
	for i := len(runes) - 1; i >= 0; i-- {
		// second number will be doubled flag
		toDouble = !toDouble
		// if not number, false immediately
		if !unicode.IsNumber(runes[i]) {
			return false
		}

		num := int(runes[i] - '0')
		// double every two number
		if toDouble {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		// add to total
		total += num
	}
	// if total divisible by 10, then true
	return total%10 == 0
}
