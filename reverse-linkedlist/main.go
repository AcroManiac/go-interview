package main

import "fmt"

func main() {
	// Create and populate list
	l := NewList()
	l.Add(42)
	l.Add(11)
	l.Add(72)
	l.Add(256)
	fmt.Printf("Source list: %s\n", l)

	// Reverse list iterative
	l.IterativeReverse()
	fmt.Printf("Reversed list: %s\n", l)

	// Reverse list recursive
	l.RecursiveReverse(l.Head)
	fmt.Printf("Reversed list: %s\n", l)
}
