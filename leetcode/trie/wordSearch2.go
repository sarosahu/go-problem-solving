/*
*

  - 212. Word Search II
    *

  - Given an m x n board of characters and a list of strings words,

  - return all words on the board.

  - Each word must be constructed from letters of sequentially adjacent cells,

  - where adjacent cells are horizontally or vertically neighboring.

  - The same letter cell may not be used more than once in a word.
    *
*/
package trie


func findWords(board [][]byte, words []string) []string {
    result := []string{}
	if len(board) == 0 {
		return result
	}
	m, n := len(board), len(board[0])
    // Build a trie and insert words.
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	root := trie.root
	// Now traverse the board
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := board[i][j]
			if _, exists := trie.root.children[c]; exists {
                res := []byte{}
				backtrack(&result, &board, i, j, root, &res)
			}
		}
	}
	return result
}

func backtrack(result *[]string, board *[][]byte, rowIdx, colIdx int, parentNode *Node, currentWord *[]byte) {
	ch := (*board)[rowIdx][colIdx]
	curr := parentNode.children[ch]

	// 1. Push current state & mark cell as visited
	*currentWord = append(*currentWord, ch)

	// 2. Collect match and immediately prevent duplicates
	if curr.isEnd {
		*result = append(*result, string(*currentWord))
		curr.isEnd = false 
	}

    // 3. Prune dynamically (Catches nodes emptied out by neighbor loops)
	if len(curr.children) == 0 {
		delete(parentNode.children, ch)
        *currentWord = (*currentWord)[:len(*currentWord)-1]
        return
	}

    (*board)[rowIdx][colIdx] = '#' // Move this up here so the cell is protected right away
	// 4. Explore Neighbors
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, dir := range dirs {
		nr, nc := rowIdx+dir[0], colIdx+dir[1]
        if !isValidCell(board, nr, nc) {
			continue
		}
        nextChar := (*board)[nr][nc]
		if _, exist := curr.children[nextChar]; !exist {
			continue
		}
        backtrack(result, board, nr, nc, curr, currentWord)
	}

	// 5. Backtrack cleanup (Always executes completely!)
	(*board)[rowIdx][colIdx] = ch  
	*currentWord = (*currentWord)[:len(*currentWord)-1]
}

func isValidCell(board *[][]byte, rowIdx int, colIdx int) bool {
	m, n := len(*board), len((*board)[0])
	if rowIdx < 0 || rowIdx >= m || colIdx < 0 || colIdx >= n {
		return false
	}
	return true
}