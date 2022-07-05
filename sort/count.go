package sort

// CountSort 作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
func CountSort(arr []int) []int {
	max := getMax(arr)
	sortedArr := make([]int, len(arr))
	countsArr := make([]int, max+1) // max+1 是为了防止 countsArr[] 计数时溢出

	// 元素计数
	for _, v := range arr {
		countsArr[v]++
	}

	// 统计独个数字个数并累加
	for i := 1; i <= max; i++ {
		countsArr[i] += countsArr[i-1]
	}

	// 让 arr 中每个元素找到其位置
	for _, v := range arr {
		sortedArr[countsArr[v]-1] = v
		//fmt.Print(countsArr[v]-1, " ")
		// 保证稳定性
		countsArr[v]--
	}
	return sortedArr
}
