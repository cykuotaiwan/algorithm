package quicksort

func Quicksort(arr []int, start int, end int) {
	if end > start {
		pivot := partition(arr, start, end)
		Quicksort(arr, pivot+1, end)
		Quicksort(arr, start, pivot-1)
	}
}

func partition(arr []int, start int, end int) int {
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
