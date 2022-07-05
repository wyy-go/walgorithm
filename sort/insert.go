package sort

func InsertSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	// 遍历所有元素
	for i := 1; i < n; i++ {
		// 向前找位置
		for j := i; j > 0; j-- {
			// 合适位置插入
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
