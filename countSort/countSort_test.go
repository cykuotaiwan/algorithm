package countsort_test

import (
	countsort "algorithms/countSort"
	"algorithms/util"
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	arr := util.GenerateRandArr(10, true)
	// arr := []int{2, 5, 3, 0, 2, 3, 0, 3}
	fmt.Println(arr)
	arr = countsort.CountSort(arr)
	fmt.Println(arr)
	if !util.IsSorted(arr) {
		t.Errorf("array is not sorted")
	}
}
