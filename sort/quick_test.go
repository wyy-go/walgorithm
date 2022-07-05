package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 17, 10, 0, 13}
	want := []int{0, 9, 10, 13, 17}

	fmt.Println("[NSORTED]:\t", arr)
	arr = QuickSort(arr)
	fmt.Println("[SORTED]:\t", arr)
	assert.Equal(t, want, arr)
}
