package course

import (
	"fmt"
	"leetcode-go/utils"
	"testing"
)

func Test(t *testing.T) {
	array := utils.GenerateRandomArray(5, 10, -10, 10)
	fmt.Println(array)

	selectionSort(array)

	fmt.Println(array)
}
