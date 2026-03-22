package arrays_hashing

// IsValidSudoku determines if a 9x9 Sudoku board is valid.
// Time Complexity: O(1) (since the board is always 9x9)
// Space Complexity: O(1)
func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			val := board[r][c]

			// Check row
			if rows[r][val] {
				return false
			}
			rows[r][val] = true

			// Check column
			if cols[c][val] {
				return false
			}
			cols[c][val] = true

			// Check 3x3 box
			boxIndex := (r/3)*3 + (c / 3)
			if boxes[boxIndex][val] {
				return false
			}
			boxes[boxIndex][val] = true
		}
	}

	return true
}
