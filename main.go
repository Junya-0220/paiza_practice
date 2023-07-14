package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	}

	if value < node.value {
		node.left = insert(node.left, value)
	} else {
		node.right = insert(node.right, value)
	}
	return node
}

func main() {
	var root *Node
	root = insert(root, 3)
	root = insert(root, 6)
	root = insert(root, 5)
	root = insert(root, 7)
	root = insert(root, 1)
	root = insert(root, 10)
	root = insert(root, 2)
	root = insert(root, 8)
	root = insert(root, 700)
	fmt.Println(root.left.value, root.right.value)
}
