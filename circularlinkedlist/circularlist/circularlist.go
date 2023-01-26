package circularlist

import (
	"fmt"
	"strings"
)

// Node - actual node to store data and next and prev pointers
type Node struct {
	data int
	next *Node
}

// CircularList ...
type CircularList struct {
	size int
	head *Node
}

// CreateNewNode - creates new node with given data
func CreateNewNode(item int) (newNode *Node) {
	newNode = new(Node)
	newNode.data = item
	newNode.next = nil
	return
}
func (cl *CircularList) AddNodeAtLast(item int) {
	curr := cl.head
	newNode := CreateNewNode(item)
	// if head is nil, then list is empty. so first node will be the head node
	if curr == nil {
		curr = newNode
		cl.head = curr
		// point head node to itself to make the list circular
		curr.next = cl.head
		cl.size++
		return
	}
	// iterate till the second last element, and curr next should not be head
	for curr.next != cl.head {
		curr = curr.next
	}

	curr.next = newNode
	// changes the last nodes i.e newnode pointer to head to become as circular list
	newNode.next = cl.head

	cl.size++

}

func (cl *CircularList) DisplayListItems() {
	curr := cl.head
	// If head is nil, then no item is in the list
	if curr == nil {
		fmt.Print("list is empty, nothing to display.\n")
		return
	}
	listDatas := ""
	// iterate over all the nodes and display the data
	for curr.next != cl.head {
		listDatas = listDatas + fmt.Sprintf("%d => ", curr.data)
		curr = curr.next
	}
	// print the last element
	listDatas = listDatas + fmt.Sprintf("%d => ", curr.data)

	trimmedList := strings.TrimSuffix(listDatas, " => ")
	fmt.Printf("list size = %d\n", cl.size)
	fmt.Println(trimmedList)
}
