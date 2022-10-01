package setsops

import (
	"errors"
	"fmt"
)

// Set - set holds unique elements
type Set struct {
	Elements map[string]struct{}
}

// NewSet - initializes the set
func NewSet() *Set {
	return &Set{
		Elements: make(map[string]struct{}),
	}
}

// Add - add elements in set
func (s *Set) Delete(element string) error {
	_, exists := s.Elements[element]
	if !exists {
		return errors.New("element not found")
	}
	delete(s.Elements, element)
	return nil
}

// Contains - checks if key exists in set
func (s *Set) Contains(element string) bool {
	_, exists := s.Elements[element]
	return exists
}

// // Add - add elements in set
// func (s *Set) Add(element string) {
// 	s.Elements[element] = struct{}{}
// }
// Add - add elements in set
func (s *Set) Add(element string) {
	s.Elements[element] = struct{}{}
}

// List - lists elements in the set
func (s *Set) Lists() {
	for key, _ := range s.Elements {
		fmt.Println(key)
	}
}
