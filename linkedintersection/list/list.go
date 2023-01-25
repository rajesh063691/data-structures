package list

import (
	"fmt"
	"strings"
)

type Node struct {
	Data int
	Next *Node
}
type LinkedList struct {
	Size int
	Head *Node
}

// add the item at last
func (l *LinkedList) AppendInList(item int) {
	newNode := new(Node)
	newNode.Data = item

	curr := l.Head
	//list is empty
	if curr == nil {
		curr = newNode
		l.Head = curr
		l.Size++
		return
	}
	// list not empty, iterate till last node
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	l.Size++
}
func (l *LinkedList) PrintListItems() {
	Head := l.Head
	if Head == nil {
		fmt.Println("no item found - list is empty")
		return
	}
	listItems := ""
	for Head != nil {
		Data := fmt.Sprintf("%d -> ", Head.Data)
		listItems = listItems + Data
		Head = Head.Next
	}
	trimmedList := strings.TrimSuffix(listItems, " -> ")
	fmt.Printf("%s (List Size = %d)", trimmedList, l.Size)
	fmt.Println()
}
