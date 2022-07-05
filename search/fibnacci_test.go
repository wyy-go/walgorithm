package search

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestFibnacci(t *testing.T) {
	expect := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	fibarr := createFibnacci(20)
	assert.Equal(t, expect, fibarr)

	assert.Equal(t, 6765, fibonacciValue(20))
	assert.Equal(t, 6765, fibonacciValueRecursive(20))
	assert.Equal(t, 6765, fibonacciValueIterate(20))
	assert.Equal(t, 6765, fibonacciValueClosure(20))

	items := []int{1, 5, 3, 7, 9, 2, 6, 8}
	// sort
	sort.Ints(items) // {1, 2, 3, 5, 6, 7, 8, 9}

	assert.Equal(t, -1, SearchFibnacci(items, 100))
	assert.Equal(t, 0, SearchFibnacci(items, 1))
}
