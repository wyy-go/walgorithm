package stack

import (
	"errors"
)

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	isEmpty() bool
	isFull() bool
}

type ArrayList struct {
	dataStore []interface{}
	theSize   int
}

func New() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10) //分配内存10个数组元素

	list.theSize = 0 //0
	//fmt.Println("new",list.theSize,cap(list.dataStore))
	return list
}

func (list *ArrayList) Append(newval interface{}) {
	//list.checkmemisfull()
	list.dataStore = append(list.dataStore, newval) //数据叠加
	//fmt.Println(list.theSize,cap(list.dataStore))
	//list.dataStore[list.theSize]=newval  //赋值
	list.theSize++ //索引移动
}

func (list *ArrayList) Remove(index int) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}

	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //删除
	list.theSize--
	return nil
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10) //清空
	list.theSize = 0
}

type Stack struct {
	myarray *ArrayList
	capsize int
}

func NewStackByArray() *Stack {
	s := new(Stack)
	s.myarray = New()
	s.capsize = 10 //0
	return s
}

func (s *Stack) Clear() {
	s.myarray.Clear()
	s.myarray.theSize = 0
}

func (s *Stack) Size() int {
	return s.myarray.theSize
}

func (s *Stack) isEmpty() bool {
	if s.myarray.theSize == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) isFull() bool {
	if s.capsize == s.myarray.theSize {
		return true
	} else {
		return false
	}
}

func (s *Stack) Pop() interface{} {
	if !s.isEmpty() {
		last := s.myarray.dataStore[s.myarray.theSize-1]
		s.myarray.Remove(s.myarray.theSize - 1)
		return last
	}
	return nil
}

func (s *Stack) Push(data interface{}) {
	if !s.isFull() {
		s.myarray.Append(data)
	}
}
