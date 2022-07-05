package b

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBPlusTree(t *testing.T) {

	tree := NewBPlusTree()

	for i := 0; i < 100; i++ {
		err := tree.Insert(i, fmt.Sprintf("test%d", i))
		assert.NoError(t, err)
	}

	for i := 0; i < 100; i++ {
		r, err := tree.Find(i, false)
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("test%d", i), r.Value)
	}

	// tree
	fmt.Println("================tree================")
	tree.PrintTree()
	// leaves
	fmt.Println("================leaves================")
	tree.PrintLeaves()
}
