package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	//              3
	//           /    \
	//          0      5
	//           \    /
	//            2  4
	tree := NewBinaryTree(3)
	tree.Left = CreateNode(0)
	tree.Left.Right = CreateNode(2)
	tree.Right = CreateNode(5)
	tree.Right.Left = CreateNode(4)

	// 前序
	tree.PreTraverseRecursion(false)
	// 中序
	tree.MidTraverseRecursion(false)
	// 后续
	tree.PostTraverseRecursion(false)
	// 广度优先遍历
	tree.BreadthFirstSearch(false)

	// 层数
	assert.Equal(t, 3, tree.Layers())
	assert.Equal(t, 3, tree.LayersByQueue())
}
