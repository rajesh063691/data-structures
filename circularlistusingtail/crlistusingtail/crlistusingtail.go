package crlistusingtail

import (
	"fmt"
	"strings"
)

// Node - actual node to store data and next and prev pointers
type Node struct {
	data int
	next *Node
}

// CircularTailList ...
type CircularTailList struct {
	size int
	head *Node
	tail *Node
}

// CreateNewNode - creates new node with given data
func CreateNewNode(item int) (newNode *Node) {
	newNode = new(Node)
	newNode.data = item
	newNode.next = nil
	return
}
func (cl *CircularTailList) AddNodeAtLast(item int) {
	curr := cl.head
	newNode := CreateNewNode(item)
	// if head is nil, then list is empty. so first node will be the head node
	if curr == nil {
		curr = newNode
		cl.head = curr
		cl.tail = curr
		// pointtail node next to head to make a list as circular
		cl.tail.next = cl.head
		cl.size++
		return
	}
	// insert newNode after tail node
	newNode.next = cl.tail.next
	cl.tail.next = newNode
	cl.tail = newNode

	cl.size++

}

func (cl *CircularTailList) DisplayListItems() {
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
	fmt.Printf("Head=%d and Tail=%d\n", cl.head.data, cl.tail.data)
	fmt.Printf("Head=%d and Tail=%d\n", cl.head.data, cl.tail.data)
}
