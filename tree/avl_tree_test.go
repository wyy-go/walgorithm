package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVL(t *testing.T) {

	bsTree := NewAVL(100, 1)
	newTree := bsTree.Insert(60)
	newTree = bsTree.Insert(120)
	newTree = bsTree.Insert(110)
	newTree = bsTree.Insert(130)
	newTree = bsTree.Insert(105)

	fmt.Println(newTree.getAll())

	newTree.Delete(110)

	fmt.Println(newTree.getAll())

	assert.True(t, newTree.Search(60))

}
