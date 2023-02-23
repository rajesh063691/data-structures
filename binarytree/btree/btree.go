package btree

import "fmt"

// Node ...
type Node struct {
	value int
	left  *Node
	right *Node
}

// CreateNode ...
func CreateNode(value int) *Node {
	return &Node{
		value: value,
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

	if data < root.value {
		root.left = InsertNode(root.left, data)
	} else if data > root.value {
		root.right = InsertNode(root.right, data)
	}

	return root
}

// traverse Inorder ...
func TraverseInorder(root *Node) {
	if root != nil {
		TraverseInorder(root.left)
		fmt.Printf("%d ", root.value)
		TraverseInorder(root.right)
	}
}

// Delete nodes...
func DeleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	// Search for the node to be deleted
	if value < root.value {
		root.left = DeleteNode(root.left, value)
		return root
	} else if value > root.value {
		root.right = DeleteNode(root.right, value)
		return root
	}

	// If the node has no children
	if root.left == nil && root.right == nil {
		return nil
	}

	// If the node has one child
	if root.left == nil {
		return root.right
	} else if root.right == nil {
		return root.left
	}

	// If the node has two children
	minRight := root.right
	for minRight.left != nil {
		minRight = minRight.left
	}
	root.value = minRight.value
	root.right = DeleteNode(root.right, minRight.value)

	return root
}
