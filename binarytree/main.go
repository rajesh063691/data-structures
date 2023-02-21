package main

import (
	"fmt"
	bt "practice_codes/data-structures/binarytree/btree"
)

func main() {
	fmt.Println("Binary Tree!!")

	var root *bt.Node
	bt.TraverseInorder(root)
	root = bt.InsertNode(root, 5)
	root = bt.InsertNode(root, 4)
	root = bt.InsertNode(root, 6)
	root = bt.InsertNode(root, 3)
	root = bt.InsertNode(root, 7)
	root = bt.InsertNode(root, 2)
	root = bt.InsertNode(root, 8)
	root = bt.InsertNode(root, 9)
	bt.TraverseInorder(root)

}
