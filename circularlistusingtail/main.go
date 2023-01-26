package main

import (
	"fmt"
	"practice_codes/data-structures/circularlistusingtail/crlistusingtail"
)

func main() {
	fmt.Printf("Circular list using tail pointer\n")

	cl := new(crlistusingtail.CircularTailList)
	// cl.AddNodeAtLast(10)
	// cl.AddNodeAtLast(20)
	// cl.AddNodeAtLast(30)
	// cl.AddNodeAtLast(40)
	// cl.AddNodeAtLast(50)

	cl.DisplayListItems()
}
