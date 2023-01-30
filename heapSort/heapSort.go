package heapSort

func HeapSort(mh MaxHeap[int]) MaxHeap[int] {
	if mh.sorted {
		return mh
	}

	for i := len(mh.heap) - 1; i >= 1; i-- {
		// exchange index 1 with index i
		tmp := mh.heap[i]
		mh.heap[i] = mh.heap[0]
		mh.heap[0] = tmp

		mh.heapSize--
		mh.MaxHeapify(0)
	}
	mh.heapSize = len(mh.heap)
	return mh
}
