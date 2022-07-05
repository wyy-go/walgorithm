package hash

import "fmt"

// hash 表, key, v 组成
type Key interface{}

type Value interface{}

// 生成一个hash表
type HashTable struct {
	Items map[int]Value
}

// 给hash表中添加元素
func (ht *HashTable) Put(key Key, value Value) {
	index := customeHash(key)
	if ht.Items == nil {
		ht.Items = make(map[int]Value)
	}
	ht.Items[index] = value
}

func customeHash(k Key) int {
	// 自定义hash算法
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *HashTable) Get(key Key) Value {
	index := customeHash(key)
	return ht.Items[index]
}

func (ht *HashTable) Remove(key Key) {
	index := customeHash(key)
	delete(ht.Items, index)
}

func (ht *HashTable) Size() int {
	return len(ht.Items)
}

func (ht *HashTable) String() {
	// map遍历无序
	for k, v := range ht.Items {
		fmt.Println(fmt.Sprintf("k: %v\tv: %v", k, v))
	}
}

func populateHashTable(count int, start int) *HashTable {
	dict := HashTable{}
	for i := start; i < (start + count); i++ {
		dict.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	return &dict
}
