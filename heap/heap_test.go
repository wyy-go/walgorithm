package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap(t *testing.T) {

	var data = &BinaryHeap{data: []int{6, 9, 2, 22, 44, 8, 12}}
	assert.Equal(t, []int{9, 2, 22, 44, 8, 12}, data.Values())

	data.Add(88)
	data.AddSlice([]int{3, 28})
	assert.Equal(t, []int{9, 2, 22, 44, 8, 12, 88, 3, 28}, data.Values())

	data.BuildMax(data.Length())
	assert.Equal(t, []int{88, 44, 9, 2, 8, 12, 22, 3, 28}, data.Values(), "max")

	data.BuildMin(data.Length())
	assert.Equal(t, []int{2, 88, 9, 44, 8, 12, 22, 3, 28}, data.Values(), "min")

	data.SortAsc()
	assert.Equal(t, []int{2, 3, 8, 9, 12, 22, 28, 44, 88}, data.Values(), "asc")

	data.SortDesc()
	assert.Equal(t, []int{88, 44, 28, 22, 12, 9, 8, 3, 2}, data.Values(), "desc")
}
