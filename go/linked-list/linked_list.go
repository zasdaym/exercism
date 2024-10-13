package linkedlist

import "fmt"

type List struct {
	head *Node
}

type Node struct {
	Value any
	prev  *Node
	next  *Node
}

func NewList(elements ...interface{}) *List {
	dummyHead := &Node{}
	current := dummyHead

	for i, element := range elements {
		var prev *Node
		if i > 0 {
			prev = current
		}
		current.next = &Node{
			Value: element,
			prev:  prev,
		}
		current = current.next
	}

	return &List{
		head: dummyHead.next,
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newHead := &Node{
		Value: v,
		next:  l.head,
	}

	if l.head != nil {
		l.head.prev = newHead
	}

	l.head = newHead
}

func (l *List) Push(v interface{}) {
	last := l.Last()
	newNode := &Node{
		Value: v,
		prev:  last,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	last.next = newNode
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, fmt.Errorf("empty list")
	}

	result := l.head.Value

	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	return result, nil
}

func (l *List) Pop() (interface{}, error) {
	last := l.Last()
	if last == nil {
		return nil, fmt.Errorf("empty list")
	}

	result := last.Value

	if last.prev == nil {
		l.head = nil
		return result, nil
	}

	last.prev.next = nil
	return result, nil
}

func (l *List) Reverse() {
	last := l.Last()
	current := last
	for current != nil {
		current.next, current.prev, current = current.prev, current.next, current.prev
	}
	l.head = last
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	current := l.head
	for current != nil && current.next != nil {
		current = current.next
	}
	return current
}
