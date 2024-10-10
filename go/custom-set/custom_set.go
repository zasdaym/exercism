package stringset

import (
	"strings"
)

type Set struct {
	elements map[string]bool
}

func New() Set {
	return Set{
		elements: make(map[string]bool),
	}
}

func NewFromSlice(l []string) Set {
	elements := make(map[string]bool)
	for _, v := range l {
		elements[v] = true
	}
	return Set{
		elements: elements,
	}
}

func (s Set) String() string {
	var builder strings.Builder
	var i int
	for element := range s.elements {
		builder.WriteString("\"" + element + "\"")
		if len(s.elements) > 1 && i != len(s.elements)-1 {
			builder.WriteString(", ")
		}
		i++
	}
	return "{" + builder.String() + "}"
}

func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.elements[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.elements[elem] = true
}

func Subset(s1, s2 Set) bool {
	if len(s1.elements) > len(s2.elements) {
		return false
	}
	for element := range s1.elements {
		if _, ok := s2.elements[element]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for element := range s1.elements {
		if _, ok := s2.elements[element]; ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.elements) != len(s2.elements) {
		return false
	}
	for element := range s1.elements {
		if _, ok := s2.elements[element]; !ok {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	for element := range s1.elements {
		if _, ok := s2.elements[element]; !ok {
			delete(s1.elements, element)
		}
	}
	return s1
}

func Difference(s1, s2 Set) Set {
	for element := range s2.elements {
		delete(s1.elements, element)
	}
	return s1
}

func Union(s1, s2 Set) Set {
	for element := range s2.elements {
		s1.elements[element] = true
	}
	return s1
}
