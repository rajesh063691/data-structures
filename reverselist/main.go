package main

import (
	"practice_codes/data-structures/reverselist/list"
)

func main() {
	lt := list.LinkedList{}
	lt.AppendInList(9)
	lt.AppendInList(8)
	lt.AppendInList(13)
	lt.AppendInList(14)
	lt.AppendInList(5)
	lt.AppendInList(2)
	lt.AppendInList(0)
	lt.AppendInList(1)
	lt.PrintListItems()

	revList := ReverseList(lt)
	revList.PrintListItems()

}

func ReverseList(lst list.LinkedList) (lt list.LinkedList) {
	lt = list.LinkedList{}
	//return nil if list is empty and return head if list has only one element
	if lst.Head == nil {
		return
	} else if lst.Head.Next == nil {
		return lst
	}

	var prev *list.Node
	curr := lst.Head
	next := lst.Head.Next
	for next != nil {
		curr.Next = prev
		prev = curr
		curr = next
		next = next.Next

		lt.Size++
	}
	// make curr next to prev and make curr as new head
	curr.Next = prev
	lt.Size++
	lt.Head = curr
	return
}
