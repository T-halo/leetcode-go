package course

import (
	"fmt"
	"leetcode-go/utils"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	array := utils.GenerateRandomArray(5, 10, -10, 10)
	fmt.Println(array)
	insertionSort(array)
	fmt.Println(array)
}
