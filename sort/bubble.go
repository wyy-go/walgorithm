package sort

// BubbleSort 检测是否已提前有序
func BubbleSort(arr []int) []int {
	n := len(arr)
	// 遍历所有元素
	var isSorted bool
	for i := 0; i < n-1; i++ {
		isSorted = true
		for j := 0; j < n-i-1; j++ {
			// 左元素 > 右元素
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 发生交换则还未排序完毕
				isSorted = false
			}
		}
		// 没有发生交换则说明排序完成
		if isSorted {
			break
		}
	}
	return arr
}

// BubbleSortByShort 缩短扫描距离
func BubbleSortByShort(arr []int) []int {
	n := len(arr)
	end := n - 1
	for end > 0 {
		cur := 0
		for j := 0; j < end; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			// 记录每趟最后发生交换的位置，此位置之后均已有序，下一趟只需遍历到此位置即可
			cur = j
		}
		end = cur
	}
	return arr
}

// CockTailSort 鸡尾酒排序(双向冒泡排序)
func CockTailSort(arr []int) []int {
	n := len(arr)
	left := 0
	right := n - 1

	// left  以左已有序
	// right 以右已有序
	// 两个区间游标相遇则集合有序
	for left < right {
		// 从左到右，选出最大值
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--
		// 从右到左，选出最小值
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		left++
	}
	return arr
}
