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
