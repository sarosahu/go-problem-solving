/*
  - 556. Next Greater Element III

    Given a positive integer n, find the smallest integer which has exactly the same digits existing in the integer n and is greater in value than n. If no such positive integer exists, return -1.

Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit integer, return -1.

Example 1:

Input: n = 12
Output: 21
Example 2:

Input: n = 21
Output: -1

Constraints:

1 <= n <= 2^31 - 1
*/
package stack

import (
	"math"
	"slices"
	"strconv"
)

func NextGreaterElement3(n int) int {
	str := strconv.Itoa(n)
	runes := []rune(str)
	
	// 1. Find the first decreasing element from the right
	i := len(runes) - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}
	
	// If no such element is found, digits are in descending order
	if i < 0 {
		return -1
	}
	
	// 2. Find the smallest element to the right of 'i' that is greater than runes[i]
	for j := len(runes) - 1; j > i; j-- { // optimized loop condition to 'j > i'
		if runes[j] > runes[i] {
			swap(runes, i, j)
			break
		}
	}
	
	// 3. FIX: Reverse the entire subarray to the right of 'i'
	slices.Reverse(runes[i+1:]) // Removed 'len(runes) - 1'

	// 4. Parse back to integer
	ans, err := strconv.Atoi(string(runes))
	if err != nil {
		return -1
	}
	
	// 5. Ensure the value fits in a standard 32-bit signed integer
	if ans > math.MaxInt32 {
		return -1
	}

	return ans
}

func swap(chars []rune, i, j int) {
	chars[i], chars[j] = chars[j], chars[i]
}