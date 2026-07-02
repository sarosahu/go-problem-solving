package dp

/**
 * 10. Regular Expression Matching

Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.

Return a boolean indicating whether the matching covers the entire input string (not partial).

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".

Constraints:

1 <= s.length <= 20
1 <= p.length <= 20
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.

 */

func isMatch(s string, p string) bool {
    //return isMatchR(s, p)

    //return isMatchDPR(s, p)
    return isMatchDP(s, p)
}

func isMatchR(s string, p string) bool {
    lenP := len(p)
    lenS := len(s)
    if lenP == 0 {
        return lenS == 0
    }
    firstMatch := lenS != 0 && (s[0] == p[0] || p[0] == '.')
    if lenP >= 2 && p[1] == '*' {
        return isMatchR(s, p[2:]) ||
            firstMatch && isMatch(s[1:], p)
    } else {
        return firstMatch && isMatch(s[1:], p[1:])
    }
}

/* */
func isMatchDPR(s string, p string) bool {
    memo := make([][]*bool, len(s) + 1)
    for i := 0; i < len(memo); i++ {
        memo[i] = make([]*bool, len(p) + 1)
    }
    return isMatchDPRHelper(s, p, 0, 0, memo)
}

func isMatchDPRHelper(t string, p string, row int, col int, memo [][]*bool) bool {
    // If the state has already been calculated, return the cached result
    if memo[row][col] != nil {
        return *memo[row][col]
    }

    var ans bool
    tLen, pLen := len(t), len(p)
    if col == pLen {
        ans = row == tLen
    } else {
        // In Go, indexing a string returns a byte. This is completely safe
		// for standard regular expression patterns containing basic ASCII characters.
        firstMatch := row < tLen && (t[row] == p[col] || p[col] == '.')
        if col + 1 < pLen && p[col + 1] == '*' {
            // Zero match branch OR character repetition branch
            ans = isMatchDPRHelper(t, p, row, col + 2, memo) ||
                    (firstMatch && isMatchDPRHelper(t, p, row + 1, col, memo))
        } else {
            // Single character advance branch
            ans = firstMatch && isMatchDPRHelper(t, p, row + 1, col + 1, memo)
        }
    }
    // Cache the result before returning
    memo[row][col] = &ans
    return ans
}
/* */

func isMatchDP(s string, p string) bool {
    sLen, pLen := len(s), len(p)
    dp := make([][]bool, sLen + 1)
    for i := range dp {
        dp[i] = make([]bool, pLen + 1)
    }
    // Base case --> when both are empty string, then return true
    dp[sLen][pLen] = true

    for i := sLen; i >= 0; i-- {
        for j := pLen - 1; j >= 0; j-- {
            firstMatch := i < sLen && (s[i] == p[j] || p[j] == '.')
            if j + 1 < pLen && p[j + 1] == '*' {
                dp[i][j] = dp[i][j + 2] || (firstMatch && dp[i + 1][j])
            } else {
                dp[i][j] = firstMatch && dp[i+1][j+1]
            }
        }
    }
    return dp[0][0]
}