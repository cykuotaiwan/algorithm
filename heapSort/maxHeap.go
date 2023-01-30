package heapSort

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MaxHeap[T constraints.Ordered] struct {
	sorted   bool
	heapSize int
	heap     []T
}

func (mh MaxHeap[T]) New(heapIn []T) MaxHeap[T] {
	mh.heap = heapIn
	mh.sorted = false
	mh.heapSize = len(heapIn)

	for i := mh.heapSize/2 - 1; i >= 1; i-- {
		mh.MaxHeapify(i)
	}

	return mh
}

func (mh MaxHeap[T]) findParent(index int) int {
	return index / 2
}

func (mh MaxHeap[T]) findLeft(index int) int {
	return 2*index + 1
}

func (mh MaxHeap[T]) findRight(index int) int {
	return 2*index + 2
}

func (mh MaxHeap[T]) MaxHeapify(index int) {
	if index >= mh.heapSize {
		return
	}
	l := mh.findLeft(index)
	r := mh.findRight(index)
	largest := index

	if l < mh.heapSize && mh.heap[l] > mh.heap[index] {
		largest = l
	}
	if r < mh.heapSize && mh.heap[r] > mh.heap[largest] {
		largest = r
	}
	if largest != index {
		tmp := mh.heap[index]
		mh.heap[index] = mh.heap[largest]
		mh.heap[largest] = tmp

		mh.MaxHeapify(largest)
	}

}

func (mh MaxHeap[T]) GetHeap() []T {
	return mh.heap
}

func (mh MaxHeap[T]) PeakMax() (T, error) {
	if mh.heapSize <= 0 {
		var empty T
		return empty, fmt.Errorf("heap is empty")
	}
	return mh.heap[0], nil
}

func (mh MaxHeap[T]) PopMax() (T, MaxHeap[T], error) {
	if mh.heapSize <= 0 {
		var empty T
		return empty, mh, fmt.Errorf("heap is empty")
	}
	ret := mh.heap[0]
	mh.heap = mh.heap[:mh.heapSize-1]
	mh.heapSize--

	return ret, mh, nil
}

func (mh MaxHeap[T]) Push(item T) MaxHeap[T] {
	mh.heap = append(mh.heap, item)
	mh.heapSize = len(mh.heap)

	index := mh.heapSize - 1
	for index > 0 && mh.heap[mh.findParent(index)] < mh.heap[index] {
		tmp := mh.heap[index]
		mh.heap[index] = mh.heap[mh.findParent(index)]
		mh.heap[mh.findParent(index)] = tmp

		index = mh.findParent(index)
	}

	return mh
}
