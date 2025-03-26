package brackets

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

var openers = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
}

func Bracket(input string) bool {
	var unclosed []rune
	for _, char := range input {
		if opening, ok := pairs[char]; ok {
			if len(unclosed) == 0 || unclosed[len(unclosed)-1] != opening {
				return false
			}
			unclosed = unclosed[:len(unclosed)-1]
		} else if openers[char] {
			unclosed = append(unclosed, char)
		}
	}
	return len(unclosed) == 0
}
