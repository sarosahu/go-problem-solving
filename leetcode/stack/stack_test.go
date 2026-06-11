package stack

import "testing"

func TestIsValid(t *testing.T) {
	// Define the test cases using a slice of structs
	tests := []struct {
		name     string // Description of the test case
		input    string // Input string to test
		expected bool   // Expected boolean output
	}{
		// --- True Cases ---
		{
			name:     "Simple matching parentheses",
			input:    "()",
			expected: true,
		},
		{
			name:     "Multiple consecutive matching pairs",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "Nested valid brackets",
			input:    "{[()]}",
			expected: true,
		},
		{
			name:     "Empty string (technically valid)",
			input:    "",
			expected: true,
		},

		// --- False Cases ---
		{
			name:     "Mismatched closing bracket",
			input:    "(]",
			expected: false,
		},
		{
			name:     "Incorrect closing order",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "Only opening brackets",
			input:    "(((",
			expected: false,
		},
		{
			name:     "Only closing brackets",
			input:    ")))",
			expected: false,
		},
		{
			name:     "Closer without an opener",
			input:    "([]))",
			expected: false,
		},
	}

	// Loop through all test cases
	for _, tc := range tests {
		// t.Run creates a subtest for clear terminal reporting
		t.Run(tc.name, func(t *testing.T) {
			actual := IsValidParentheses(tc.input)
			if actual != tc.expected {
				t.Errorf("isValid(%q) = %v; want %v", tc.input, actual, tc.expected)
			}
		})
	}
}
