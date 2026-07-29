/**
 * 127. Word Ladder

 * A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.

Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
Example 2:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
 

Constraints:

1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord, and wordList[i] consist of lowercase English letters.
beginWord != endWord
All the words in wordList are unique. 
*/
package bst

func ladderLength(beginWord string, endWord string, wordList []string) int {
    // { Initialize Set (Hash)
    setOfDict := make(map[string]struct{})
    for _, val := range wordList {
        setOfDict[val] = struct{}{}
    }
    // } Initialize Set (Hash)
    if _, exist := setOfDict[endWord]; !exist {
        return 0
    }
    n := len(beginWord)
    table := map[string][]string{}
    for _, word := range wordList {
        for i := 0; i < n; i++ {
            key := word[:i] + "*" + word[i+1:]
            if _, exist := table[key]; !exist {
                table[key] = []string{}
            }
            table[key] = append(table[key], word)
        }
    }
    queue := Queue[string]{}
    queue.Enqueue(beginWord)
    level := 1
    delete(setOfDict, beginWord)
    for !queue.IsEmpty() {
        sz := queue.Size()
        for sz > 0 {
            curr, _ := queue.Dequeue()
            if curr == endWord {
                return level
            }
            for i := 0; i < n; i++ {
                key := curr[:i] + "*" + curr[i+1:]
                for _, w := range table[key] {
                    if _, exist := setOfDict[w]; exist {
                        queue.Enqueue(w)
                        delete(setOfDict, w)
                    }
                }
            }
            sz--
        }
        level++
    }
    return 0
}