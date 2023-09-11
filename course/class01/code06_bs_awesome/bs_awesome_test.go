package code06_bs_awesome

import (
	"fmt"
	"leetcode-go/utils"
	"testing"
)

func TestBSAwesome(t *testing.T) {
	array := utils.GenerateRandomArray(5, 10, 0, 10)
	fmt.Println(array)
	index := getLessIndex(array)
	fmt.Println(index)
}
