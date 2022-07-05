package sort

func ShellSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	step := n / 2

	// 步长减少到 0 则排序完毕
	for step > 0 {

		// 遍历第一个步长区间之后的所有元素
		for i := step; i < n; i++ {
			j := i
			// 前一个元素更大则交换值
			// j >= step	// 避免向下越界
			for j >= step && arr[j-step] > arr[j] {
				arr[j-step], arr[j] = arr[j], arr[j-step]
				j -= step
			}
		}
		step /= 2
	}
	return arr
}
