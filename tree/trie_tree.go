package tree

// 字典树(前缀树)
type trieNode struct {
	isWord bool // 是否是单词结尾
	next   map[interface{}]*trieNode
}

type trie struct {
	size int // 节点数量
	root *trieNode
}

func NewTrie() *trie {
	return &trie{
		0,
		&trieNode{false, make(map[interface{}]*trieNode)},
	}
}

func (this *trie) GetSize() int {
	return this.size
}

// 非递归算法
func (this *trie) Add(word string) {
	if len(word) == 0 {
		return
	}

	cur := this.root
	for _, v := range word {
		_, ok := cur.next[v] // 在NewTrie中已经初始化，能直接用
		if !ok {
			cur.next[v] = &trieNode{false, map[interface{}]*trieNode{}}
		}
		cur = cur.next[v]
	}
	// 判断该单词之前是否已经添加到tree中了
	if !cur.isWord {
		cur.isWord = true
		this.size++
	}
}

// 递归算法
func (this *trie) Add2(word string) {
	if len(word) == 0 {
		return
	}
	cur := this.root
	this.size = this.size + cur.insert(word)
}

// 辅助完成递归函数:在node节点中插入word,如果是已经存在的单词,返回0,如果不存在返回1
func (node *trieNode) insert(word string) int {
	_, ok := node.next[rune(word[0])]
	if !ok {
		node.next[rune(word[0])] = &trieNode{false,
			map[interface{}]*trieNode{}}
	}
	node = node.next[rune(word[0])]

	if len(word) == 1 {
		if !node.isWord {
			node.isWord = true
			return 1
		}
		return 0
	}

	return node.insert(word[1:])
}

// 查询是否包含某个单词
func (this *trie) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}

	cur := this.root
	for _, v := range word {
		t1, ok := cur.next[v]
		if !ok {
			return false
		}
		cur = t1
	}
	return cur.isWord
}

// 前缀是否有以prefix为前缀的单词
func (this *trie) IsPrefix(word string) bool {
	if len(word) == 0 {
		return false
	}

	cur := this.root
	for _, v := range word {
		t1, ok := cur.next[v]
		if !ok {
			return false
		}
		cur = t1
	}
	return true
}
