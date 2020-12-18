// Package brackets is solution for problem Matching Brackets.
package brackets

var (
	pairs = map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	bracket = map[byte]bool{
		'{': true,
		'(': true,
		'[': true,
		'}': true,
		')': true,
		']': true,
	}
)

// Bracket validates all bracket pairs in a string are matched and nested correctly.
func Bracket(s string) bool {
	var stack []byte
	for i := range s {
		if _, ok := bracket[s[i]]; !ok {
			continue
		}
		if val, ok := pairs[s[i]]; ok {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if val == last {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
