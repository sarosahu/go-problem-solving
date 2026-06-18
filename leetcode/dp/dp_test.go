package dp

import (
	"reflect"
	"sort"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	// Define our test cases using an anonymous struct slice
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Standard odd-length palindrome",
			input:    "babad",
			expected: "bab", // "aba" is also valid, but "bab" is found first
		},
		{
			name:     "Standard even-length palindrome",
			input:    "cbbd",
			expected: "bb",
		},
		{
			name:     "Single character string",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Entire string is a palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "The baab edge case (Even-length skip check)",
			input:    "bbaab",
			expected: "baab",
		},
		{
			name:     "No palindromes larger than 1 char",
			input:    "abcdef",
			expected: "a", // Returns the first character as the fallback
		},
	}

	// Loop through and run each test case independently
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := LongestPalindrome(tc.input, true)
			if actual != tc.expected {
				t.Errorf("longestPalindrome(%q) = %q; want %q", tc.input, actual, tc.expected)
			}
		})
	}

	// Loop through and run each test case independently
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := LongestPalindrome(tc.input, false)
			if actual != tc.expected {
				t.Errorf("longestPalindrome(%q) = %q; want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestGenerateParenthesis(t *testing.T) {
	// Define table-driven test cases
	tests := []struct {
		name     string
		input    int
		expected []string
	}{
		{
			name:     "n = 1",
			input:    1,
			expected: []string{"()"},
		},
		{
			name:     "n = 2",
			input:    2,
			expected: []string{"(())", "()()"},
		},
		{
			name:     "n = 3",
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get the actual result from your function
			actual := GenerateParenthesis(tt.input)

			// Sort both slices to ensure order doesn't cause a false failure
			sort.Strings(actual)
			sort.Strings(tt.expected)

			// Deeply compare the actual slice with the expected slice
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("generateParenthesis(%d) failed.\nGot:      %v\nExpected: %v", 
					tt.input, actual, tt.expected)
			}
		})
	}
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "Standard true case",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "Reuse dictionary words",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "Standard false case",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "Single character match",
			s:        "a",
			wordDict: []string{"a"},
			want:     true,
		},
		{
			name:     "Single character mismatch",
			s:        "b",
			wordDict: []string{"a"},
			want:     false,
		},
		{
			name:     "Overlapping words requiring backtracking",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			want:     true,
		},
		{
			name:     "Empty dictionary",
			s:        "abcdef",
			wordDict: []string{},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakEPI(tt.s, tt.wordDict)
			if got != tt.want {
				t.Errorf("wordBreak() = %v, want %v for string %q", got, tt.want, tt.s)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakBFS(tt.s, tt.wordDict)
			if got != tt.want {
				t.Errorf("wordBreak() = %v, want %v for string %q", got, tt.want, tt.s)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakBottomUpDP(tt.s, tt.wordDict)
			if got != tt.want {
				t.Errorf("wordBreak() = %v, want %v for string %q", got, tt.want, tt.s)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakTrie(tt.s, tt.wordDict)
			if got != tt.want {
				t.Errorf("wordBreak() = %v, want %v for string %q", got, tt.want, tt.s)
			}
		})
	}
}

func TestWordBreakTwo(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     []string
	}{
		{
			name:     "Standard multiple results",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			want:     []string{"cat sand dog", "cats and dog"},
		},
		{
			name:     "Pineapple combinations",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:     []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		},
		{
			name:     "No valid segmentation",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     []string{},
		},
		{
			name:     "Single character repetition",
			s:        "aaaaaaa",
			wordDict: []string{"a", "aa", "aaa"},
			// Will generate many valid combinations
			want: []string{
				"a a a a a a a", "a a a a a aa", "a a a a aa a", "a a a a aaa",
				"a a a aa a a", "a a a aa aa", "a a a aaa a", "a a aa a a a",
				"a a aa a aa", "a a aa aa a", "a a aa aaa", "a a aaa a a",
				"a a aaa aa", "a aa a a a a", "a aa a a aa", "a aa a aa a",
				"a aa a aaa", "a aa aa a a", "a aa aa aa", "a aa aaa a",
				"a aaa a a a", "a aaa a aa", "a aaa aa a", "a aaa aaa",
				"aa a a a a a", "aa a a a aa", "aa a a aa a", "aa a a aaa",
				"aa a aa a a", "aa a aa aa", "aa a aaa a", "aa aa a a a",
				"aa aa a aa", "aa aa aa a", "aa aa aaa", "aa aaa a a",
				"aa aaa aa", "aaa a a a a", "aaa a a aa", "aaa a aa a",
				"aaa a aaa", "aaa aa a a", "aaa aa aa", "aaa aaa a",
    		},
		},
		{
			name:     "Exact match single word",
			s:        "apple",
			wordDict: []string{"apple"},
			want:     []string{"apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakTwoTrie(tt.s, tt.wordDict)

			// Sort both slices to ensure order independence in validation
			sort.Strings(got)
			sort.Strings(tt.want)

			// Handle nil vs empty slice comparison gracefully
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordBreakTwoBottomUpDP(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     []string
	}{
		{
			name:     "Standard multiple results",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			want:     []string{"cat sand dog", "cats and dog"},
		},
		{
			name:     "Pineapple combinations",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:     []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		},
		{
			name:     "No valid segmentation",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     []string{},
		},
		{
			name:     "Single character repetition",
			s:        "aaaaaaa",
			wordDict: []string{"a", "aa", "aaa"},
			// Will generate many valid combinations
			want: []string{
				"a a a a a a a", "a a a a a aa", "a a a a aa a", "a a a a aaa",
				"a a a aa a a", "a a a aa aa", "a a a aaa a", "a a aa a a a",
				"a a aa a aa", "a a aa aa a", "a a aa aaa", "a a aaa a a",
				"a a aaa aa", "a aa a a a a", "a aa a a aa", "a aa a aa a",
				"a aa a aaa", "a aa aa a a", "a aa aa aa", "a aa aaa a",
				"a aaa a a a", "a aaa a aa", "a aaa aa a", "a aaa aaa",
				"aa a a a a a", "aa a a a aa", "aa a a aa a", "aa a a aaa",
				"aa a aa a a", "aa a aa aa", "aa a aaa a", "aa aa a a a",
				"aa aa a aa", "aa aa aa a", "aa aa aaa", "aa aaa a a",
				"aa aaa aa", "aaa a a a a", "aaa a a aa", "aaa a aa a",
				"aaa a aaa", "aaa aa a a", "aaa aa aa", "aaa aaa a",
    		},
		},
		{
			name:     "Exact match single word",
			s:        "apple",
			wordDict: []string{"apple"},
			want:     []string{"apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakTwoBottomUpDP(tt.s, tt.wordDict)

			// Sort both slices to ensure order independence in validation
			sort.Strings(got)
			sort.Strings(tt.want)

			// Handle nil vs empty slice comparison gracefully
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
