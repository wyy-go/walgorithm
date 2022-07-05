package linked

import (
	"fmt"
	"sync"
)

// Item the type of the linked list
type Item interface{}

// Node a single node that composes the list
type Node struct {
	data Item
	next *Node
}

// LinkedList the linked list of Items
type LinkedList struct {
	head *Node
	len  int
	lock sync.RWMutex
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds an Item to the end of the linked list
func (ll *LinkedList) Append(t Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	node := Node{t, nil}
	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.len++
}

// Insert adds an Item at position i
func (ll *LinkedList) Insert(i int, t Item) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.len {
		return fmt.Errorf("Index out of bounds")
	}
	addNode := Node{t, nil}

	node := ll.head
	if i == 0 {
		addNode.next = node
		ll.head = &addNode
	} else {
		for j := 0; j < i-1; j++ {
			node = node.next
		}
		addNode.next = node.next
		node.next = &addNode
	}

	ll.len++
	return nil
}

// RemoveAt removes a node at position i
func (ll *LinkedList) RemoveAt(i int) (*Item, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i >= ll.len {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := ll.head
	if i == 0 {
		ll.head = node.next
		ll.len--
		return &node.data, nil
	}

	j := 0
	for j < i {
		node = node.next
		j++
	}
	remove := node.next
	node.next = remove.next
	ll.len--
	return &remove.data, nil
}

// IndexOf returns the position of the Item t
func (ll *LinkedList) IndexOf(t Item) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	i := 0
	for {
		if node.data == t {
			return i
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		i++
	}
}

// IsEmpty returns true if the list is empty
func (ll *LinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if ll.head == nil {
		return true
	}
	return false
}

// Len returns the linked list len
func (ll *LinkedList) Len() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	last := ll.head
	if last == nil {
		return 0
	}

	size := 1
	for {
		if last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}

// Insert adds an Item at position i
func (ll *LinkedList) String() {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	for i := 0; node != nil; i++ {
		fmt.Print(node.data)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

// Head returns a pointer to the first node of the list
func (ll *LinkedList) Head() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head
}
