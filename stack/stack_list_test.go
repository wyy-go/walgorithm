package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackByList(t *testing.T) {

	s := NewStackByArray()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	assert.Equal(t, 4, s.Pop())
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())

}
