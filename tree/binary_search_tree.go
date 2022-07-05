package tree

import (
	"fmt"
	"strings"
)

type Node struct {
	E     int
	left  *Node
	right *Node
}

// 约定在此样例代码中我们的二分搜索树中没有重复元素
// 如果想包涵重复元素的话，只需要以下定义：
// 左子树小于等于此节点，或右子树大于等于节点
type BinarySearchTree struct {
	root *Node
	size int
}

// NewBinarySearchTree 二叉搜索树
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (t BinarySearchTree) Size() int {
	return t.size
}

func (t BinarySearchTree) Root() *Node {
	return t.root
}

// 判断二叉树是否为空
func (t *BinarySearchTree) IsEmpty() bool {
	if t.size == 0 {
		return true
	}
	return false
}

func NewNode(E int) *Node {
	return &Node{
		E:     E,
		left:  nil,
		right: nil,
	}
}

func (t *BinarySearchTree) AddE(e int) {
	t.root = t.add(t.root, e)
}

// 向以node为根的二分搜索树中插入元素E，递归算法
func (t *BinarySearchTree) add(node *Node, e int) *Node {
	if node == nil {
		t.size++
		return NewNode(e)
	}

	if e > node.E {
		node.right = t.add(node.right, e)
	} else if e < node.E {
		node.left = t.add(node.left, e)
	}
	return node
}

// 查找二分搜索中是否含有元素E
func (t *BinarySearchTree) Contains(e int) bool {
	return t.contains(t.root, e)
}

// 递归的方式查找元素是否存在
func (t *BinarySearchTree) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}
	if e == node.E {
		return true
	} else if e > node.E {
		return t.contains(node.right, e)
	} else {
		return t.contains(node.left, e)
	}
}

// 遍历算法
// 1.前序遍历
func (t *BinarySearchTree) PreOrder() {
	PreOrder(t.root)
	fmt.Println()
}

func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.E)
	PreOrder(node.left)
	PreOrder(node.right)
}

// 非递归的前序遍历
func (t *BinarySearchTree) PreOrderNR() {
	stack := make([]*Node, 0)
	stack = append(stack, t.root)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curNode.E)
		if curNode.right != nil {
			stack = append(stack, curNode.right)
		}
		if curNode.left != nil {
			stack = append(stack, curNode.left)
		}
	}
	fmt.Println()
}

// 2.中序遍历
func (t *BinarySearchTree) MidOrder() {
	MidOrder(t.root)
	fmt.Println()
}

func MidOrder(node *Node) {
	if node == nil {
		return
	}
	MidOrder(node.left)
	fmt.Printf("%d ", node.E)
	MidOrder(node.right)
}

// 3.后序遍历
func (t *BinarySearchTree) BackOrder() {
	BackOrder(t.root)
	println()
}

func BackOrder(node *Node) {
	if node == nil {
		return
	}
	BackOrder(node.left)
	BackOrder(node.right)
	fmt.Printf("%d ", node.E)
}

// 二分搜索树的层序遍历
func (t *BinarySearchTree) LevelOrder() {
	queue := make([]*Node, 0)
	queue = append(queue, t.root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", curNode.E)
		if curNode.left != nil {
			queue = append(queue, curNode.left)
		}
		if curNode.right != nil {
			queue = append(queue, curNode.right)
		}
	}
	println()
}

// 二分搜索树中搜索最小值
func (t *BinarySearchTree) FindMin() int {
	if t.IsEmpty() {
		panic("二叉树为空，无法删除任何节点")
	}
	return minimum(t.root).E
}
func minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return minimum(node.left)
}

// 二分搜索树中搜索最大值
func (t *BinarySearchTree) FindMax() int {
	if t.IsEmpty() {
		panic("二叉树为空，无法删除任何节点")
	}
	return maximum(t.root).E
}
func maximum(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return maximum(node.right)
}

// 从二分搜索树中删除最小值
func (t *BinarySearchTree) RemoveMin() int {
	var ret = t.FindMin()
	t.root = t.rmMin(t.root)
	return ret
}

// 删除掉以node为根的二分搜索树的最小节点
// 返回删除节点后新的二分搜索树的根
func (t *BinarySearchTree) rmMin(node *Node) *Node {
	if node.left == nil {
		nodeRight := node.right
		node.right = nil
		t.size--
		return nodeRight
	}
	node.left = t.rmMin(node.left)
	return node
}

// 从二分搜索树种删除最大值
func (t *BinarySearchTree) RemoveMax() int {
	var ret = t.FindMax()
	t.root = t.rmMax(t.root)
	return ret
}

// 删除掉以node为根的二分搜索树的最小节点
// 返回删除节点后新的二分搜索树的根
func (t *BinarySearchTree) rmMax(node *Node) *Node {
	if node.right == nil {
		nodeLeft := node.left
		node.left = nil
		t.size--
		return nodeLeft
	}
	node.right = t.rmMax(node.right)
	return node
}

// 在二分搜索树中删除值为e的方法
func (t *BinarySearchTree) Remove(e int) {
	t.root = t.remove(t.root, e)
}
func (t *BinarySearchTree) remove(node *Node, e int) *Node {
	if node == nil {
		return nil
	}
	if e > node.E {
		node.right = t.remove(node.right, e)
		return node
	} else if e < node.E {
		node.left = t.remove(node.left, e)
		return node
	} else {
		// 如果左子树为空的时候
		if node.left == nil {
			nodeRight := node.right
			node.right = nil
			t.size--
			return nodeRight
		}
		// 如果右子树为空
		if node.right == nil {
			nodeLeft := node.left
			node.left = nil
			t.size--
			return nodeLeft
		}
		// 如果左右子树都不为空，那么我们需要找到node的后继
		// 就是所有比node值大的节点中值最小的那个，显然它在node的右子树中
		nodeNext := minimum(node.right)
		nodeNext.right = t.rmMin(node.right)
		nodeNext.left = node.left
		node.left = nil
		node.right = nil
		return nodeNext
	}
}

// 二叉树的打印方法
func (t *BinarySearchTree) String() string {
	var (
		builder strings.Builder
	)
	generateBSTString(t.root, 0, &builder)
	return builder.String()
}

func generateBSTString(node *Node, depth int, builder *strings.Builder) {
	if node == nil {
		fmt.Fprintln(builder, generateDepthString(depth)+"null")
		return
	}
	fmt.Fprintln(builder, generateDepthString(depth), node.E)
	generateBSTString(node.left, depth+1, builder)
	generateBSTString(node.right, depth+1, builder)
}
func generateDepthString(depth int) string {
	var builder strings.Builder
	for i := 0; i < depth; i++ {
		fmt.Fprintf(&builder, "--")
	}
	return builder.String()
}
