package quicksort

import (
	"math/rand"
	"time"
)

func QuicksortShuffle(arr []int, start int, end int, isShuffle bool) {
	if end > start {
		if isShuffle {
			shuffle(arr, start, end)
		}

		pivot := partition(arr, start, end)
		QuicksortShuffle(arr, pivot+1, end, false)
		QuicksortShuffle(arr, start, pivot-1, false)
	}
}

func shuffle(arr []int, start int, end int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}
