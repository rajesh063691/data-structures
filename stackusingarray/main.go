package main

import (
	"fmt"
	"practice_codes/data-structures/stackusingarray/st1array"
)

func main() {
	fmt.Printf("Stack imp using array\n")

	s := st1array.Stack{}
	s.InitStack(5)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Display()
	s.Peek()
	s.Pop()
	s.Peek()
	s.Display()
	s.Pop()
	s.Pop()
	s.Display()
}
