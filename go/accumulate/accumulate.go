// Package accumulate is solution for problem Accumulate.
package accumulate

// Accumulate will accumulate given strings with given function.
func Accumulate(s []string, f func(string) string) []string {
	var result []string
	for i := range s {
		result = append(result, f(s[i]))
	}
	return result
}
