package main

import (
	"practice_codes/data-structures/linkedlist/list"
)

func main() {
	//fmt.Println("Linked List")
	lt := new(list.LinkedList)
	lt.AppendInList(1)
	lt.AppendInList(2)
	//lt.AppendInList(3)
	lt.AppendInList(4)
	lt.AppendInList(5)
	lt.AppendInList(6)
	lt.PrintListItems()
	// lt.DeleteItemFromList(1)
	// lt.DeleteItemFromList(3)
	// lt.DeleteItemFromList(5)
	// lt.DeleteItemFromList(6)
	// lt.PrintListItems()

	//lt.PrependInList(1)
	// lt.PrependInList(2)
	// lt.PrependInList(3)
	// lt.PrependInList(4)
	// lt.PrependInList(5)
	// lt.DeleteItemFromList(1)
	// lt.DeleteItemFromList(3)
	// lt.DeleteItemFromList(5)
	//lt.PrintListItems()
	lt.AddAtSpecificIndex(2, 3)
	lt.AddAtSpecificIndex(6, 7)
	lt.AddAtSpecificIndex(7, 8)
	lt.AddAtSpecificIndex(8, 9)
	lt.AddAtSpecificIndex(9, 10)
	lt.AddAtSpecificIndex(0, 0)
	lt.AddAtSpecificIndex(0, -1)
	lt.AddAtSpecificIndex(0, -2)

	lt.PrintListItems()
}
