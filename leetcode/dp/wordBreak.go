/*
 * 139. Word Break
 * Given a string s and a dictionary of strings wordDict, 
 * return true if s can be segmented into a space-separated sequence of one or more dictionary words.

 * Note that the same word in the dictionary may be reused multiple times in the segmentation.

 Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
 

Constraints:

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.
 */

func wordBreak(s string, wordDict []string) bool {
    //return wordBreakEPI(s, wordDict)
    //return wordBreakBFS(s, wordDict)
    return wordBreakBottomUpDP(s, wordDict)
}

func wordBreakEPI(s string, wordDict []string) bool {
    n := len(s)
    // 1. Create a fast lookup set using an empty struct (0 bytes memory)
    wordSet := make(map[string]struct{}, len(wordDict))
    for _, val := range wordDict {
        wordSet[val] = struct{}{}
    }
    // table[i] represents if s[0:i] can be segmented
    table := make([]bool, n + 1)
    table[0] = true // Base case: empty prefix is valid

    // 2. Iterate through all possible prefix lengths
    for i := 1; i <= n; i++ {
        subs := s[0:i]
        if _, exist := wordSet[subs]; exist {
            table[i] = true
        }

        // Iterate only if table[i] is false
        for j := 0; !table[i] && j < i; j++ {
            if table[j] == false {
                continue
            }
            // If s[0:j] is valid AND s[j:i] is in the dictionary
            nextSubs := s[j:i]
            if _, exist := wordSet[nextSubs]; exist {
                table[i] = true
            }
        }
    }
    return table[n]
}
type Queue[T any] struct {
    items []T
}
func (q *Queue[T]) Size() int {
    return len(q.items)
}
func (q *Queue[T]) IsEmpty() bool {
    return q.Size() == 0
}
func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}
func (q *Queue[T]) Dequeue() (T, error) {
    var zero T
    if q.IsEmpty() {
        return zero, errors.New("Queue is empty")
    }
    item := q.items[0]
    q.items[0] = zero
    q.items = q.items[1:]

    return item, nil
}

func wordBreakBFS(s string, wordDict []string) bool {
    wordSet := make(map[string]struct{})
    for _, word := range wordDict {
        wordSet[word] = struct{}{}
    }
    n := len(s)
    seen := make([]bool, n + 1)
    queue := &Queue[int]{}
    queue.Enqueue(0)

    for !queue.IsEmpty() {
        start, _ := queue.Dequeue()
        if start == n {
            return true
        }
        for end := start + 1; end <= n; end++ {
            if seen[end] {
                continue
            }
            subStr := s[start:end]
            if _, exist := wordSet[subStr]; exist {
                queue.Enqueue(end)
                seen[end] = true
            }
        }
    }
    return false
}

func wordBreakBottomUpDP(s string, wordDict []string) bool {
    n := len(s)
    dp := make([]bool, n)

    
    for i := 0; i < n; i++ {
        for _, word := range wordDict {
            wn := len(word)
            if i < wn - 1 {
                continue
            }
            if i == wn - 1 || dp[i - wn] {
                curr := s[i - wn + 1:i + 1]
                if curr == word {
                    dp[i] = true
                    break
                }
            }
        }
    }
    return dp[n - 1]
}