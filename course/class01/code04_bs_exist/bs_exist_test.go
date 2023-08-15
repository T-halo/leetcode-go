package code04_bs_exist

import (
	"fmt"
	"testing"
)

func TestBSExist(t *testing.T) {
	arr := []int64{1, 3, 5, 7, 8, 9, 35, 67, 89}
	exist := bsExist(arr, 8)
	fmt.Println(exist)
}
