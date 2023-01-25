package taildoublylist

import (
	"fmt"
	"strings"
)

// Node - actual node to store data and next and prev pointers
type Node struct {
	data int
	next *Node
	prev *Node
}

// DoublyList ...
type DoublyList struct {
	size int
	head *Node
	tail *Node
}

// CreateNewNode - creates new node with given data
func CreateNewNode(item int) (newNode *Node) {
	newNode = new(Node)
	newNode.data = item
	newNode.next = nil
	newNode.prev = nil
	return
}
func (dl *DoublyList) AddNodeAtLast(item int) {
	curr := dl.head
	newNode := CreateNewNode(item)
	// if head is nil, then list is empty. so first node will become head as well as tail
	if curr == nil {
		curr = newNode
		dl.head = curr
		dl.tail = curr
		dl.size++
		return
	}
	// iterate till the second last element
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
	newNode.prev = curr
	dl.tail = newNode
	dl.size++

}

func (dl *DoublyList) ReverseDoublyList() {
	if dl.head == nil {
		fmt.Printf("list is empty, can not reverse.\n")
		return
	}

	curr := dl.head
	// swap each nodes next and prev reference
	for curr != nil {
		next := curr.next
		curr.next = curr.prev
		curr.prev = next
		curr = next
	}
	// at last swap head and tail reference
	curr = dl.head
	dl.head = dl.tail
	dl.tail = curr

}

func (dl *DoublyList) DisplayListItems() {
	curr := dl.head
	// If head is nil, then no item is in the list
	if curr == nil {
		fmt.Print("List is empty...\n")
		return
	}
	listDatas := ""
	// iterate over all the nodes and display the data
	for curr != nil {
		listDatas = listDatas + fmt.Sprintf("%d <=> ", curr.data)
		curr = curr.next
	}

	trimmedList := strings.TrimSuffix(listDatas, " <=> ")
	fmt.Printf("list size = %d\n", dl.size)
	fmt.Println(trimmedList)

	fmt.Printf("Head =>%d && Tail =>%d\n", dl.head.data, dl.tail.data)
	//fmt.Printf("Head Prev =>%v && Tail Next =>%v\n", dl.head.prev, dl.tail.next)
}
