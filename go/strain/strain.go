// Package strain provides Strain solution.
package strain

// Ints represents a slice of integers.
type Ints []int

// Keep returns a copy that keeps only values that returns true when used as an argument to f.
func (ints Ints) Keep(f func(int) bool) Ints {
	var result Ints
	for _, v := range ints {
		if !f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

// Discard returns a copy that discards all values that returns true when used as an argument to f.
func (ints Ints) Discard(f func(int) bool) Ints {
	var result Ints
	for _, v := range ints {
		if f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

// Lists represents a slice of slices of integers.
type Lists [][]int

// Keep returns a copy that keeps only values that returns true when used as an argument to f.
func (l Lists) Keep(f func([]int) bool) Lists {
	var result Lists
	for _, v := range l {
		if !f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

// Strings represents a slice of strings.
type Strings []string

// Keep returns a copy that keeps only values that returns true when used as an argument to f.
func (s Strings) Keep(f func(string) bool) Strings {
	var result Strings
	for _, v := range s {
		if !f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}
