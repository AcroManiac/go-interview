package main

import "github.com/pkg/errors"

func main() {
	l := &LinkedList{}
	l.Add("A", 0)
	l.Add("D", 1)
}

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(value interface{}, pos int) error {

	// Check for empty list
	if l.head == nil && pos == 0 {
		l.head = &Node{value: value}
		return nil
	}

	n, err := l.Search(pos)
	if err != nil {
		return errors.Wrap(err, "failed adding new value")
	}

	// Create and insert new node O(1)
	newNode := &Node{
		value: value,
		next:  n,
		prev:  n.prev,
	}
	n.prev = newNode

	return nil
}

// Search position in list O(n)
func (l *LinkedList) Search(pos int) (*Node, error) {
	iter := l.head
	var cur int
	for iter != nil {
		if cur == pos {
			return iter, nil
		}
		iter = iter.next
		cur++
	}
	return nil, errors.New("list position is out of range")
}

func (l *LinkedList) String() string {
	var s string
	iter := l.head
	for iter != nil {
		if len(s) != 0 {
			s += ", "
		}
		s += iter.value.(string)
		iter = iter.next
	}
	return s
}
