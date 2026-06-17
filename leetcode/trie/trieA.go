package trie

// Size defines the num of possible child nodes (a-z)
const Size = 26

// Node represents a single element in the Trie

type NodeA struct {
	children[Size] *NodeA
	isEnd bool
}

// Trie represents teh prefix tree structure
type TrieA struct {
	root *NodeA
}

// NewTrie initializes an empty Trie
func NewTrieA() *TrieA {
	return &TrieA{root: &NodeA{}}
}

// Insert adds a word to the TrieA
func (t *TrieA) Insert(word string) {
	curr := t.root
	for _, char := range word {
		idx := char - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &NodeA{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

// Search returns true if the exact word exists in the Trie
func (t *TrieA) Search(word string) bool {
	curr := t.root
	for _, char := range word {
		idx := char - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

// StartsWith returns true if there is any word in the TrieA that
// starts with the given prefix
func (t *TrieA) StartsWith(prefix string) bool {
	curr := t.root
	for _, char := range prefix {
		idx := char - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}