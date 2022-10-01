package search

// Binary Search Implementation
func BinarySearch(sl []int, item int, left int, right int) int {
	//base case
	if right < left {
		// means item is not present
		return -1
	}
	// find the middle element
	middle_index := (left + right) / 2
	mid_value := sl[middle_index]
	if mid_value == item {
		//item found, returning index
		return middle_index
	}
	if mid_value > item {
		//item could be after mid element as its sorted so calling BinarySearch from index mid+1 till last
		return BinarySearch(sl, item, left, middle_index-1)
	}

	if mid_value < item {
		//item could be before mid element as its sorted so calling BinarySearch from index 0 to mid-1
		return BinarySearch(sl, item, middle_index+1, right)
	}
	return -1
}

func BinaryWithWhileLoop(sl []int, target int) int {
	first := 0
	last := len(sl) - 1
	for first <= last {
		mid := (first + last) / 2
		if sl[mid] == target {
			return mid
		} else if sl[mid] < target {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return -1
}
