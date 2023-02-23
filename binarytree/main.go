package main

import (
	"fmt"
	bt "practice_codes/data-structures/binarytree/btree"
)

func main() {
	fmt.Println("Binary Tree!!")

	var root *bt.Node
	bt.TraverseInorder(root)
	root = bt.InsertNode(root, 30)
	root = bt.InsertNode(root, 20)
	root = bt.InsertNode(root, 40)
	root = bt.InsertNode(root, 10)
	root = bt.InsertNode(root, 21)
	root = bt.InsertNode(root, 35)
	root = bt.InsertNode(root, 42)
	root = bt.InsertNode(root, 5)
	root = bt.InsertNode(root, 11)
	root = bt.InsertNode(root, 31)
	root = bt.InsertNode(root, 36)
	root = bt.InsertNode(root, 41)
	root = bt.InsertNode(root, 44)
	bt.TraverseInorder(root)
	root = bt.DeleteNode(root, 40)
	fmt.Println()
	bt.TraverseInorder(root)

}
