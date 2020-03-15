// Package raindrops is solution for problem Raindrops.
package raindrops

import (
	"fmt"
	"strconv"
)

// Convert will convert given number to raindrop sounds.
func Convert(n int) string {
	result := ""

	if n%3 == 0 {
		result = fmt.Sprintf("%s%s", result, "Pling")
	}

	if n%5 == 0 {
		result = fmt.Sprintf("%s%s", result, "Plang")
	}

	if n%7 == 0 {
		result = fmt.Sprintf("%s%s", result, "Plong")
	}

	if result == "" {
		return strconv.FormatInt(int64(n), 10)
	}

	return result
}
