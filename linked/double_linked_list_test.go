package linked

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleLinkedListMisc(t *testing.T) {
	list := NewDoubleLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	assert.Equal(t, 5, list.Len())
	list.Clear()
	assert.Equal(t, 0, list.Len())
	assert.Nil(t, list.Front())
	assert.Nil(t, list.Back())
}

func TestDoubleLinkedListPushFront(t *testing.T) {
	list := NewDoubleLinkedList()
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushFront(4)
	list.PushFront(5)
	assert.Equal(t, []interface{}{5, 4, 3, 2, 1}, list.Values())
}

func TestDoubleLinkedListPushBack(t *testing.T) {
	list := NewDoubleLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5}, list.Values())
}

func TestDoubleLinkedListInsert(t *testing.T) {
	list := NewDoubleLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5}, list.Values())
}
