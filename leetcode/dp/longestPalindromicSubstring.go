package dp

func LongestPalindrome(s string, approachOne bool) string {
	if (approachOne) {
		return longestPalindromeDP(s)
	} else {
		return longestPalindromeDP2(s)
	}
}

func longestPalindromeDP(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }
    ans := [2]int{0, 0}
    palindromDP := make([][]bool, n)
    for i := range palindromDP {
        palindromDP[i] = make([]bool, n)
    }
    for end := 0; end < n; end++ {
        for start := 0; start <= end; start++ {
            if string(s[start]) == string(s[end]) &&
               (end - start <= 2 || palindromDP[start+1][end-1]) {
                palindromDP[start][end] = true
                if end - start > ans[1] - ans[0] {
                    ans[0], ans[1] = start, end
                }
            }
        }
    }
    substr := s[ans[0]:ans[1] + 1]

    return substr
}

type Boundary struct {
    start int
    end int
}

func longestPalindromeDP2(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }
    b := &Boundary{0, 0}
    for i := 0; i < n - 1; i++ {
        expandAroundCenter(s, i, i, b)
        expandAroundCenter(s, i, i + 1, b)
    }
    
    substr := s[b.start:b.end + 1]
    return substr
}

func expandAroundCenter(str string, left int, right int, b *Boundary) {
    n := len(str)
    for left >= 0 && right <= n - 1 && string(str[left]) == string(str[right]) {
        currLen := right - left + 1
        maxLen := b.end - b.start + 1
        if currLen > maxLen {
            b.start, b.end = left, right
        }
        left--
        right++
    }
}