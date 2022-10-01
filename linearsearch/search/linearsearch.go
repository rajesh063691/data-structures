package search

// linear search using iteration

func LinearSearch(sl []int, item int) int {
	for i, val := range sl {
		if item == val {
			// return index if item is found
			return i
		}
	}
	// -1 means item could not be found
	return -1
}

// linear search using recursion

func LinearSearchRecursion(sl []int, item int, index int) int {
	if index >= len(sl) {
		return -1
	}
	if item == sl[index] {
		return index
	}
	return LinearSearchRecursion(sl, item, index+1)
}
