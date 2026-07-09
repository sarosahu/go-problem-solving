package stack

import "strings"

/**
 * 1249. Minimum Remove to Make Valid Parentheses

 * Given a string s of '(' , ')' and lowercase English characters.

Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions )
so that the resulting parentheses string is valid and return any valid string.

Formally, a parentheses string is valid if and only if:

It is the empty string, contains only lowercase characters, or
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.

Example 1:

Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
Example 2:

Input: s = "a)b(c)d"
Output: "ab(c)d"
Example 3:

Input: s = "))(("
Output: ""
Explanation: An empty string is also valid.


Constraints:

1 <= s.length <= 10^5
s[i] is either '(' , ')', or lowercase English letter.
*/

func minRemoveToMakeValidUsingStack(s string) string {
	// Convert string to a slice of runes to handle modification
	arr := []byte(s)
	stack := Stack[int]{}

	for i, c := range arr {
		if c == '(' || c == ')' {
			if c == '(' {
				// Push the index to the stack
				stack.Push(i)
			} else {
				if !stack.IsEmpty() {
					// Pop the last element from the slice
					stack.Pop()
				} else {
					// Mark invalid closing parenthesis
					arr[i] = '*'
				}
			}
		}
	}

	// Mark all remaining unclosed opening parentheses
	for !stack.IsEmpty() {
		idx, _ := stack.Peek()
		arr[idx] = '*'
        stack.Pop()
	}

	// Build the final string efficiently using strings.Builder
	var result strings.Builder
	for _, c := range arr {
		if c != '*' {
			result.WriteByte(c)
		}
	}

	return result.String()
}

func minRemoveToMakeValidUsing2Ptr(s string) string {
	// Convert input string to byte slice for easy manipulation
	arr := []byte(s)
	
	// Counter for open parentheses
	openPCount := 0

	// First pass: mark excess closing parentheses with '*'
	for i := 0; i < len(arr); i++ {
		if arr[i] == '(' {
			openPCount++
		} else if arr[i] == ')' {
			if openPCount == 0 {
				arr[i] = '*' // Mark excess closing parentheses
			} else {
				openPCount--
			}
		}
	}

	// Second pass: mark excess opening parentheses from the end
	for i := len(arr) - 1; i >= 0; i-- {
		if openPCount > 0 && arr[i] == '(' {
			arr[i] = '*' // Mark excess opening parentheses
			openPCount--
		}
	}

	// Filter out marked characters in-place
	p := 0 // Pointer for updating the byte slice
	for i := 0; i < len(arr); i++ {
		if arr[i] != '*' {
			arr[p] = arr[i]
			p++
		}
	}

	// Construct the result string from the sliced byte array
	return string(arr[:p])
}

