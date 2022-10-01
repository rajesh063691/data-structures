package main

import (
	"fmt"
	"practice_codes/data-structures/linearsearch/search"
)

func main() {

	fmt.Println("Linear Search")
	intSlice := []int{1, 2, 3, 8, 0, 5, 4, 2, 2, 4, 5, 10, 199, 12, 191, 456}
	fmt.Printf("item found at index: %d\n", search.LinearSearch(intSlice, 10))
	fmt.Printf("item found at index: %d\n", search.LinearSearch(intSlice, 5))
	fmt.Printf("item found at index: %d\n", search.LinearSearch(intSlice, 500))

	fmt.Printf("item found at index: %d\n", search.LinearSearchRecursion(intSlice, 10, 0))
	fmt.Printf("item found at index: %d\n", search.LinearSearchRecursion(intSlice, 5, 0))
	fmt.Printf("item found at index: %d\n", search.LinearSearchRecursion(intSlice, 500, 0))

}
