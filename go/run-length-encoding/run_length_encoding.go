// Package encode is solution for problem Run Length Encoding.
package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode encodes given string with Run-length data compression.
func RunLengthEncode(s string) string {
	var result strings.Builder
	count := 1
	for i := range s {
		if i != len(s)-1 && s[i] == s[i+1] {
			count++
			continue
		}

		if count > 1 {
			result.WriteString(strconv.Itoa(count))
		}

		result.WriteByte(s[i])
		count = 1
	}

	return result.String()
}

// RunLengthDecode decodes given Run-length encoded string.
func RunLengthDecode(s string) string {
	var result strings.Builder
	var count string
	for _, r := range s {
		if unicode.IsDigit(r) {
			count += string(r)
			continue
		}

		if count == "" {
			count = "1"
		}

		repeat, err := strconv.Atoi(count)
		if err != nil {
			panic(err)
		}

		data := strings.Repeat(string(r), repeat)
		result.WriteString(data)

		count = ""
	}

	return result.String()
}
