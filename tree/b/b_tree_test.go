package b

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBtree(t *testing.T) {
	bTree := NewBTree(2)

	for _, item := range perm(100) {
		if x := bTree.ReplaceOrInsert(item); x != nil {
			fmt.Println("insert didn't find item", item)
		} else {
			assert.Nil(t, x)
		}

	}

	fmt.Println(bTree.Min())
	fmt.Println(bTree.Max())
	fmt.Println(bTree.Len())

	fmt.Println(bTree.DeleteMin())
	fmt.Println(bTree.DeleteMax())
	fmt.Println(bTree.Len())

	fmt.Println(bTree.Get(Int(55)))
	fmt.Println(bTree.Has(Int(111)))
	fmt.Println(bTree.Delete(Int(98)))

	// 升序遍历
	bTree.Ascend(func(i Item) bool {
		fmt.Println(i)
		return true
	})

	bTree.AscendRange(Int(10), Int(20), func(i Item) bool {
		fmt.Println(i)
		return true
	})

	// 降序遍历
	bTree.Descend(func(i Item) bool {
		fmt.Println(i)
		return true
	})
}
