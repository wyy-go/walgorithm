package search

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestIdxSearch(t *testing.T) {
	arr := []int{1, 8, 14, 6, 9, 10, 22, 34, 18, 19, 31, 40, 38, 54, 66, 46, 71, 78, 68, 80, 85, 100, 94, 88, 96, 87}
	// 先排序
	sort.Ints(arr)

	for i, v := range arr {
		assert.Equal(t, i, IdxSearch(arr, v))
	}
}
