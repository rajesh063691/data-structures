package infixevaluation

import "fmt"

// Stack
type Stack struct {
	item []int
	size int
	top  int
}

// InitStack - initializes the stack with default value
func (s *Stack) InitStack(size int) {
	s.size = size
	s.top = -1
	s.item = make([]int, size)
}

// Push - push item into stack
func (s *Stack) Push(ele int) {
	// check for stack overflow condition
	if s.top == s.size-1 {
		fmt.Printf("stack is overflow\n")
		return
	}
	// increase the top index and insert ele
	s.top++
	s.item[s.top] = ele

}

// Pop - pop item from stack
func (s *Stack) Pop() {
	// check for stack underflow condition
	if s.top == -1 {
		fmt.Printf("stack is underflow\n")
		return
	}
	poppedItem := s.item[s.top]
	fmt.Printf("Popped Item : %d\n", poppedItem)
	s.top--

}

func (s *Stack) EvaluatePostFixExpression(exp string) {

}
