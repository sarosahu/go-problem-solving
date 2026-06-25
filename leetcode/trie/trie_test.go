package trie

import (
	"reflect"
	"sort"
	"testing"
)

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

func TestFindWords(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		words    []string
		want     []string
	}{
		{
			name: "Standard Grid Case",
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"eat", "oath"},
		},
		{
			name: "Prevent Duplicate Word Collection",
			// The word "a" can be formed from 4 different cells, 
			// but should only be reported once in the output.
			board: [][]byte{
				{'a', 'a'},
				{'a', 'a'},
			},
			words: []string{"a"},
			want:  []string{"a"},
		},
		{
			name: "No Words Found",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words: []string{"abcd", "xyz"},
			want:  []string{},
		},
		{
			name: "Overlapping Paths and Backtracking States",
			// Verifies that backing out of a dead end doesn't corrupt state
			board: [][]byte{
				{'a', 'b', 'c'},
				{'a', 'e', 'd'},
				{'a', 'f', 'g'},
			},
			words: []string{"abcdefg", "aba"},
			want:  []string{"abcdefg"},
		},
		{
			name:  "Empty Board Edge Case",
			board: [][]byte{},
			words: []string{"hello"},
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWords(tt.board, tt.words)

			// Sort slices to make testing order-independent
			sort.Strings(got)
			sort.Strings(tt.want)

			// Gracefully handle comparison of nil vs empty slices
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
