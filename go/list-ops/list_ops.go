// Package listops is solution for problem List Ops.
package listops

// IntList is a list of integers.
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Length return total items in list.
func (intList IntList) Length() int {
	count := 0
	for range intList {
		count++
	}
	return count
}

// Reverse will return the list backwards.
func (intList IntList) Reverse() IntList {
	for i := len(intList)/2 - 1; i >= 0; i-- {
		opp := len(intList) - 1 - i
		intList[i], intList[opp] = intList[opp], intList[i]
	}
	return intList
}

// Append will return the list appended with the new list.
func (intList IntList) Append(newList IntList) IntList {
	return append(intList, newList...)
}

// Concat will return the list appended with the new lists.
func (intList IntList) Concat(newLists []IntList) IntList {
	result := intList
	for _, newList := range newLists {
		result = result.Append(newList)
	}
	return result
}

// Map will return the list applied with the given function.
func (intList IntList) Map(f unaryFunc) IntList {
	result := make(IntList, intList.Length())
	for i, item := range intList {
		result[i] = f(item)
	}
	return result
}

// Filter will return the list filtered with the given function.
func (intList IntList) Filter(f predFunc) IntList {
	result := IntList{}
	for _, item := range intList {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// Foldr will fold each item in list into the accumulator from the right using function(item, accumulator).
func (intList IntList) Foldr(f binFunc, initial int) int {
	length := intList.Length()
	if length == 0 {
		return initial
	}

	counter := initial
	for i := length - 1; i >= 0; i-- {
		counter = f(intList[i], counter)
	}
	return counter
}

// Foldl will fold each item in list into the accumulator from the left using function(accumulator, item).
func (intList IntList) Foldl(f binFunc, initial int) int {
	length := intList.Length()
	if length == 0 {
		return initial
	}

	counter := initial
	for _, item := range intList {
		counter = f(counter, item)
	}
	return counter
}
