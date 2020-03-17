// Package strain is solution for problem Strain.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep will returns a new Ints containing elements where the predicate is true.
func (ints Ints) Keep(f func(int) bool) Ints {
	var result Ints
	for i := range ints {
		if f(ints[i]) {
			result = append(result, ints[i])
		}
	}
	return result
}

// Discard will returns a new Ints containing elements where the predicate is false.
func (ints Ints) Discard(f func(int) bool) Ints {
	var result Ints
	for i := range ints {
		if !f(ints[i]) {
			result = append(result, ints[i])
		}
	}
	return result
}

// Keep will returns a new Lists containing elements where the predicate is true.
func (lists Lists) Keep(f func([]int) bool) Lists {
	var result Lists
	for _, list := range lists {
		if f(list) {
			result = append(result, list)
		}
	}
	return result
}

// Keep will returns a new Strings containing elements where the predicate is true.
func (strings Strings) Keep(f func(string) bool) Strings {
	var result Strings
	for _, s := range strings {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}
