package trie

type Node struct {
	// A map matches each character to the next node
	// down the tree branch
	children map[byte]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{children: make(map[byte]*Node)},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root

	// Ranging over a string in Go yeilds 'rune' types automatically
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if _, exists := curr.children[ch]; !exists {
			curr.children[ch] = &Node{children: make(map[byte]*Node)}
		}
		curr = curr.children[ch]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if _, exists := curr.children[ch]; !exists {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for i := 0; i < len(prefix); i++  {
		ch := prefix[i]
		if _, exists := curr.children[ch]; !exists {
			return false
		}
		curr = curr.children[ch]
	}
	return true
}