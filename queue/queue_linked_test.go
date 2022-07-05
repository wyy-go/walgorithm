package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkQueue(t *testing.T) {
	myq := NewLinkQueue()
	for i := 0; i < 10; i++ {
		myq.Enqueue(i)
	}

	assert.Equal(t, 10, myq.length(), "length")
	for i := 0; i < 10; i++ {
		v, err := myq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}
}
