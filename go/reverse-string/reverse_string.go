// Package reverse provides Reverse String solution.
package reverse

// Reverse reverses given input.
func Reverse(input string) string {
	runes := []rune(input)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
