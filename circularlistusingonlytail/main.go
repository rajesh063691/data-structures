package main

import (
	"fmt"
	"practice_codes/data-structures/circularlistusingonlytail/crlistusingtailonly"
)

func main() {
	fmt.Printf("Circular list using only tail pointer\n")

	cl := new(crlistusingtailonly.CircularTailList)
	cl.AddNodeAtLast(10)
	cl.AddNodeAtLast(20)
	cl.AddNodeAtLast(30)
	cl.AddNodeAtLast(40)
	cl.AddNodeAtLast(50)

	cl.DisplayListItems()

	// cl.DeleteAtGivenPosition(0)
	// cl.DeleteAtGivenPosition(1)
	// cl.DeleteAtGivenPosition(1)
	// cl.DeleteAtGivenPosition(3)
	// cl.DeleteAtGivenPosition(4)

	// cl.DisplayListItems()

	cl.ReverseCircularList()

	cl.DisplayListItems()
}
