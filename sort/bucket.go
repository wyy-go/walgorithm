package sort

// BucketSort
// 桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
// 桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，
// 每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。
func BucketSort(arr []int) []int {
	//桶数
	num := len(arr)
	//k（数组最大值）
	max := getMax(arr)
	//二维切片
	buckets := make([][]int, num)

	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max //分配桶index = value * (n-1) /k

		buckets[index] = append(buckets[index], arr[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])

			copy(arr[tmpPos:], buckets[i])

			tmpPos += bucketLen
		}
	}

	return arr
}

// 此处实现插入排序方式，其实可以用任意其他排序方式
func sortInBucket(bucket []int) {
	length := len(bucket)
	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i - 1
		//将选出的被排数比较后插入左边有序区
		//注意j >= 0必须在前边，否则会数组越界
		for j >= 0 && backup < bucket[j] {
			bucket[j+1] = bucket[j] //移动有序数组
			j--                     //反向移动下标
		}
		bucket[j+1] = backup //插队插入移动后的空位
	}
}
