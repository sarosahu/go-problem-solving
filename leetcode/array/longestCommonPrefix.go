/*
14. Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

 

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
 

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.

*/

package array

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {
    n := len(strs)
    if (n <= 1) {
        if n == 0 {
            return ""
        } else {
            return strs[0]
        }
    }

    prefix := strs[0]
    for i := 1; i < n; i++ {
        for strings.Index(strs[i], prefix) != 0 {
            prefix = prefix[:len(prefix) - 1]
            if (prefix == "") {
                return ""
            }
        }
    }
    return prefix
}