package main

import (
	"fmt"
	"practice_codes/data-structures/stackusinglist/st2array"
)

func main() {
	fmt.Printf("<======Stack imp using linked list======>\n")
	s := st2array.Stack{}
	s.InitStack(5)
	s.Display()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	s.Display()

	s.Pop()
	s.Pop()
	s.Display()
}
