package util

func IsSorted(arr []int) bool {
	isSorted := true
	for i := 0; isSorted && i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			isSorted = false
		}
	}

	return isSorted
}
