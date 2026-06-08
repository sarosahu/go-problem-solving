/*
  - 36. Valid Sudoku

    Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/
package array

func IsValidSudoku(board [][]byte) bool {
	const N = 9

	// Use bitmasking to record previous occurrences
	rows := make([]int, N)
	cols := make([]int, N)
	boxes := make([]int, N)

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			// Check if the position is filled with a number
			currChar := board[r][c]
			if currChar == '.' {
				continue
			}
			val := int(currChar - '0')
			pos := 1 << (val - 1)

			// Check the row
			if (rows[r] & pos) > 0 {
				return false
			}
			rows[r] |= pos

			// Check the column
			if (cols[c] & pos) > 0 {
				return false
			}
			cols[c] |= pos

			// Check the box
			idx := (r / 3) * 3 + c / 3
			if (boxes[idx] & pos) > 0 {
				return false
			}
			boxes[idx] |= pos
		}
	}
	return true
}
