package main

import (
	"practice_codes/data-structures/linked_mergedsort/linkedsort"
	"practice_codes/data-structures/linked_mergedsort/list"
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

	sorted_list := linkedsort.MergeSort(lt)
	sorted_list.PrintListItems()

}
