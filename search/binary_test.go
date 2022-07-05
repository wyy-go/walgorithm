package search

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	items := []int{1, 5, 3, 7, 9, 2, 6, 8}

	// sort
	sort.Ints(items) // {1, 2, 3, 5, 6, 7, 8, 9}

	// 非递归
	assert.Equal(t, 6, BinarySearch(items, 8))
	assert.Equal(t, -1, BinarySearch(items, 100))
	// 递归
	assert.Equal(t, 6, BinarySearchRecursive(items, 8))
	assert.Equal(t, -1, BinarySearchRecursive(items, 100))
	// 查找第一个等于给定值的元素
	assert.Equal(t, 6, BinarySearchFirst(items, 8))
	assert.Equal(t, -1, BinarySearchFirst(items, 100))
	// 查找最后一个值等于给定值的元素
	assert.Equal(t, 6, BinarySearchLast(items, 8))
	assert.Equal(t, -1, BinarySearchLast(items, 100))
	// 查找第一个大于等于给定值的元素
	assert.Equal(t, 6, BinarySearchFirstGT(items, 8))
	assert.Equal(t, -1, BinarySearchFirstGT(items, 100))
	// 查找最后一个小于等于给定值的元素
	assert.Equal(t, 6, BinarySearchLastLT(items, 8))
	assert.Equal(t, -1, BinarySearchLastLT(items, 0))
}
