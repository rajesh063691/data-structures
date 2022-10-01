package main

import "fmt"

//1. Selection phase
func QuickSelect(sl []int, startIndex, endIndex int, k int) int {

	pivot := Partition(sl, startIndex, endIndex)
	if pivot > k {
		// means item should be on left side of array. So, discard the right array
		return QuickSelect(sl, startIndex, pivot-1, k)
	} else if pivot < k {
		// means ites should be on right side of the arr. So, discard the left array
		QuickSelect(sl, pivot+1, endIndex, k)
	}
	// means, item is at pivot element
	return sl[pivot]
}

// 2. Partion phase
func Partition(sl []int, startIndex, endIndex int) int {
	//considering last index as pivot
	pivot := (startIndex + endIndex) / 2
	temp := sl[pivot]
	sl[pivot] = sl[endIndex]
	sl[endIndex] = temp

	for i := startIndex; i < endIndex; i++ {
		if sl[i] < sl[endIndex] {
			//swap ith item with start index and increase the start index
			//Swap(i, startIndex)
			temp := sl[i]
			sl[i] = sl[startIndex]
			sl[startIndex] = temp
			startIndex = startIndex + 1
		}
	}
	// swapping startindex with lastindex
	//Swap(startIndex, endIndex)
	temp = sl[startIndex]
	sl[startIndex] = sl[endIndex]
	sl[endIndex] = temp

	return startIndex
}

// func Swap(i, j int) {
// 	temp := sl[i]
// 	sl[i] = sl[j]
// 	sl[j] = temp
// }

func main() {
	fmt.Println("Selection Search")
	k := 3
	var sl = []int{1, -2, 5}
	fmt.Printf("2nd smallest number= %d\n", QuickSelect(sl, 0, len(sl)-1, k-1))
}
