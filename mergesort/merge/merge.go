package merge

func Merge_Sort(arr []int) []int {
	//base case
	if len(arr) <= 1 {
		return arr
	}

	left_half, right_half := SpiltArray(arr)
	left := Merge_Sort(left_half)
	right := Merge_Sort(right_half)
	return Merge(left, right)
}

// actual functions which merge the 2 arrays
func Merge(left, right []int) (sorted_list []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted_list = append(sorted_list, left[i])
			i++
		} else {
			sorted_list = append(sorted_list, right[j])
			j++
		}
	}
	// incase left is greater then last element will be missed out
	for i < len(left) {
		sorted_list = append(sorted_list, left[i])
		i++
	}
	// incase right is greater then last element will be missed out
	for j < len(right) {
		sorted_list = append(sorted_list, right[j])
		j++
	}
	return
}

// SpiltArray - return spilted slices
func SpiltArray(arr []int) (left, right []int) {
	mid := len(arr) / 2
	left = arr[:mid]
	right = arr[mid:]
	return left, right
}

// verify that merge sort retured list is sorted or not
func VerifyIfListIsSorted(arr []int) bool {
	if len(arr) == 0 || len(arr) == 1 {
		return true
	}
	return arr[0] < arr[1] && VerifyIfListIsSorted(arr[1:])
}
