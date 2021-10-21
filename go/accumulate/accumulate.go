// Package accumulate provides Accumulate solution.
package accumulate

// Accumulate accumulates each element in words with given converter.
func Accumulate(words []string, convert func(string) string) []string {
	accumulated := make([]string, len(words))
	for i := range words {
		accumulated[i] = convert(words[i])
	}
	return accumulated
}
