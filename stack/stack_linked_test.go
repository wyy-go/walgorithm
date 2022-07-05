package stack

import (
	"fmt"
	"testing"
)

func TestStackByLinked(t *testing.T) {

	stack := NewStackByLinked()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	for data, err := stack.Pop(); err == nil; data, err = stack.Pop() {
		fmt.Println(data)
	}

}
