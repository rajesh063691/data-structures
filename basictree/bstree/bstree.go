package bstree

import "fmt"

// Node ...
type Node struct {
	Value    int
	Children []*Node
}

// NewNode ...
func NewNode(data int) *Node {
	return &Node{Value: data}
}

// Add Child Node
func (n *Node) AddChild(child *Node) {
	if child == nil {
		fmt.Print("child is nil, can not append\n")
		return
	}
	n.Children = append(n.Children, child)
}

func (n *Node) DisplayTree() {
	if n == nil {
		fmt.Print("no child to display\n")
		return
	}

	fmt.Printf("%d ", n.Value)

	for _, child := range n.Children {
		child.DisplayTree()
	}
}
