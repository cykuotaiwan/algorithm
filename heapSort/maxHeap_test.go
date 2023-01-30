package heapSort_test

import (
	"algorithms/heapSort"
	"fmt"
	"testing"
)

func isSorted(arr []int, isAsc bool) bool {
	isSorted := true
	for i := 0; isSorted && (i < len(arr)-1); i++ {
		if isAsc && (arr[i] > arr[i+1]) {
			isSorted = false
		}
		if !isAsc && (arr[i] < arr[i+1]) {
			isSorted = false
		}
	}

	return isSorted
}

func isMaxHeapified(arr []int) bool {
	isMaxHeapified := true

	for i := 0; isMaxHeapified && (i < len(arr)/2-1); i++ {
		if arr[i] <= arr[2*i+1] || arr[i] <= arr[2*i+2] {
			isMaxHeapified = false
		}
	}
	return isMaxHeapified
}

func TestMaxHeap(t *testing.T) {
	t.Run("arr w/ even number elements", func(t *testing.T) {
		arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

		var hp heapSort.MaxHeap[int]
		hp = hp.New(arr)
		res := hp.GetHeap()
		fmt.Println(res)

		isMaxHeapified := isMaxHeapified(res)
		if !isMaxHeapified {
			t.Errorf("heap is not sorted")
		}
	})
	t.Run("arr w/ odd number elements", func(t *testing.T) {
		arr := []int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}

		var hp heapSort.MaxHeap[int]
		hp = hp.New(arr)
		res := hp.GetHeap()
		fmt.Println(res)

		isMaxHeapified := isMaxHeapified(res)
		if !isMaxHeapified {
			t.Errorf("heap is not sorted")
		}
	})
}

func TestSort(t *testing.T) {
	t.Run("arr w/ even number elements", func(t *testing.T) {
		arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

		var hp heapSort.MaxHeap[int]
		hp = heapSort.HeapSort(hp.New(arr))

		res := hp.GetHeap()
		fmt.Println(res)

		isSorted := isSorted(res, true)
		if !isSorted {
			t.Errorf("heap is not sorted")
		}
	})
	t.Run("arr w/ odd number elements", func(t *testing.T) {
		arr := []int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}

		var hp heapSort.MaxHeap[int]
		hp = heapSort.HeapSort(hp.New(arr))

		res := hp.GetHeap()
		fmt.Println(res)

		isSorted := isSorted(res, true)
		if !isSorted {
			t.Errorf("heap is not sorted")
		}
	})
}

func TestPeak(t *testing.T) {
	t.Run("arr w/ even number elements", func(t *testing.T) {
		arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

		var hp heapSort.MaxHeap[int]
		hp = hp.New(arr)

		res, _ := hp.PeakMax()
		heap := hp.GetHeap()
		fmt.Println(heap)

		if res != 16 {
			t.Errorf("heap Peak() error: %d", res)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("arr w/ even number elements", func(t *testing.T) {
		arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

		var hp heapSort.MaxHeap[int]
		hp = hp.New(arr)

		res, hp, _ := hp.PopMax()
		heap := hp.GetHeap()
		fmt.Println(heap)

		if res != 16 {
			t.Errorf("heap Pop() error: %d", res)
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("arr w/ even number elements", func(t *testing.T) {
		arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

		var hp heapSort.MaxHeap[int]
		hp = hp.New(arr).Push(20)

		heap := hp.GetHeap()
		fmt.Println(heap)

		if heap[0] != 20 {
			t.Errorf("heap Push() error: %d", heap[0])
		}
	})
}
