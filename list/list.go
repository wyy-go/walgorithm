package list

import (
	"errors"
	"fmt"
)

var ListLength = 10 //定义列表全局长度

//interface  实例化int，int., string ,string
type List interface {
	Size() int                                  //函数大小，返回大小
	Get(index int) (interface{}, error)         //根据索引抓取数据
	Set(index int, newval interface{}) error    //设置
	Insert(index int, newval interface{}) error //插入
	Append(newval interface{}) error            //追加
	Remove(index int) error                     //删除
	Clear()                                     //清空
	String() string                             //返回字符串
}

// 结构体
type ArrayList struct {
	dataStore []interface{}
	len       int
}

// 创建新的链表
func New() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, ListLength) //分配内存10个数组元素

	list.len = 0 //0
	//fmt.Println("new",list.theSize,cap(list.dataStore))
	return list
}

// 获取链表长度
func (list *ArrayList) Size() int {
	return list.len //返回大小

}

// 追加数据
func (list *ArrayList) Append(newval interface{}) {

	list.dataStore = append(list.dataStore, newval) //数据叠加
	list.len++                                      //索引移动
}

// 获取索引所在数据
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.len {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

// 根据索引更改数据
func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.len {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval //赋值新的值
	return nil
}

// 检查切片是否已满
func (list *ArrayList) checkmemisfull() {
	if list.Size() == cap(list.dataStore) {
		newDataStore := make([]interface{}, 0, 2*list.Size()) //开辟更大内存
		copy(newDataStore, list.dataStore)                    //拷贝
		list.dataStore = newDataStore                         //赋值
	}
}

// 插入数据
func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.len {
		return errors.New("索引越界")
	}
	list.checkmemisfull()
	list.dataStore = list.dataStore[:list.Size()+1] //开辟内存,延展使用的内存
	for i := list.Size(); i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1] //从后往前赋值
	}
	list.dataStore[index] = newval //插入数据
	list.len++                     //索引加1

	return nil
}

// 根据索引移出数据
func (list *ArrayList) Remove(index int) error {
	if index < 0 || index >= list.len {
		return errors.New("索引越界")
	}

	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //删除
	list.len--
	return nil
}

// 清空数据
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10) //清空
	list.len = 0
}

// 返回字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
