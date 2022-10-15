package main

import (
	"fmt"
	"practice_codes/data-structures/mergesort/merge"
	usingindex "practice_codes/data-structures/mergesort/sortusingindex"
)

/*
1. find the mid element
2. call sort method for left sub array
3. call sort method for right sub array
4. finally, call the merge method for merging 2 sub arrays
*/

func main() {
	fmt.Println("Merge Sort...")
	sl := []int{12, 11, 13, 5, 6, 7}

	//arr := merge.Merge_Sort(sl)
	usingindex.MergeSortWithIndex(sl, 0, len(sl))
	fmt.Println(sl)

	fmt.Println("Is List Sorted: ", merge.VerifyIfListIsSorted(sl))

}
