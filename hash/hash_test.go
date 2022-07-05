package hash

import (
	"fmt"
	"testing"
)

// TODO: 哈希 冲突

func TestHashTable(t *testing.T) {
	dict := populateHashTable(3, 0)
	size := dict.Size()
	fmt.Println(fmt.Sprintf("Test failed, expected 3 and got %d", size))
	dict.String()
}
