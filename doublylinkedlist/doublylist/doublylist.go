package doublylist

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
	// if head is nil, then list is empty. so first node will be the head node
	if curr == nil {
		curr = newNode
		dl.head = curr
		dl.size++
		return
	}
	// iterate till the second last element
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
	newNode.prev = curr
	dl.size++

}

func (dl *DoublyList) AddNodeAtFirstPosition(item int) {
	curr := dl.head
	newNode := CreateNewNode(item)
	if curr == nil {
		curr = newNode
		dl.head = curr
		dl.size++
		return
	}

	newNode.next = curr
	curr.prev = newNode
	dl.head = newNode
	dl.size++
}

func (dl *DoublyList) AddNodeAtGivenPosition(item int, pos int) {
	// check for position boundaries conditions
	if pos < 1 || pos > dl.size {
		fmt.Printf("position ::%d:: not found.\n", pos)
		return
	}

	curr := dl.head
	newNode := CreateNewNode(item)
	if curr == nil {
		curr = newNode
		dl.head = curr
		dl.size++
		return
	}
	i := 1
	for curr != nil {
		// insert new node just before the given position
		if i == pos-1 {
			newNode.prev = curr
			newNode.next = curr.next
			curr.next = newNode
			newNode.next.prev = newNode
			dl.size++
			return
		}
		curr = curr.next
		i++
	}
}

func (dl *DoublyList) DeleteNodeFromGivenPosition(pos int) {
	curr := dl.head
	// If no item present in list
	if curr == nil {
		fmt.Printf("list is empty. position ::%d:: not found.\n", pos)
		return
	}
	// position given is outside the boundaries condition
	if pos < 1 || pos > dl.size {
		fmt.Printf("given node position ::%d:: is outside the list boundaries so can not delete the node.\n", pos)
		return
	}
	i := 1
	for curr != nil {
		// insert new node just before the given position
		if i == pos-1 {
			// this will check if deleted positionis the last position
			if curr.next.next == nil {
				curr.next.prev = nil
				curr.next = nil
			} else {
				curr.next.next.prev = curr
				curr.next = curr.next.next
			}
			dl.size--
			return
		}
		curr = curr.next
		i++
	}
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
}
