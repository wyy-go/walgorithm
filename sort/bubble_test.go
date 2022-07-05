package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 17, 10, 0, 13}
	want := []int{0, 9, 10, 13, 17}

	fmt.Println("[UNSORTED]:\t", arr)

	t.Run("BubbleSort", func(t *testing.T) {
		arr = BubbleSort(arr)
		fmt.Println("[SORTED]:\t", arr)
		assert.Equal(t, want, arr)
	})

	t.Run("BubbleSortByShort", func(t *testing.T) {
		arr = BubbleSortByShort(arr)
		fmt.Println("[SORTED]:\t", arr)
		assert.Equal(t, want, arr)
	})

	t.Run("CockTailSort", func(t *testing.T) {
		arr = CockTailSort(arr)
		fmt.Println("[SORTED]:\t", arr)
		assert.Equal(t, want, arr)
	})
}
