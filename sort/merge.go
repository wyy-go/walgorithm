package sort

func merge(left, right []int) []int {
	// 遍历比较后合并两个数组
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		// 数组提前排序完毕
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		// 比较更小的值追加到 res[] 中
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}

func MergeSort(arr []int) []int {
	// 递归结束条件
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	arr = merge(left, right)
	return arr
}
