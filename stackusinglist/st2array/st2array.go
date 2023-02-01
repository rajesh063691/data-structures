package st2array

import (
	"fmt"
	"strings"
)

// Node ...
type Node struct {
	data int
	next *Node
}

// Stack ...
type Stack struct {
	size int
	top  *Node
}

// InitStack - initializes the stack
func (s *Stack) InitStack(size int) {
	s.size = size
	s.top = nil
}

// CreateNewNode - creates new node
func CreateNewNode(data int) (newNode *Node) {
	newNode = new(Node)
	newNode.data = data
	newNode.next = nil
	return
}

// Push - pushes data into stack
func (s *Stack) Push(ele int) {
	temp := s.top
	newNode := CreateNewNode(ele)
	// underflow condition
	if temp == nil {
		s.top = newNode
		return
	}
	// if stack is already populated - create new node and make newnode to top
	newNode.next = temp
	s.top = newNode
}

// Pop - pops data from stack
func (s *Stack) Pop() (popedValue int) {
	temp := s.top
	if temp == nil {
		fmt.Printf("stack is underflow, nothing to pop out.\n")
		return
	}
	// take poped item
	popedValue = temp.data
	// remove top and make top->next to new top
	s.top = temp.next
	temp.next = nil
	return
}

// Peek - returns top ref data
func (s *Stack) Peek() (peek int) {
	temp := s.top
	if temp == nil {
		fmt.Printf("stack is underflow, nothing to pop out.")
		return
	}

	peek = temp.data
	return
}

// Peek - display the stack data
func (s *Stack) Display() {
	temp := s.top
	if temp == nil {
		fmt.Printf("stack is underflow, nothing to display.\n")
		return
	}
	stackData := ""
	for temp != nil {
		stackData = stackData + fmt.Sprintf("%d ", temp.data)
		temp = temp.next
	}
	trimmedStack := strings.TrimSuffix(stackData, " ")
	fmt.Println(trimmedStack)
	fmt.Printf("Top : %d\n", s.top.data)
}
