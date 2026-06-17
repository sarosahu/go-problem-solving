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
