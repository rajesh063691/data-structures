package main

import (
	"fmt"
	"practice_codes/data-structures/basictree/bstree"
)

func main() {
	fmt.Println("Basic Tree!!")

	root := bstree.NewNode(1)
	child1 := bstree.NewNode(2)
	child2 := bstree.NewNode(3)
	child3 := bstree.NewNode(4)
	child4 := bstree.NewNode(5)

	root.AddChild(child1)
	root.AddChild(child2)
	root.AddChild(child3)
	root.AddChild(child4)

	root.DisplayTree()

}
