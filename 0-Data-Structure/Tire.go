package DataStructure

// TrieNode O(m)
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func InitTrie(items []string) *Trie {
	trie := NewTrie()

	for _, item := range items {
		trie.Insert(item)
	}

	return trie
}

func (t Trie) Insert(W string) {
	wLen := len(W)
	currentNode := t.root

	for i := 0; i < wLen; i++ {
		index := W[i] - 'a'
		if currentNode.children[index] == nil {
			currentNode.children[index] = NewTrieNode()
		}
		currentNode = currentNode.children[index]
	}
	currentNode.isEnd = true
}

func (t Trie) Search(W string) bool {
	wLen := len(W)
	currentNode := t.root

	for i := 0; i < wLen; i++ {
		index := W[i] - 'a'
		if currentNode.children[index] == nil {
			return false
		}
		currentNode = currentNode.children[index]
	}

	return currentNode.isEnd
}
