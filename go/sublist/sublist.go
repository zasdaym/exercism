// Package sublist is solution for problem Sublist.
package sublist

// Relation represents two list relation.
type Relation string

// Sublist determines relation of two list.
func Sublist(first, second []int) Relation {
	switch {
	case len(first) == len(second) && isSublist(first, second):
		return "equal"
	case len(first) < len(second) && isSublist(first, second):
		return "sublist"
	case len(first) > len(second) && isSublist(second, first):
		return "superlist"
	}
	return "unequal"
}

func isSublist(shorter, longer []int) bool {
	for i := 0; i < len(longer)-len(shorter)+1; i++ {
		match := true
		for j := range shorter {
			if shorter[j] != longer[i+j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// 1 2 3 4 5	(5)
// 2 3 4		(3)
