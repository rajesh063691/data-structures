package usingindex

func MergeSortWithIndex(arr []int, start, end int) {
	//base case
	if start < end {

		//int mid = l + (r - l) / 2;
		mid := (start + end) / 2
		MergeSortWithIndex(arr, start, mid)
		MergeSortWithIndex(arr, mid+1, end)
		MergeWithIndex(arr, start, mid, end)
	}
}

func MergeWithIndex(arr []int, left, mid, end int) {
	L1 := arr[left:mid]
	L2 := arr[mid:end]

	i, j, k := 0, 0, 0
	for i < len(L1) && j < len(L2) {
		if L1[i] <= L2[j] {
			arr[k] = L1[i]
			i++
		} else {
			arr[k] = L2[j]
			j++
		}
		k++
	}
	// incase left is greater then last element will be missed out
	for i < len(L1) {
		arr[k] = L1[i]
		i++
		k++
	}
	// incase right is greater then last element will be missed out
	for j < len(L2) {
		arr[k] = L2[j]
		j++
		k++
	}
}
