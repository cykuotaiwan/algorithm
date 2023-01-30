package quicksort

func QuicksortMedian(arr []int, start int, end int) {
	if end > start {
		pivot := partitionMedian(arr, start, end)
		QuicksortMedian(arr, pivot+1, end)
		QuicksortMedian(arr, start, pivot-1)
	}
}

func medianOfThree(arr []int, start int, end int) int {
	mid := (start + end) / 2
	if arr[start] < arr[end] {
		arr[start], arr[end] = arr[end], arr[start]
	}
	if arr[start] < arr[mid] {
		arr[start], arr[mid] = arr[mid], arr[start]
	}
	if arr[mid] < arr[end] {
		arr[mid], arr[end] = arr[end], arr[mid]
	}

	return mid
}

func partitionMedian(arr []int, start int, end int) int {
	index := medianOfThree(arr, start, end)
	arr[index], arr[end] = arr[end], arr[index]
	pivot := arr[end]
	finish := start - 1

	for i := start; i <= end-1; i++ {
		if arr[i] <= pivot {
			finish++
			arr[finish], arr[i] = arr[i], arr[finish]
		}
	}
	arr[finish+1], arr[end] = arr[end], arr[finish+1]

	return finish + 1
}
