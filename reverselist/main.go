package main

import (
	"practice_codes/data-structures/reverselist/list"
	"practice_codes/data-structures/reverselist/reverse"
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

	revList := reverse.ReverseList(lt)
	revList.PrintListItems()

}
