package main

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head *Node
}

func NewList() *List {
	return &List{}
}

// Get tail
func (l *List) Tail() *Node {
	var tail *Node
	p := l.Head
	for p != nil {
		if p.Next == nil {
			tail = p
		}
		p = p.Next
	}
	return tail
}

// Add value to list tail
func (l *List) Add(v interface{}) {
	newNode := &Node{Value: v}
	tail := l.Tail()

	// Check for empty list
	if tail == nil {
		l.Head = newNode
		return
	}
	tail.Next = newNode
}

func (l *List) String() string {
	var s string
	for p := l.Head; p != nil; p = p.Next {
		s += fmt.Sprintf("%d", p.Value)
		if p.Next != nil {
			s += " -> "
		}
	}
	return s
}

func (l *List) IterativeReverse() {
	var (
		curr *Node
		prev *Node
		p    *Node = l.Head
	)

	for p != nil {
		curr = p
		p = p.Next

		// Reverse pointer
		curr.Next = prev
		prev = curr
		l.Head = curr
	}
}

func (l *List) RecursiveReverse(n *Node) *Node {

	// Check for tail
	if n.Next == nil {
		return n
	}
	l.Head = l.RecursiveReverse(n.Next)

	// Reverse links
	n.Next.Next = n
	n.Next = nil

	return l.Head
}
