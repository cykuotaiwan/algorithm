package quicksort

import (
	"math/rand"
	"time"
)

func QuicksortRandPivot(arr []int, start int, end int) {
	if end > start {
		pivot := partitionRand(arr, start, end)
		QuicksortRandPivot(arr, pivot+1, end)
		QuicksortRandPivot(arr, start, pivot-1)
	}
}

func partitionRand(arr []int, start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(end-start) + start
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
