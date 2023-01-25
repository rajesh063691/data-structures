package main

import (
	"fmt"
	list "practice_codes/data-structures/doublylistusingtail/taildoublylist"
)

func main() {
	fmt.Print("Doubly Linked list using tail pointer\n")
	dl := new(list.DoublyList)
	dl.AddNodeAtLast(10)
	dl.AddNodeAtLast(11)
	dl.AddNodeAtLast(12)
	dl.AddNodeAtLast(13)
	dl.AddNodeAtLast(14)
	dl.AddNodeAtLast(15)
	dl.AddNodeAtLast(16)
	dl.AddNodeAtLast(17)
	dl.AddNodeAtLast(18)
	dl.AddNodeAtLast(19)
	dl.AddNodeAtLast(144)
	dl.AddNodeAtLast(145)
	dl.AddNodeAtLast(146)
	dl.AddNodeAtLast(147)

	dl.DisplayListItems()

	dl.ReverseDoublyList()

	dl.DisplayListItems()

}
