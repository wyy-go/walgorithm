package sort

// SelectSort
// 选择排序是一种简单直观的排序算法，无论什么数据进去都是 O(n2) 的时间复杂度。
// 所以用到它的时候，数据规模越小越好。唯一的好处可能就是不占用额外的内存空间了吧。
func SelectSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 假设无序区间第一个值为最小值
		min := i
		for next := min + 1; next < n; next++ {
			// 找到更小值，记录其位置
			if arr[min] > arr[next] {
				min = next
			}
		}
		// 将无序区间的最小值追加到有序区间
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
