package bst

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func NewCompleteTree(data []int) *Tree {
	return &Tree{
		root: levelOrderBinaryTree(data, 0, len(data)),
	}
}

func levelOrderBinaryTree(data []int, start, size int) *Node {
	curr := &Node{
		Value: data[start],
	}

	left := 2*start + 1
	right := left + 1

	if left < size {
		curr.Left = levelOrderBinaryTree(data, left, size)
	}
	if right < size {
		curr.Right = levelOrderBinaryTree(data, right, size)
	}
	return curr
}
