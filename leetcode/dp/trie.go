package dp

type Node struct {
	// A map matches each character to the next node
	// down the tree branch
	children map[rune]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{children: make(map[rune]*Node)},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root

	// Ranging over a string in Go yeilds 'rune' types automatically
	for _, char := range word {
		if _, exists := curr.children[char]; !exists {
			curr.children[char] = &Node{children: make(map[rune]*Node)}
		}
		curr = curr.children[char]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, char := range word {
		if _, exists := curr.children[char]; !exists {
			return false
		}
		curr = curr.children[char]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, char := range prefix {
		if _, exists := curr.children[char]; !exists {
			return false
		}
		curr = curr.children[char]
	}
	return true
}