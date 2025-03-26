package transpose

import (
	"slices"
)

func Transpose(input []string) []string {
	var length int
	for _, line := range input {
		if len(line) > length {
			length = len(line)
		}
	}

	var result []string
	for col := 0; col < length; col++ {
		var shouldPad bool
		chars := make([]byte, 0, length)
		for row := len(input) - 1; row >= 0; row-- {
			if col < len(input[row]) {
				chars = append(chars, input[row][col])
				shouldPad = true
				continue
			}
			if shouldPad {
				chars = append(chars, ' ')
			}
		}
		slices.Reverse(chars)
		result = append(result, string(chars))
	}
	return result
}
