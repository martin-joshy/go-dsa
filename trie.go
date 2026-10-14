package dsa

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{&TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, ch := range word {
		if _, exist := node.children[ch]; !exist {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for _, ch := range word {
		if _, exist := node.children[ch]; !exist {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}
