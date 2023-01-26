package main

import (
	"fmt"
	crl "practice_codes/data-structures/circularlinkedlist/circularlist"
)

func main() {
	fmt.Printf("Circular list...\n")

	cls := new(crl.CircularList)
	cls.AddNodeAtLast(10)
	cls.AddNodeAtLast(20)
	cls.AddNodeAtLast(30)
	cls.AddNodeAtLast(40)
	cls.AddNodeAtLast(50)

	cls.DisplayListItems()

}
