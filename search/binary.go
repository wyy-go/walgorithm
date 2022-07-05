package search

// BinarySearch 非递归
func BinarySearch(arr []int, key int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] < key {
			low = mid + 1
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// BinarySearchRecursive 递归
func BinarySearchRecursive(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	return bs(arr, v, 0, n-1)
}

func bs(arr []int, v int, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) >> 1
	if arr[mid] == v {
		return mid
	} else if arr[mid] > v {
		return bs(arr, v, low, mid-1)
	} else {
		return bs(arr, v, mid+1, high)
	}
}

// BinarySearchFirst 查找第一个等于给定值的元素
func BinarySearchFirst(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] > v {
			high = mid - 1
		} else if arr[mid] < v {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != v {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// BinarySearchLast 查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] > v {
			high = mid - 1
		} else if arr[mid] < v {
			low = mid + 1
		} else {
			if mid == n-1 || arr[mid+1] != v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

// BinarySearchFirstGT 查找第一个大于等于给定值的元素
func BinarySearchFirstGT(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		//避免溢出
		mid := low + (high-low)>>1
		if arr[mid] >= v {
			if mid == 0 || arr[mid-1] < v {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchLastLT 查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > v {
			high = mid - 1
		} else {
			if mid == n-1 || arr[mid+1] > v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
