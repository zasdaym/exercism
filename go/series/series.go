// Package series is solution for problem Series.
package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var result []string
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n, will return false if substring length is bigger than length of s.
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
