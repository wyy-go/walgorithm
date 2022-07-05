package linked

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedAppend(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "a", "b", "c", "d"}

	list := NewLinkedList()
	for _, a := range arr {
		list.Append(a)
	}
	assert.Equal(t, len(arr), list.Len())
	assert.Equal(t, false, list.IsEmpty())
	assert.NotNil(t, list.Head())

	err := list.Insert(0, "x")
	assert.NoError(t, err)
}

func TestLinkedRemoveAt(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "a", "b", "c", "d"}

	list := NewLinkedList()
	for _, a := range arr {
		list.Append(a)
	}

	for _, a := range arr {
		item, err := list.RemoveAt(0)
		assert.NoError(t, err)
		assert.Equal(t, a, *item)
	}
	assert.Equal(t, 0, list.Len())
}

func TestLinkedIndexOf(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "a", "b", "c", "d"}

	list := NewLinkedList()
	for _, a := range arr {
		list.Append(a)
	}
	for i, a := range arr {
		assert.Equal(t, i, list.IndexOf(a))
	}
}

func TestLinked(t *testing.T) {
	list := NewLinkedList()
	err := list.Insert(10, 0)
	assert.Error(t, err)
	err = list.Insert(0, "x")
	assert.NoError(t, err)
	assert.Equal(t, 1, list.Len())
	_, err = list.RemoveAt(1)
	assert.Error(t, err)
	item, err := list.RemoveAt(0)
	assert.Equal(t, "x", *item)
}
