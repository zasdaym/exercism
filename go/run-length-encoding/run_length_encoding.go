// Package encode provides Run Length Encoding solution.
package encode

import (
	"fmt"
	"strconv"
	"strings"
)

// RunLengthEncode encodes input using run-length encoding.
func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	lastChar, count := input[0], 1
	var sb strings.Builder

	for i := 1; i < len(input); i++ {
		if input[i] == lastChar {
			count++
			continue
		}

		sb.WriteString(encodeCharCount(lastChar, count))
		lastChar, count = input[i], 1
	}
	sb.WriteString(encodeCharCount(lastChar, count))

	return sb.String()
}

// encodeCharCount encodes a char and its count into a string.
func encodeCharCount(c byte, count int) string {
	if count == 1 {
		return fmt.Sprintf("%c", c)
	} else {
		return fmt.Sprintf("%d%c", count, c)
	}
}

// RunLengthDecode decodes a run-length encoded string.
func RunLengthDecode(input string) string {
	count := 1
	var left int
	var sb strings.Builder

	for i := 0; i < len(input); i++ {
		if n, err := strconv.Atoi(input[left : i+1]); err == nil {
			count = n
			continue
		}

		sb.WriteString(strings.Repeat(string(input[i]), count))
		count, left = 1, i+1
	}
	return sb.String()
}
