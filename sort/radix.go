package sort

// RadixSort
// 基数排序是一种非比较型整数排序算法，其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较。
// 由于整数也可以表达字符串（比如名字或日期）和特定格式的浮点数，所以基数排序也不是只能使用于整数
func RadixSort(arr []int) []int {
	max := getMax(arr)
	// 数组中最大值决定了循环次数，101 循环三次(个十百)
	for bit := 1; max/bit > 0; bit *= 10 {
		arr = bitSort(arr, bit)
	}
	return arr
}

//
// 对指定的位进行排序
// bit 可取 1，10，100 等值
//
func bitSort(arr []int, bit int) []int {
	n := len(arr)
	// 各个位的相同的数统计到 bitCounts[] 中
	bitCounts := make([]int, 10)
	for i := 0; i < n; i++ {
		num := (arr[i] / bit) % 10
		bitCounts[num]++
	}
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1]
	}

	tmp := make([]int, 10)
	for i := n - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitCounts[num]-1] = arr[i]
		bitCounts[num]--
	}
	for i := 0; i < n; i++ {
		arr[i] = tmp[i]
	}
	return arr
}
