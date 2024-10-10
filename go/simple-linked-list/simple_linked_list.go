package linkedlist

import (
	"fmt"
	"slices"
)

type List struct {
	element *Element
	size    int
}

type Element struct {
	value int
	next  *Element
}

func New(elements []int) *List {
	var (
		head    = &Element{}
		current = head
		size    = 0
	)
	for _, element := range elements {
		current.next = &Element{value: element}
		current = current.next
		size++
	}
	return &List{
		element: head.next,
		size:    size,
	}
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	if l.size == 0 {
		l.element = &Element{value: element}
		l.size += 1
		return
	}

	current := l.element
	for current.next != nil {
		current = current.next
	}
	current.next = &Element{value: element}
	l.size += 1
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, fmt.Errorf("unable to pop from an empty list")
	}
	if l.size == 1 {
		result := l.element.value
		l.element = nil
		l.size -= 1
		return result, nil
	}

	current := l.element
	for current.next != nil && current.next.next != nil {
		current = current.next
	}
	result := current.next.value
	current.next = nil
	l.size -= 1
	return result, nil
}

func (l *List) Array() []int {
	result := make([]int, 0, l.size)
	current := l.element
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}

func (l *List) Reverse() *List {
	elements := l.Array()
	reversed := make([]int, len(elements))
	copy(reversed, elements)
	slices.Reverse(reversed)
	return New(reversed)
}
