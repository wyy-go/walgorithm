package queue

//链表单节点
type QNode struct {
	data interface{} //数据
	next *QNode      //地址
}

type QueueLink struct { //队列的头部，尾部
	rear  *QNode
	front *QNode
}

type SQueue interface {
	length() int
	Enqueue(value interface{})
	Dequeue() (interface{}, error)
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) length() int {
	length := 0
	pnode := qlk.front //备份
	if pnode != nil {
		length = 1
	}
	for pnode != qlk.rear { //一直到循环到重合为止
		pnode = pnode.next //循环链表尾部
		length++
	}
	return length
}
func (qlk *QueueLink) Enqueue(value interface{}) {
	newnode := &QNode{data: value} //构造一个节点，返回地址
	if qlk.front == nil {          //只有一个节点
		qlk.front = newnode
		qlk.rear = newnode

	} else {
		qlk.rear.next = newnode
		qlk.rear = qlk.rear.next
	}
}

func (qlk *QueueLink) Dequeue() (value interface{}, err error) {
	if qlk.front == nil {
		return nil, nil
	}
	newnode := qlk.front
	if qlk.front == qlk.rear {
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.next //删除一个元素
	}
	return newnode.data, nil
}
