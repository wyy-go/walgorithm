package search

// 是顺序查找的一种结合改进；
// 将n个数据元素"按块有序"划分为m块（m ≤ n）。每一块中的结点不必有序，但块与块之间必须"按块有序"；
// 即第1块中任一元素的关键字都必须小于第2块中任一元素的关键字；
// 而第2块中任一元素又都必须小于第3块中的任一元素，……
//　　算法流程：
//　　step1 ：先选取各块中的最大关键字构成一个索引表；
//　　step2 ：查找分两个部分：先对索引表进行二分查找或顺序查找，以确定待查记录在哪一块中；然后，在已确定的块中用顺序法进行查找。

// 每个块中的元素个数
const defaultElementCount = 10

type indextype struct {
	key  int
	link int
}

// IdxSearch 分块查询 b块数 key查询关键字
func IdxSearch(arr []int, key int) int {

	// 先选取各块中的最大关键字构成一个索引表
	idx, block := createIndex(arr)
	low, high, mid, i := 0, block-1, 0, 0
	for low <= high {
		mid = (low + high) / 2
		if idx[mid].key >= key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	i = idx[high+1].link
	//fmt.Println(i, key)
	for i <= idx[high+1].link+defaultElementCount-1 && arr[i] != key {
		i++
	}

	if i <= idx[high+1].link+defaultElementCount-1 {
		return i
	} else {
		return -1
	}
}

func createIndex(arr []int) ([]indextype, int) {
	var bb int
	n := len(arr)
	b := n / defaultElementCount
	m := n % defaultElementCount

	bb = b
	if m > 0 {
		bb = bb + 1
	}

	idx := make([]indextype, bb)
	for i := 0; i < bb; i++ {
		idx[i].link = i * defaultElementCount
		idx[i].key = 0
	}

	for i := 0; i < bb; i++ {
		for j := 0; j < defaultElementCount; j++ {

			index := i*defaultElementCount + j
			if index > n-1 {
				break
			}

			temp := arr[index]
			if temp > idx[i].key {
				idx[i].key = temp
			}
		}
	}

	return idx, bb
}
