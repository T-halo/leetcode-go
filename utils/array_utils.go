package utils

import "math/rand"

func GenerateRandomArray(minSize, maxSize, minValue, maxValue int64) []int64 {
	n := rand.Int63n(maxSize-minSize+1) + minSize
	array := make([]int64, n)
	for i := range array {
		array[i] = rand.Int63n(maxValue-minValue+1) + minValue
	}
	return array
}

type Int64Slice []int64

func (arr Int64Slice) Len() int {
	return len(arr)
}

func (arr Int64Slice) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Int64Slice) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
