package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	key := 251
	expectIndex := 7

	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320, 500}
	index := LinearSearch(items, key)
	assert.Equal(t, expectIndex, index)
}
