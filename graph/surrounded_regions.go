package graph

// Solve captures all regions that are 4-directionally surrounded by 'X'.
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func Solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r == rows || c == cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'T'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// 1. Capture unsurrounded regions (O -> T)
	for r := 0; r < rows; r++ {
		dfs(r, 0)
		dfs(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		dfs(0, c)
		dfs(rows-1, c)
	}

	// 2. Capture surrounded regions (O -> X)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
		}
	}

	// 3. Uncapture unsurrounded regions (T -> O)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'T' {
				board[r][c] = 'O'
			}
		}
	}
}
