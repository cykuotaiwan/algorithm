package countsort

import (
	"math"
)

func CountSort(arr []int) []int {
	max := math.MinInt
	min := math.MaxInt

	for _, val := range arr {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}

	cntArr := make([]int, max-min+1)
	for _, val := range arr {
		cntArr[val-min]++
	}
	for i := 1; i < len(cntArr); i++ {
		cntArr[i] += cntArr[i-1]
	}

	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[cntArr[arr[i]-min]-1] = arr[i]
		cntArr[arr[i]-min] -= 1
		// fmt.Println(result)
		// fmt.Println(cntArr)
	}

	return result
}
