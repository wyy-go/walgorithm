package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {

	trie := NewTrie()
	trie.Add("a傲慢")
	trie.Add("傲慢")
	trie.Add2("ff")

	assert.Equal(t, 3, trie.GetSize()) // 单词数

	assert.True(t, trie.Contains("a傲慢"))
	assert.True(t, trie.Contains("傲慢"))
	assert.True(t, trie.Contains("ff"))
	assert.False(t, trie.Contains("f"))
	assert.False(t, trie.Contains("傲慢a"))

	assert.True(t, trie.IsPrefix("傲"))
	assert.True(t, trie.IsPrefix("a"))
	assert.True(t, trie.IsPrefix("f"))
	assert.False(t, trie.IsPrefix("fff"))
	assert.False(t, trie.IsPrefix("傲慢a"))

}
