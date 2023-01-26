package crlistusingtailonly

import (
	"fmt"
	"strings"
)

// Node - actual node to store data and next and prev pointers
type Node struct {
	data int
	next *Node
}

// CircularTailList  - stores only tail pointer
type CircularTailList struct {
	size int
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
	curr := cl.tail
	newNode := CreateNewNode(item)
	// if head is nil, then list is empty. so first node will be the head node
	if curr == nil {
		cl.tail = newNode
		// newnodes only becomes tails next
		cl.tail.next = newNode
		cl.size++
		return
	}
	// insert newNode after tail node
	newNode.next = curr.next
	curr.next = newNode
	cl.tail = newNode

	cl.size++

}

func (cl *CircularTailList) DeleteAtGivenPosition(pos int) {
	if pos < 1 || pos > cl.size {
		fmt.Printf("position not found %d, can not delete the node.\n", pos)
		return
	}

	curr := cl.tail
	// if head is nil, then list is empty. so first node will be the head node
	if curr == nil {
		fmt.Printf("list is empty, can not delete the node.\n")
		return
	}
	// insert at begining
	if pos == 1 {
		temp := curr.next
		curr.next = curr.next.next
		temp.next = nil
		// make tail to poin t to nil if there is only one node present in the list
		if curr.next == nil {
			cl.tail = nil
		}
		cl.size--
		return
	}

	// delete from last
	if pos == cl.size {
		head := cl.tail.next
		temp := cl.tail.next
		for temp.next != cl.tail {
			temp = temp.next
		}
		// point tails previous node to head and make previous node as tail node
		temp.next = head
		cl.tail = temp
		cl.size--
	}

	// delete at given position
	i := 1
	for curr.next != cl.tail {
		if i == pos-1 {
			temp := curr.next
			curr.next = curr.next.next
			// free deleted node
			temp.next = nil
		}
		curr = curr.next
	}

	cl.size--

}

func (cl *CircularTailList) DisplayListItems() {
	curr := cl.tail
	// If head is nil, then no item is in the list
	if curr == nil {
		fmt.Print("list is empty, nothing to display.\n")
		return
	}
	listDatas := ""
	// iterate over all the nodes and display the data
	for curr.next != cl.tail {
		listDatas = listDatas + fmt.Sprintf("%d => ", curr.next.data)
		curr = curr.next
	}
	// print the last element
	listDatas = listDatas + fmt.Sprintf("%d => ", curr.next.data)

	trimmedList := strings.TrimSuffix(listDatas, " => ")
	fmt.Printf("list size = %d\n", cl.size)
	fmt.Println(trimmedList)
	fmt.Printf("Head=%d and Tail=%d\n", cl.tail.next.data, cl.tail.data)
}
