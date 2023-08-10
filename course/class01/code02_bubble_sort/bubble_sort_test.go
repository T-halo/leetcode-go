package course

import (
	"fmt"
	"leetcode-go/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := utils.GenerateRandomArray(5, 10, -10, 10)
	fmt.Println(array)
	bubbleSort(array)
	fmt.Println(array)
}
