package code01_selection_sort

// 选择排序
//
// 将所选择的元素下标作为临时变量的值，依次与后面的元素进行比较，如果遇到比其更小（大）的值
// 将临时变量的下标替换为该值的下标，再继续向后进行比较，直到最后一个元素
// 此时将临时变量下标对应元素与该轮选择的元素进行替换
//
// 时间复杂度为 O(n^2)
func selectionSort(arr []int64) {
	//临界值判断
	if arr == nil || len(arr) < 2 {
		return
	}

	// 0~N-1 找到最小值，与 0 位置进行交换
	// 1~N-1 找到最小值，与 1 位置进行交换
	// 2~N-1 找到最小值，与 2 位置进行交换
	for i := 0; i < len(arr); i++ {
		minIndex := i
		// i~N-1 找到最小值下标
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		swap(arr, minIndex, i)
	}
}

func swap(arr []int64, minIndex, i int) {
	if minIndex != i {
		arr[i] = arr[i] ^ arr[minIndex]
		arr[minIndex] = arr[i] ^ arr[minIndex]
		arr[i] = arr[i] ^ arr[minIndex]
	}
}
