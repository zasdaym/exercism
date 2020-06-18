// Package secret is solution for problem Secret Handshake.
package secret

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake returns secret handshake events from given code.
func Handshake(code uint) []string {
	var result []string
	for i := 0; i < len(actions); i++ {
		if code%2 == 1 {
			result = append(result, actions[i])
		}

		code /= 2
	}

	if code%2 == 1 {
		result = reverse(result)
	}

	return result
}

func reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}
