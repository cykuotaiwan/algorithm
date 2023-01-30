package util

import (
	"math"
	"math/rand"
	"time"
)

func GenerateRandArr(length int, shuffle bool) []int {
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

func GenerateRandFixDigit(digitCnt int, length int) []int {
	arr := make([]int, 0, length)
	low := int(math.Pow10(digitCnt))
	high := int(math.Pow10(digitCnt+1)) - 1

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(high-low)+low)
	}

	return arr
}

func GenerateRandUniform(length int) []float64 {
	arr := make([]float64, 0, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Float64())
	}

	return arr
}
