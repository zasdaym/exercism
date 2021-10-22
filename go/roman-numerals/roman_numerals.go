// Package romannumerals provides Roman Numerals solution.
package romannumerals

import (
	"errors"
	"strings"
)

// arabicRoman is a pair of arabic & roman number.
type arabicRoman struct {
	arabic int
	roman  string
}

var arabicRomans = []arabicRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral converts arabic number to roman numeral
func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("to be converted number must be between 1 and 3000")
	}

	var sb strings.Builder
	for arabic > 0 {
		for _, v := range arabicRomans {
			if arabic >= v.arabic {
				n := arabic / v.arabic
				arabic -= n * v.arabic
				sb.WriteString(strings.Repeat(v.roman, n))
			}
		}
	}

	return sb.String(), nil
}
