package main

import (
	"fmt"
	"log"
	stackops "practice_codes/data-structures/stack/stackops"
)

func main() {
	myStack := new(stackops.Stack)
	fmt.Println(myStack)

	peek, err := myStack.Peek()
	fmt.Printf("Peek : %d && err=%v\n", peek, err)

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Push(400)
	fmt.Println(myStack)

	popedItem, err := myStack.Pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Popped Item: ", popedItem)
	fmt.Println(myStack)

	popedItem, err = myStack.Pop()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Popped Item: ", popedItem)

	fmt.Println(myStack)
	peek, err = myStack.Peek()
	fmt.Printf("Peek : %d && err=%v\n", peek, err)

}
