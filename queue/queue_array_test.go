package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	myq := NewQueue()
	myq.Enqueue(1)
	myq.Enqueue(2)
	myq.Enqueue(3)
	myq.Enqueue(4)

	assert.Equal(t, 4, myq.Size())
	assert.Equal(t, 1, myq.Dequeue())
	assert.Equal(t, 2, myq.Dequeue())
	assert.Equal(t, 3, myq.Dequeue())
	assert.Equal(t, 4, myq.Dequeue())

	assert.Equal(t, 0, myq.Size())

	myq.Enqueue(1)
	myq.Enqueue(2)
	myq.Enqueue(3)
	myq.Enqueue(4)
	assert.Equal(t, 4, myq.Size())
	myq.Clear()
	assert.Equal(t, 0, myq.Size())
}
