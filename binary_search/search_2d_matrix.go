package binary_search

// Search a 2D Matrix
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1

	for l <= r {
		m := l + (r-l)/2
		val := matrix[m/cols][m%cols]
		if val == target {
			return true
		} else if val < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return false
}
