package tree

import "fmt"

type BinaryTreeNode struct {
	Data  interface{}     //数据
	Left  *BinaryTreeNode //左子树
	Right *BinaryTreeNode //右子树
}

// NewBinaryTree 二叉树
func NewBinaryTree(data interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data: data,
	}
}

func CreateNode(data interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data: data,
	}
}

func (tree *BinaryTreeNode) String() string {
	return fmt.Sprintf(" %v", tree.Data)
}

//前序遍历:以当前节点为根节点，根——>左——>右
func (tree *BinaryTreeNode) PreTraverseRecursion(print bool) (treeString string) {

	if tree == nil {
		return
	}

	treeString += tree.String()

	if tree.Left != nil {
		treeString += tree.Left.PreTraverseRecursion(false)
	}

	if tree.Right != nil {
		treeString += tree.Right.PreTraverseRecursion(false)
	}

	if print {
		fmt.Println(fmt.Sprintf("前序遍历:[%v]", treeString))
	}

	return
}

//中序遍历:以当前节点为根节点，左——>根——>右
func (tree *BinaryTreeNode) MidTraverseRecursion(print bool) (treeString string) {
	if tree == nil {
		return
	}

	if tree.Left != nil {
		treeString += tree.Left.MidTraverseRecursion(false)
	}

	treeString += tree.String()

	if tree.Right != nil {
		treeString += tree.Right.MidTraverseRecursion(false)
	}
	if print {
		fmt.Println(fmt.Sprintf("中序遍历:[%v]", treeString))
	}

	return
}

//后续遍历：以当前节点为根节点，左——>右——>根
func (tree *BinaryTreeNode) PostTraverseRecursion(print bool) (treeString string) {
	if tree == nil {
		return
	}
	if tree.Left != nil {
		treeString += tree.Left.PostTraverseRecursion(false)
	}
	if tree.Right != nil {
		treeString += tree.Right.PostTraverseRecursion(false)
	}
	treeString += tree.String()
	if print {
		fmt.Println(fmt.Sprintf("后序遍历:[%v]", treeString))
	}
	return
}

//层次遍历(广度优先遍历)
func (tree *BinaryTreeNode) BreadthFirstSearch(print bool) {
	if tree == nil {
		return
	}
	var result []interface{}
	nodes := []*BinaryTreeNode{tree}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result, curNode.Data)
		if curNode.Left != nil {
			nodes = append(nodes, curNode.Left)
		}
		if curNode.Right != nil {
			nodes = append(nodes, curNode.Right)
		}
	}

	if print {
		fmt.Print(fmt.Sprintf("层次遍历:["))
		for _, v := range result {
			fmt.Print(" ")
			fmt.Print(v)
		}
		fmt.Println("]")
	}

}

//层数(递归实现)
//对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
func (tree *BinaryTreeNode) Layers() int {
	if tree == nil {
		return 0
	}
	leftLayers := tree.Left.Layers()
	rightLayers := tree.Right.Layers()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

//层数(非递归实现)
//借助队列，在进行按层遍历时，记录遍历的层数即可
func (tree *BinaryTreeNode) LayersByQueue() int {
	if tree == nil {
		return 0
	}
	layers := 0
	nodes := []*BinaryTreeNode{tree}
	for len(nodes) > 0 {
		layers++
		size := len(nodes) //每层的节点数
		count := 0
		for count < size {
			count++
			curNode := nodes[0]
			nodes = nodes[1:]
			if curNode.Left != nil {
				nodes = append(nodes, curNode.Left)
			}
			if curNode.Right != nil {
				nodes = append(nodes, curNode.Right)
			}
		}
	}
	return layers
}
