package course

// 冒泡排序
//
// 重复 n-1轮以下这个步骤（最后一个元素自己不用和自己进行比较）
// 将第 i 个元素依次与后面的 n-i 元素进行比较
// 如果遇到比自己大（小）的元素，进行交换，然后进行下一轮循环，i+1 继续重复步骤
// 每一轮，i~n-i 中最大（小）的元素会冒泡到 n-i 的位置
//
// 时间复杂度：O(n^2)
func bubbleSort(arr []int64) {
	//临界条件判断
	if arr == nil || len(arr) < 2 {
		return
	}

	//最后一个元素自己不用和自己进行比较
	//只需要循环 n-1
	//外层 for 循环控制循环次数
	for i := 1; i < len(arr); i++ {
		//添加标志位，如果 flag 没有变为 false，说明没有进行交换，已经为有序，不需要继续排序
		flag := true
		//第一轮过后最大的元素会出现在末尾
		for j := 0; j < len(arr)-1; j++ {
			//当前位与下一位进行比较
			//较大的交换往末尾冒泡
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func swap(arr []int64, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
