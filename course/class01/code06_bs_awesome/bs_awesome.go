package code06_bs_awesome

/**
 * 局部最小值，找到一个区间内的最小值，相邻的两个数不相等，其两边的数都大于它
 * 如[3,6,4,3,7,9,2,5,1]中，位置处于0,3,6,8的数都是区间最小值
 * <p>
 * 当区间内数字趋势有变化时，该区间内必有一个转折点，即极值，就是说左边呈下降趋势，而右边呈上升趋势时，该区间内必有一个最小值
 * 通过比较两边的大小趋势，可以找出有最小值的区间，由于这种性质，可以使用二分法
 *
 * 最大时间复杂度：O(log n)
 */
func getLessIndex(arr []int64) int {
	l := len(arr)
	if arr == nil || l == 0 {
		return -1
	}
	if l == 1 || arr[0] < arr[1] {
		return 0
	}
	if arr[l-2] > arr[l-1] {
		return l - 1
	}
	// 如果上述条件都不满足，说明中间点在中间，左边趋势下降，右边趋势上升
	// \            /
	//  \/\  /\    /
	//     \/  \  /
	//          \/
	left := 1
	right := l - 2
	mid := 0
	for left < right {
		mid = left + (right-left)>>1
		// 去中间点判断趋势，比左边高，说明最小值在左边，比右边高说明最小值在右边
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
