package sort

import (
	"math/rand"
	"time"
)

func getArr(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max + 1)
	}
	return arr
}

// 获取数组中最大的值
func getMax(arr []int) (max int) {
	max = arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return
}
