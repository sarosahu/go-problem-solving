package array

/**
* 79. Word Search

Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are
horizontally or vertically neighboring. The same letter cell may not be used more than once.
**/

func exist(board [][]byte, word string) bool {
    m := len(board)
    if m == 0 {
        return false
    }
    n := len(board[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] && backtrack(board, word, visited, i, j, 0) {
                return true
            }
        }
    }
    return false
}

func backtrack(board [][]byte, word string, visited [][]bool, i, j, idx int) bool {
    if idx == len(word) {
        return true
    }
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] || board[i][j] != word[idx] {
        return false
    }
    visited[i][j] = true

    // Define the 4 directions: down, left, up, right
	dirs := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	for _, dir := range dirs {
		if backtrack(board, word, visited, i+dir[0], j+dir[1], idx+1) {
			return true
		}
	}

	// Backtrack
	visited[i][j] = false
    return false
}