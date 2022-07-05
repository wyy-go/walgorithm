package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := New()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}

	assert.Equal(t, 10, list.Size())
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, list.dataStore)

	err := list.Remove(3)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{0, 1, 2, 4, 5, 6, 7, 8, 9}, list.dataStore)

	list.Insert(3, 100)
	assert.Equal(t, []interface{}{0, 1, 2, 100, 4, 5, 6, 7, 8, 9}, list.dataStore)

}
