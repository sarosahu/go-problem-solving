/**
 * 267. Palindrome permutation II

 Given a string s, return all the palindromic permutations (without duplicates) of it.

You may return the answer in any order. If s has no palindromic permutation, return an empty list.

 

Example 1:

Input: s = "aabb"
Output: ["abba","baab"]
Example 2:

Input: s = "abc"
Output: []

Constraints:

1 <= s.length <= 16

 */
package backtrack

func generatePalindromes(s string) []string {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}
	mid := ""
	ans := []string{}
	for k, v := range cnt {
		if v % 2 == 1 {
			if mid != "" {
				return ans
			}
			mid = string(k)
		}
	}

	dfs(mid, s, &ans, cnt)
	return ans
}

func dfs(t string, s string, ans *[]string, cnt map[byte]int) {
	if len(t) == len(s) {
		*ans = append(*ans, t)
		return
	}
	for k, v := range cnt {
		if v > 1 {
			cnt[k] -= 2
			c := string(k)
			dfs(c + t + c, s, ans, cnt)
			cnt[k] += 2
		}
	}
}