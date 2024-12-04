package exercise

// https://leetcode.cn/problems/contains-duplicate-ii/description/
// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，
// 满足 nums[i] == nums[j] 且 abs(i - j) <= k 。
// 如果存在，返回 true ；否则，返回 false 。
func containsNearbyDuplicate(nums []int, k int) bool {
	// 设置一个滑动窗口，窗口大小为 k
	set := map[int]struct{}{}
	for i, num := range nums {
		// 判断当前元素下标是否大于滑动窗口
		if i > k {
			// 大于移除窗口最左边元素 i-k-1
			delete(set, nums[i-k-1])
		}
		// 判断是否有相同的元素在滑动窗口中
		if _, ok := set[num]; ok {
			return true
		}
		// 将当前元素加入滑动窗口中
		set[num] = struct{}{}
	}
	return false
}
