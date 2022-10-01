package main

import (
	"fmt"
	"practice_codes/data-structures/binarysearch/search"
	"time"
)

func main() {
	fmt.Println("Binary Search!!!")
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 15, 17, 19, 20, 15, 30, 34, 40, 46, 50, 53, 55}
	st_time := time.Now()
	fmt.Printf("item 15 found at index: %d\n", search.BinarySearch(sl, 123, 0, len(sl)-1))
	fmt.Println("Total Run Time: ", time.Since(st_time).Microseconds())
	//fmt.Printf("item 0 found at index: %d\n", search.BinarySearch(sl, 0, 0, len(sl)-1))
	// fmt.Printf("item 17 found at index: %d\n", search.BinarySearch(sl, 17, 0, len(sl)-1))
	// fmt.Printf("item 55 found at index: %d\n", search.BinarySearch(sl, 55, 0, len(sl)-1))
	// fmt.Printf("item 56 found at index: %d\n", search.BinarySearch(sl, 56, 0, len(sl)-1))
	st_time = time.Now()
	fmt.Printf("item 15 found at index: %d\n", search.BinaryWithWhileLoop(sl, 123))
	fmt.Println("Total Run Time: ", time.Since(st_time).Microseconds())
}
