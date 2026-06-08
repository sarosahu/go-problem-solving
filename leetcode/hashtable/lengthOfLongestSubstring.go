/*
 * 3. Longest Substring Without Repeating Characters

 Given a string s, find the length of the longest substring without duplicate characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 10^4
s consists of English letters, digits, symbols and spaces.


Hint 1
There are less than 100 unique characters. We can check all substrings with length at most 100 for example. This is a good enough approximation.
 */
package hashtable

func LengthOfLongestSubstring(s string) int {
    n := len(s)
    if n <= 1 {
        return n
    }
    maxLen := 1
    charToIdx := make(map[rune]int)
    startIdx := 0
    endIdx := 0
    for endIdx < n {
        c := s[endIdx]
        idx, exist := charToIdx[rune(c)]
        if exist {
            startIdx = max(idx + 1, startIdx)
        }
        charToIdx[rune(c)] = endIdx
        maxLen = max(maxLen, endIdx - startIdx + 1)
        endIdx++
    }
    return maxLen
}