package main

import (
	"fmt"
	"practice_codes/data-structures/doublylinkedlist/doublylist"
)

func main() {
	fmt.Printf("Doubly Linked List Example\n")
	dt := new(doublylist.DoublyList)
	// dt.AddNodeAtLast(10)
	// dt.AddNodeAtLast(20)
	// dt.AddNodeAtLast(30)
	// dt.AddNodeAtLast(40)
	// dt.AddNodeAtLast(50)

	// dt.AddNodeAtFirstPosition(10)
	// dt.AddNodeAtFirstPosition(20)
	// dt.AddNodeAtFirstPosition(30)
	// dt.AddNodeAtFirstPosition(40)
	// dt.AddNodeAtFirstPosition(50)

	dt.AddNodeAtLast(10)
	dt.AddNodeAtLast(20)
	dt.AddNodeAtLast(30)
	dt.AddNodeAtLast(40)
	dt.AddNodeAtLast(50)

	// dt.DisplayListItems()

	dt.AddNodeAtGivenPosition(25, 3)
	dt.AddNodeAtGivenPosition(26, 4)
	dt.AddNodeAtGivenPosition(27, 5)
	dt.AddNodeAtGivenPosition(28, 6)
	dt.AddNodeAtGivenPosition(29, 7)

	dt.DisplayListItems()

	dt.DeleteNodeFromGivenPosition(3)
	dt.DeleteNodeFromGivenPosition(4)
	dt.DeleteNodeFromGivenPosition(5)
	dt.DeleteNodeFromGivenPosition(6)
	//dt.DeleteNodeFromGivenPosition(7)

	dt.DisplayListItems()

}

//25,27,29,
