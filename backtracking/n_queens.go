package backtracking

import "strings"

// SolveNQueens returns all distinct solutions to the n-queens puzzle.
// Time Complexity: O(n!)
// Space Complexity: O(n^2)
func SolveNQueens(n int) [][]string {
	col := make(map[int]bool)
	posDiag := make(map[int]bool) // (r + c)
	negDiag := make(map[int]bool) // (r - c)

	var res [][]string
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	var backtrack func(int)
	backtrack = func(r int) {
		if r == n {
			var copyBoard []string
			for i := range board {
				copyBoard = append(copyBoard, strings.Join(board[i], ""))
			}
			res = append(res, copyBoard)
			return
		}

		for c := 0; c < n; c++ {
			if col[c] || posDiag[r+c] || negDiag[r-c] {
				continue
			}

			col[c] = true
			posDiag[r+c] = true
			negDiag[r-c] = true
			board[r][c] = "Q"

			backtrack(r + 1)

			delete(col, c)
			delete(posDiag, r+c)
			delete(negDiag, r-c)
			board[r][c] = "."
		}
	}

	backtrack(0)
	return res
}
