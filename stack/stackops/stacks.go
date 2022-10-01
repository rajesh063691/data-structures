package stack

import "errors"

// Stack - struct to hold stack data
type Stack struct {
	Items []int
}

// Push - pushes elements into Stack
func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

// Pop - pops elements from stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() == -1 {
		err := errors.New("can not pop element :: stack is empty")
		return -1, err
	}
	l := len(s.Items) - 1
	popedItem := s.Items[l]
	s.Items = s.Items[:l]
	return popedItem, nil
}

// IsEmpty - Checks if stack is empty, return -1 if empty
func (s *Stack) IsEmpty() int {
	size := len(s.Items)
	if size == 0 {
		return -1
	}

	return size
}
