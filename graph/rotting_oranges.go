package graph

// OrangesRotting returns the minimum number of minutes that must elapse until no cell has a fresh orange.
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func OrangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var queue [][2]int
	fresh := 0
	time := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				fresh++
			} else if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 && fresh > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			r, c := curr[0], curr[1]

			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					queue = append(queue, [2]int{nr, nc})
					fresh--
				}
			}
		}
		time++
	}

	if fresh == 0 {
		return time
	}
	return -1
}
