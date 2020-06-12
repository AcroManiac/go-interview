package bst

import "fmt"

func (t *Tree) PreOrderTraversal() {
	printPreOrder(t.root)
	fmt.Println()
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value.(int))
	printPreOrder(n.Left)
	printPreOrder(n.Right)
}
