package trie

import "testing"

func TestTrieA(t *testing.T) {
	trie := NewTrieA()

	// Test Case 1: Insert words
	wordsToInsert := []string{"apple", "app", "apricot", "banana"}
	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	// Test Case 2: Search exact words
	searchTests := []struct {
		word     string
		expected bool
	}{
		{"apple", true},
		{"app", true},
		{"apricot", true},
		{"banana", true},
		{"ap", false},       // Prefix only, not a full word
		{"apples", false},   // Extension of a word
		{"orange", false},   // Completely missing word
	}

	for _, tt := range searchTests {
		t.Run("Search_"+tt.word, func(t *testing.T) {
			actual := trie.Search(tt.word)
			if actual != tt.expected {
				t.Errorf("Search(%q) = %v; want %v", tt.word, actual, tt.expected)
			}
		})
	}

	// Test Case 3: StartsWith / Prefix search
	prefixTests := []struct {
		prefix   string
		expected bool
	}{
		{"ap", true},
		{"app", true},
		{"apple", true},
		{"ban", true},
		{"bana", true},
		{"cat", false},
		{"apples", false},
	}

	for _, tt := range prefixTests {
		t.Run("StartsWith_"+tt.prefix, func(t *testing.T) {
			actual := trie.StartsWith(tt.prefix)
			if actual != tt.expected {
				t.Errorf("StartsWith(%q) = %v; want %v", tt.prefix, actual, tt.expected)
			}
		})
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()

	// Test Case 1: Insert words
	wordsToInsert := []string{"apple", "app", "apricot", "banana"}
	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	// Test Case 2: Search exact words
	searchTests := []struct {
		word     string
		expected bool
	}{
		{"apple", true},
		{"app", true},
		{"apricot", true},
		{"banana", true},
		{"ap", false},       // Prefix only, not a full word
		{"apples", false},   // Extension of a word
		{"orange", false},   // Completely missing word
	}

	for _, tt := range searchTests {
		t.Run("Search_"+tt.word, func(t *testing.T) {
			actual := trie.Search(tt.word)
			if actual != tt.expected {
				t.Errorf("Search(%q) = %v; want %v", tt.word, actual, tt.expected)
			}
		})
	}

	// Test Case 3: StartsWith / Prefix search
	prefixTests := []struct {
		prefix   string
		expected bool
	}{
		{"ap", true},
		{"app", true},
		{"apple", true},
		{"ban", true},
		{"bana", true},
		{"cat", false},
		{"apples", false},
	}

	for _, tt := range prefixTests {
		t.Run("StartsWith_"+tt.prefix, func(t *testing.T) {
			actual := trie.StartsWith(tt.prefix)
			if actual != tt.expected {
				t.Errorf("StartsWith(%q) = %v; want %v", tt.prefix, actual, tt.expected)
			}
		})
	}
}