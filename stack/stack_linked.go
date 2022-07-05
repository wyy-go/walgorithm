package stack

import (
	"errors"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(value interface{})
	Pop() (interface{}, error)
	Length() int
}

func NewStackByLinked() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool { //判断是否为空
	return n.next == nil
}
func (n *Node) Push(value interface{}) {
	newnode := &Node{data: value} //初始化
	newnode.next = n.next
	n.next = newnode
}
func (n *Node) Pop() (interface{}, error) {
	if n.IsEmpty() == true {
		return nil, errors.New("bug")
	}
	value := n.next.data
	n.next = n.next.next
	return value, nil
}
func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.next != nil { //返回长度
		pnext = pnext.next
		length++
	}
	return length
}
