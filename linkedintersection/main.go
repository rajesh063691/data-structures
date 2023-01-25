package main

import (
	"fmt"
	"practice_codes/data-structures/linkedintersection/intersection"
	"practice_codes/data-structures/linkedintersection/list"
)

func main() {
	lt := new(list.Node)
	lt.Data = 1
	lt1 := new(list.Node)
	lt1.Data = 2

	intNode := intersection.GetIntersectionNode(lt, lt1)
	fmt.Printf("%v", &intNode)

}
