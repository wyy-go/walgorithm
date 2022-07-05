package redblack

import (
	"fmt"
	"testing"
)

func TestRedBlack(t *testing.T) {

	rbTree := New()
	items := []*testStruct{
		{1, "this"},
		{2, "is"},
		{3, "a"},
		{4, "test"},
	}

	for i := range items {
		rbTree.Insert(items[i])
	}

	rbTree.Descend(items[2], func(i Item) bool {
		fmt.Println(i)
		return true
	})

	newItem := &testStruct{items[0].id, "not"}
	newItem = rbTree.InsertOrGet(newItem).(*testStruct)

	if newItem.text != items[0].text {
		fmt.Println(fmt.Errorf("tree.InsertOrGet = {id: %d, text: %s}, expect {id %d, text %s}", newItem.id, newItem.text, items[0].id, items[0].text))
	}

	newItem = &testStruct{5, "new"}
	newItem = rbTree.InsertOrGet(newItem).(*testStruct)

	if newItem.text != "new" {
		fmt.Println(fmt.Errorf("tree.InsertOrGet = {id: %d, text: %s}, expect {id %d, text %s}", newItem.id, newItem.text, 5, "new"))
	}

}
