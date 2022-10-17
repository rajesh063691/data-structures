package linkedsort

import "practice_codes/data-structures/linked_mergedsort/list"

func MergeSort(lt list.LinkedList) list.LinkedList {
	//base case
	if lt.Head == nil || lt.Head.Next == nil {
		return lt
	}

	// split list into 2 halves
	left_half, right_half := SplitList(lt)
	left := MergeSort(left_half)
	right := MergeSort(right_half)

	return SortedMerge(left, right)
}

// SplitList ...
func SplitList(lt list.LinkedList) (leftHalf, rightHalf list.LinkedList) {

	// if list has zero or one item
	if lt.Head == nil || lt.Head.Next == nil {
		leftHalf = lt
		rightHalf = list.LinkedList{}
		return
	}
	//split the list into two halves
	curr := lt.Head
	slow := lt.Head
	fast := lt.Head.Next

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}

	leftHalf.Head = curr
	rightHalf.Head = slow.Next
	slow.Next = nil

	return
}

// Merge ...
func SortedMerge(left, right list.LinkedList) (lt list.LinkedList) {
	lt = list.LinkedList{}
	// if both lists are empty
	if left.Head == nil && right.Head == nil {
		return
	}
	//if left list is empty
	if left.Head == nil {
		lt.Head = right.Head
		return
	}
	//if right list is empty
	if right.Head == nil {
		lt.Head = left.Head
		return
	}

	// adding dummy head
	lt.Size = 0
	dummy_node := new(list.Node)
	dummy_node.Data = 0
	lt.Head = dummy_node
	current := lt.Head

	for left.Head != nil && right.Head != nil {
		// data comparision
		if left.Head.Data < right.Head.Data {
			current.Next = left.Head
			left.Head = left.Head.Next
			lt.Size++
		} else {
			current.Next = right.Head
			right.Head = right.Head.Next
			lt.Size++
		}
		current = current.Next
	}

	// check if any items are left in the left lists
	for left.Head != nil {
		current.Next = left.Head
		left.Head = left.Head.Next
		current = current.Next
		lt.Size++
	}

	// check if any items are left in the right lists
	for right.Head != nil {
		current.Next = right.Head
		right.Head = right.Head.Next
		current = current.Next
		lt.Size++
	}

	// discard the dummy head
	lt.Head = lt.Head.Next
	return
}
