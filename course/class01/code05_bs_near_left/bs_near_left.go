package course

// 在一个有序数组中，找 >= 某个数最左侧的位置
// 如 1,2,4,5,6,7,7,8,8,9,10 中，找到 >= 7 最左侧的位置,找到的位置为5
func nearLeft(arr []int64, num int64) int {
	// 临界条件判断
	if arr == nil || len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 && arr[0] < num {
		return -1
	}

	index := -1
	left, right := 0, len(arr)-1
	var mid int
	// 需要判断 left == right 的情况，此时 mid 刚好为 left 或 right
	for left <= right {
		mid = left + (right-left)>>1
		if arr[mid] >= num {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}

func nearRight(arr []int64, num int64) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 && arr[0] <= num {
		return -1
	}

	index := -1
	left, right := 0, len(arr)-1
	var mid int
	for left <= right {
		mid = left + (right-left)>>1
		if arr[mid] > num {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}
