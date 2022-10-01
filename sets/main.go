package main

import (
	"fmt"
	setops "practice_codes/data-structures/sets/setsops"
)

func main() {
	set := setops.NewSet()
	set.Add("Rajesh")
	set.Add("Kumar")
	set.Add("Yadav")
	fmt.Println("Set elements")
	set.Lists()
	fmt.Println("If Rajesh exists: ", set.Contains("Rajesh"))
	_ = set.Delete("Yadav")
	set.Lists()
}
