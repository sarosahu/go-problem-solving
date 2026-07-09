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

func TestMinRemoveToMakeValidUsing2Ptr(t *testing.T) {
	// Define the test cases using a table-driven approach
	tests := []struct {
		name     string
		input    string
		expected []string // Array of possible valid answers since multiple are allowed
	}{
		{
			name:     "Standard case with excess closing",
			input:    "lee(t(c)o)de)",
			expected: []string{"lee(t(c)o)de", "lee(t(co)de)", "lee(t(c)ode)"},
		},
		{
			name:     "Standard case with excess opening",
			input:    "a)b(c)d",
			expected: []string{"ab(c)d"},
		},
		{
			name:     "All invalid parentheses",
			input:    "))((",
			expected: []string{""},
		},
		{
			name:     "Already valid nested structure",
			input:    "(a(b(c)d)e)",
			expected: []string{"(a(b(c)d)e)"},
		},
		{
			name:     "No parentheses",
			input:    "abcde",
			expected: []string{"abcde"},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "Complex mixed characters",
			input:    "(()abc(d)",
			expected: []string{"()abc(d)", "(()abcd)"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minRemoveToMakeValidUsing2Ptr(tt.input)
			
			// Verify if the actual output matches any of the accepted valid variations
			matched := false
			for _, exp := range tt.expected {
				if actual == exp {
					matched = true
					break
				}
			}

			if !matched {
				t.Errorf("minRemoveToMakeValidUsing2Ptr(%q) = %q; want one of %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestMinRemoveToMakeValidUsingStack(t *testing.T) {
	// Define the test cases using a table-driven approach
	tests := []struct {
		name     string
		input    string
		expected []string // Array of possible valid answers since multiple are allowed
	}{
		{
			name:     "Standard case with excess closing",
			input:    "lee(t(c)o)de)",
			expected: []string{"lee(t(c)o)de", "lee(t(co)de)", "lee(t(c)ode)"},
		},
		{
			name:     "Standard case with excess opening",
			input:    "a)b(c)d",
			expected: []string{"ab(c)d"},
		},
		{
			name:     "All invalid parentheses",
			input:    "))((",
			expected: []string{""},
		},
		{
			name:     "Already valid nested structure",
			input:    "(a(b(c)d)e)",
			expected: []string{"(a(b(c)d)e)"},
		},
		{
			name:     "No parentheses",
			input:    "abcde",
			expected: []string{"abcde"},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "Complex mixed characters",
			input:    "(()abc(d)",
			expected: []string{"()abc(d)", "(()abcd)"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minRemoveToMakeValidUsingStack(tt.input)
			
			// Verify if the actual output matches any of the accepted valid variations
			matched := false
			for _, exp := range tt.expected {
				if actual == exp {
					matched = true
					break
				}
			}

			if !matched {
				t.Errorf("minRemoveToMakeValidUsing2Ptr(%q) = %q; want one of %v", tt.input, actual, tt.expected)
			}
		})
	}
}
