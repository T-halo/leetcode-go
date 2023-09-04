package course

import (
	"fmt"
	"leetcode-go/utils"
	"sort"
	"testing"
)

func TestBSNearLeft(t *testing.T) {
	array := utils.GenerateRandomArray(5, 10, 0, 10)
	sort.Sort(utils.Int64Slice(array))
	fmt.Println(array)
	index := nearLeft(array, 5)
	//index := nearRight(array, 5)
	fmt.Println(index)
}
