package btree

import "fmt"

// Node ...
type Node struct {
	data  int
	left  *Node
	right *Node
}

// CreateNode ...
func CreateNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

// InsertNode ...
func InsertNode(root *Node, data int) *Node {
	if root == nil {
		newNode := CreateNode(data)
		return newNode
	}

	if data < root.data {
		root.left = InsertNode(root.left, data)
	} else if data > root.data {
		root.right = InsertNode(root.right, data)
	}

	return root
}

// traverse Inorder ...
func TraverseInorder(root *Node) {
	if root != nil {
		TraverseInorder(root.left)
		fmt.Printf("%d ", root.data)
		TraverseInorder(root.right)
	}
}
