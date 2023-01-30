package quicksort_test

import (
	qsort "algorithms/quickSort"
	"math/rand"
	"testing"
	"time"
)

func generateRandArr(length int, shuffle bool) []int {
	arr := make([]int, 0, length)

	for i := 0; i < length; i++ {
		arr = append(arr, i+1)
	}
	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
	}
	return arr
}

func isSorted(arr []int) bool {
	isSorted := true
	for i := 0; isSorted && i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			isSorted = false
		}
	}

	return isSorted
}

func TestQuickSort(t *testing.T) {
	arr1ku := generateRandArr(1000, false)
	arr10ku := generateRandArr(10000, false)
	arr100ku := generateRandArr(100000, false)
	arr1ks := generateRandArr(1000, true)
	arr10ks := generateRandArr(10000, true)
	arr100ks := generateRandArr(100000, true)
	t.Run("size 1k unshuffled", func(t *testing.T) {
		qsort.Quicksort(arr1ku, 0, len(arr1ku)-1)
		if !isSorted(arr1ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k unshuffled", func(t *testing.T) {
		qsort.Quicksort(arr10ku, 0, len(arr10ku)-1)
		if !isSorted(arr10ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k unshuffled", func(t *testing.T) {
		qsort.Quicksort(arr100ku, 0, len(arr100ku)-1)
		if !isSorted(arr100ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 1k shuffled", func(t *testing.T) {
		qsort.Quicksort(arr1ks, 0, len(arr1ks)-1)
		if !isSorted(arr1ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k shuffled", func(t *testing.T) {
		qsort.Quicksort(arr10ks, 0, len(arr10ks)-1)
		if !isSorted(arr10ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k shuffled", func(t *testing.T) {
		qsort.Quicksort(arr100ks, 0, len(arr100ks)-1)
		if !isSorted(arr100ks) {
			t.Error("array is not sorted")
		}
	})
}

func TestQuickSortRandPivot(t *testing.T) {
	arr1ku := generateRandArr(1000, false)
	arr10ku := generateRandArr(10000, false)
	arr100ku := generateRandArr(100000, false)
	arr1ks := generateRandArr(1000, true)
	arr10ks := generateRandArr(10000, true)
	arr100ks := generateRandArr(100000, true)
	t.Run("size 1k unshuffled", func(t *testing.T) {
		qsort.QuicksortRandPivot(arr1ku, 0, len(arr1ku)-1)
		if !isSorted(arr1ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k unshuffled", func(t *testing.T) {
		qsort.QuicksortRandPivot(arr10ku, 0, len(arr10ku)-1)
		if !isSorted(arr10ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k unshuffled", func(t *testing.T) {
		qsort.QuicksortRandPivot(arr100ku, 0, len(arr100ku)-1)
		if !isSorted(arr100ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 1k shuffled", func(t *testing.T) {
		qsort.QuicksortRandPivot(arr1ks, 0, len(arr1ks)-1)
		if !isSorted(arr1ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k shuffled", func(t *testing.T) {
		qsort.QuicksortRandPivot(arr10ks, 0, len(arr10ks)-1)
		if !isSorted(arr10ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k shuffled", func(t *testing.T) {
		qsort.QuicksortRandPivot(arr100ks, 0, len(arr100ks)-1)
		if !isSorted(arr100ks) {
			t.Error("array is not sorted")
		}
	})
}

func TestQuickSortMedian(t *testing.T) {
	arr1ku := generateRandArr(1000, false)
	arr10ku := generateRandArr(10000, false)
	arr100ku := generateRandArr(100000, false)
	arr1ks := generateRandArr(1000, true)
	arr10ks := generateRandArr(10000, true)
	arr100ks := generateRandArr(100000, true)
	t.Run("size 1k unshuffled", func(t *testing.T) {
		qsort.QuicksortMedian(arr1ku, 0, len(arr1ku)-1)
		if !isSorted(arr1ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k unshuffled", func(t *testing.T) {
		qsort.QuicksortMedian(arr10ku, 0, len(arr10ku)-1)
		if !isSorted(arr10ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k unshuffled", func(t *testing.T) {
		qsort.QuicksortMedian(arr100ku, 0, len(arr100ku)-1)
		if !isSorted(arr100ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 1k shuffled", func(t *testing.T) {
		qsort.QuicksortMedian(arr1ks, 0, len(arr1ks)-1)
		if !isSorted(arr1ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k shuffled", func(t *testing.T) {
		qsort.QuicksortMedian(arr10ks, 0, len(arr10ks)-1)
		if !isSorted(arr10ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k shuffled", func(t *testing.T) {
		qsort.QuicksortMedian(arr100ks, 0, len(arr100ks)-1)
		if !isSorted(arr100ks) {
			t.Error("array is not sorted")
		}
	})
}

func TestQuickSortShuffle(t *testing.T) {
	arr1ku := generateRandArr(1000, false)
	arr10ku := generateRandArr(10000, false)
	arr100ku := generateRandArr(100000, false)
	arr1ks := generateRandArr(1000, true)
	arr10ks := generateRandArr(10000, true)
	arr100ks := generateRandArr(100000, true)
	t.Run("size 1k unshuffled", func(t *testing.T) {
		qsort.QuicksortShuffle(arr1ku, 0, len(arr1ku)-1, true)
		if !isSorted(arr1ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k unshuffled", func(t *testing.T) {
		qsort.QuicksortShuffle(arr10ku, 0, len(arr10ku)-1, true)
		if !isSorted(arr10ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k unshuffled", func(t *testing.T) {
		qsort.QuicksortShuffle(arr100ku, 0, len(arr100ku)-1, true)
		if !isSorted(arr100ku) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 1k shuffled", func(t *testing.T) {
		qsort.QuicksortShuffle(arr1ks, 0, len(arr1ks)-1, true)
		if !isSorted(arr1ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 10k shuffled", func(t *testing.T) {
		qsort.QuicksortShuffle(arr10ks, 0, len(arr10ks)-1, true)
		if !isSorted(arr10ks) {
			t.Error("array is not sorted")
		}
	})
	t.Run("size 100k shuffled", func(t *testing.T) {
		qsort.QuicksortShuffle(arr100ks, 0, len(arr100ks)-1, true)
		if !isSorted(arr100ks) {
			t.Error("array is not sorted")
		}
	})
}
