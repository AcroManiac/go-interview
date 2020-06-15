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

func (t *Tree) Depth() int {
	return findDepth(t.root)
}

func findDepth(n *Node) int {
	if n == nil {
		return 0
	}

	leftDepth := findDepth(n.Left)
	rightDepth := findDepth(n.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
