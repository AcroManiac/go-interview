package main

import "github.com/AcroManiac/go-interview/binary-search-tree/bst"

func main() {
	t := bst.NewCompleteTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	t.PreOrderTraversal()
}
