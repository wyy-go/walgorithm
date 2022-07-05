package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getBinarySearchTree() *BinarySearchTree {
	tree := NewBinarySearchTree()
	tree.AddE(3)
	tree.AddE(0)
	tree.AddE(2)
	tree.AddE(5)
	tree.AddE(4)
	return tree
}

func TestBinarySearchTreeAdd(t *testing.T) {

	tree := getBinarySearchTree()

	assert.Equal(t, 5, tree.Size())
	assert.NotNil(t, tree.Root())
	assert.False(t, tree.IsEmpty())

	tree.PreOrder()
	tree.MidOrder()
	tree.BackOrder()
	tree.LevelOrder()
}

func TestBinarySearchTreeContains(t *testing.T) {
	tree := getBinarySearchTree()
	assert.True(t, tree.Contains(5))
}

func TestBinarySearchTreeRemove(t *testing.T) {
	tree := getBinarySearchTree()
	tree.Remove(3)
	tree.PreOrder()

	tree.RemoveMin()
	tree.PreOrder()

	tree.RemoveMax()
	tree.PreOrder()
}

func TestBinarySearchTreeFind(t *testing.T) {
	tree := getBinarySearchTree()
	assert.Equal(t, 0, tree.FindMin())
	assert.Equal(t, 5, tree.FindMax())
}
