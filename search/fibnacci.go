package search

// SearchFibnacci 斐波那契查找
// 有序的，非动态的序列
func SearchFibnacci(slice []int, key int) int {
	n := len(slice)
	fib := createFibnacci(20)
	//1、斐波那契下标，需满足F(k)-1>=n
	k := 0
	for !(fib[k]-1 >= n) {
		k++
	}
	//2、构建新序列，多出位补slice[n-1]
	tempS := make([]int, fib[k]-1)
	copy(tempS, slice)
	for i := n; i < len(tempS); i++ {
		tempS[i] = slice[n-1]
	}
	//3、开始斐波那契查找
	left, right := 0, n-1
	for left <= right {
		mid := left + fib[k-1] - 1
		if tempS[mid] > key {
			right = mid - 1
			k -= 1 //查找值在前面的F(k-1)位中
		} else if tempS[mid] < key {
			left = mid + 1
			k -= 2 //查找值在后面的F(k-2)位中
		} else {
			if mid < n {
				return mid
			} else { //位于tempS的填补位
				return n - 1
			}
		}
	}
	return -1
}

//  构建斐波那契数列

// createFibnacci
func createFibnacci(n int) []int {
	res := []int{1, 1}
	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res
}

func fibonacciValue(n int) int {
	if n < 2 {
		return n
	}

	var fibarry = [3]int{0, 1, 0}
	for i := 2; i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}

// fibonacciRecursive 递归
func fibonacciValueRecursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return fibonacciValueRecursive(n-1) + fibonacciValueRecursive(n-2)
	} else {
		return -1
	}
}

// fibonacciIterate 迭代
func fibonacciValueIterate(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		a, b := 1, 1
		result := 0
		for i := 3; i <= n; i++ {
			result = a + b
			a, b = b, result
		}
		return result
	}
}

// 闭包
func fibonacciValueClosure(n int) int {
	if n < n {
		return -1
	} else {
		f := fibonacciValueFunc()
		result := 0
		for i := 0; i < n; i++ {
			result = f()
		}
		return result
	}
}

func fibonacciValueFunc() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
