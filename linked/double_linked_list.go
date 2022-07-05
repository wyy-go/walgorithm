package linked

import (
	"fmt"
	"sync"
)

//type LinkList interface {
//	Clear()
//	Len()
//	Front() *Element
//	Back() *Element
//	Remove(e *Element) *Element
//	SetData(old, new interface{})
//	Insert(e, mark *Element) *Element
//	InsertValue(v interface{}, mark *Element) *Element
//	PushFront(data interface{}) *Element
//	PushBack(data interface{}) *Element
//	InsertBefore(data interface{}, mark *Element) *Element
//	InsertAfter(data interface{}, mark *Element) *Element
//	MoveBefore(e, mark *Element)
//	MoveAfter(e, mark *Element)
//	Show()
//}

type Element struct {
	prev, next *Element
	value      interface{}
	list       *DoubleLinkedList
}

type DoubleLinkedList struct {
	root Element
	len  int
	rw   sync.RWMutex
}

func NewDoubleLinkedList() *DoubleLinkedList {
	list := new(DoubleLinkedList)
	list.Clear()
	return list
}

func (dl *DoubleLinkedList) Clear() {
	dl.rw.Lock()
	defer dl.rw.Unlock()
	if firstElem := dl.root.next; firstElem != nil && firstElem.list == dl {
		firstElem.prev = nil
	}
	if lastElem := dl.root.prev; lastElem != nil && lastElem.list == dl {
		lastElem.prev = nil
	}
	dl.root.next = &dl.root
	dl.root.prev = &dl.root
	dl.len = 0
}

func (dl *DoubleLinkedList) Len() int {
	dl.rw.RLock()
	defer dl.rw.RUnlock()
	return dl.len
}

func (dl *DoubleLinkedList) Front() *Element {
	dl.rw.RLock()
	defer dl.rw.RUnlock()
	if dl.len == 0 {
		return nil
	}
	return dl.root.next
}

func (dl *DoubleLinkedList) Back() *Element {
	dl.rw.RLock()
	defer dl.rw.RUnlock()
	if dl.len == 0 {
		return nil
	}
	return dl.root.prev
}

func (dl *DoubleLinkedList) Insert(e, mark *Element) *Element {
	dl.rw.Lock()
	defer dl.rw.Unlock()
	if mark == nil {
		return nil
	}
	next := mark.next
	mark.next = e
	e.prev = mark
	e.next = next
	next.prev = e
	e.list = dl
	dl.len++
	return e
}

func (dl *DoubleLinkedList) InsertValue(v interface{}, mark *Element) *Element {
	return dl.Insert(&Element{value: v}, mark)
}

func (dl *DoubleLinkedList) PushFront(data interface{}) *Element {
	return dl.InsertValue(data, &dl.root)
}

func (dl *DoubleLinkedList) PushBack(data interface{}) *Element {
	return dl.InsertValue(data, dl.root.prev)
}

func (dl *DoubleLinkedList) Remove(e *Element) *Element {
	dl.rw.Lock()
	defer dl.rw.Unlock()
	if e == nil {
		return nil
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	dl.len--
	return e
}

func (dl *DoubleLinkedList) RemoveData(data interface{}) interface{} {
	for pHead := dl.root.next; pHead != &dl.root; pHead = pHead.next {
		if pHead.value == data {
			dl.Remove(pHead) //查找并删除
			return pHead.value

		}

	}
	return nil
}

func (dl *DoubleLinkedList) Show() {
	dl.rw.RLock()
	defer dl.rw.RUnlock()
	for pHead := dl.root.next; pHead != &dl.root; pHead = pHead.next {
		fmt.Print(pHead.value)
		fmt.Print(" ")
	}
	fmt.Println()
}

func (dl *DoubleLinkedList) Values() []interface{} {
	dl.rw.RLock()
	defer dl.rw.RUnlock()
	var v []interface{}
	for pHead := dl.root.next; pHead != &dl.root; pHead = pHead.next {
		v = append(v, pHead.value)
	}
	return v
}

func (dl *DoubleLinkedList) InsertBefore(data interface{}, mark *Element) *Element {
	if mark.list != dl {
		return nil
	}
	return dl.InsertValue(data, mark.prev)
}

func (dl *DoubleLinkedList) InsertAfter(data interface{}, mark *Element) *Element {
	if mark.list != dl {
		return nil
	}
	return dl.InsertValue(data, mark)
}

func (dl *DoubleLinkedList) MoveBefore(e, mark *Element) {
	if e.list != dl || dl.root.next == e {
		return
	}

	dl.Insert(dl.Remove(e), mark.prev)
}

func (dl *DoubleLinkedList) MoveAfter(e, mark *Element) {
	if e.list != dl || dl.root.prev == e {
		return
	}
	dl.Insert(dl.Remove(e), &dl.root)
}

func (dl *DoubleLinkedList) SetData(old, new interface{}) {
	dl.rw.Lock()
	defer dl.rw.Unlock()
	for pHead := dl.root.next; pHead != &dl.root; pHead = pHead.next {
		if pHead.value == old {
			pHead.value = new
		}
	}
}
