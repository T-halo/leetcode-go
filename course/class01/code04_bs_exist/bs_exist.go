package code04_bs_exist

// 在一个有序数组里，判断一个数字是否存在
//
// 因为是有序数组，所以使用二分法可以较快找出数字是否存在
//
// 时间复杂度：最坏情况下 O(logn)
func bsExist(arr []int64, num int64) bool {
	//判断临界条件
	if arr == nil || len(arr) == 0 {
		return false
	}

	//二分法套路
	//先定义左边界、右边界、中间点
	l := 0
	r := len(arr) - 1
	mid := 0
	//如果 l < r，则说明两边界已经相遇，退出循环
	for l < r {
		//中间点 (l+r)/2
		//如果两边界较大时，上述情况可能会出现溢出的情况
		//所以用 l + ((r - l) >> 1)
		mid = l + ((r - l) >> 1)
		if arr[mid] > num {
			//右边界变为中间点-1
			r = mid - 1
		} else if arr[mid] < num {
			//左边界变为中间点+1
			l = mid + 1
		} else if arr[mid] == num {
			return true
		}
	}
	//判断相遇的位置是不是符合的值
	return arr[l] == num
}
