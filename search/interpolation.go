package search

// InterpolationSearch
// 插值查找算法又称 插值搜索算法 ，是在 二分查找算法 的基础上改进得到的一种查找算法。
// 插值查找算法只适用于有序序列，换句话说，它只能在升序序列或者降序序列中查找目标元素
// 对于分布不均匀的有序序列来说，该算法不一定比二分查找要好。
func InterpolationSearch(arr []int, key int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := size * (key - arr[low]) / (arr[high] - arr[low])
			guess = low + offset
		}

		if arr[guess] < key {
			low = guess + 1
		} else if arr[guess] > key {
			high = guess - 1
		} else {
			return guess
		}

	}

	return -1
}
