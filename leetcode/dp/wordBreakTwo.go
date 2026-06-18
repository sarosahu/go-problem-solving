package dp

import "strings"

func WordBreakTwo(s string, wordDict []string) []string {
    return wordBreakTwoTrie(s, wordDict)
}

func wordBreakTwoTrie(s string, wordDict []string) []string {
	// Create a Trie
    trie := NewTrie()

	// Initialize the dp map --> it stores map of int to slices of string
	// where each string represents --> the valid words that must be present till the end of string 's'
	dp := make(map[int][]string)

	// Insert wordDict into trie
	for _, word := range wordDict {
		trie.Insert(word)
	}

	n := len(s)
	for start := n; start >= 0; start -- {
		var validSentenses []string
		curr := trie.root
		for end := start; end < n; end++ {
			c := rune(s[end])
			if _, exists := curr.children[c]; !exists {
				break;
			}
			curr = curr.children[c]
			if curr.isEnd {
				// Found a word in trie
				w := s[start: end+1]
				if end == n - 1 {
					validSentenses = append(validSentenses, w)
				} else {
					// Fetch all the valid sentenses from end + 1
					for _, sentense := range dp[end + 1] {
						var strBuilder strings.Builder
						strBuilder.WriteString(w)
						strBuilder.WriteString(" ")
						strBuilder.WriteString(sentense)
						validSentenses = append(validSentenses, strBuilder.String())
					}
				}
				
			}
		}
		dp[start] = validSentenses
	}

	return dp[0]
}