package list

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	size int
	head *Node
}

// add the item at last
func (l *LinkedList) AppendInList(item int) {
	newNode := new(Node)
	newNode.data = item

	curr := l.head
	//list is empty
	if curr == nil {
		curr = newNode
		l.head = curr
		l.size++
		return
	}
	// list not empty, iterate till last node
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	l.size++
}
func (l *LinkedList) DeleteItemFromList(item int) {
	curr := l.head
	if curr == nil {
		return
	}

	// item found at head location
	if curr.data == item {
		l.head = curr.next
		l.size--
		return
	}

	for curr.next != nil {
		if curr.next.data == item {
			curr.next = curr.next.next
			l.size--
		} else {
			curr = curr.next
		}
	}
}

func (l *LinkedList) PrintListItems() {
	head := l.head
	if head == nil {
		fmt.Println("no item found - list is empty")
		return
	}
	listItems := ""
	for head != nil {
		data := fmt.Sprintf("%d -> ", head.data)
		listItems = listItems + data
		head = head.next
	}
	trimmedList := strings.TrimSuffix(listItems, " -> ")
	fmt.Printf("%s (List size = %d)", trimmedList, l.size)
	fmt.Println()
}

// add the item at the begining
func (l *LinkedList) PrependInList(item int) {
	newNode := new(Node)
	newNode.data = item
	curr := l.head
	newNode.next = curr
	l.head = newNode

}

// add element at a specific index
func (l *LinkedList) AddAtSpecificIndex(index, item int) {
	prev := l.head
	pos := 0
	newNode := new(Node)
	newNode.data = item
	//add at begining
	if index == pos {
		newNode.next = prev
		l.head = newNode
		l.size++
		return
	}
	for prev.next != nil {
		pos = pos + 1
		if index == pos {
			newNode.next = prev.next
			prev.next = newNode
			l.size++
			return
		}
		prev = prev.next
	}
	// finally, if it's last element
	pos = pos + 1
	if index == pos {
		prev.next = newNode
		l.size++
	}
}
