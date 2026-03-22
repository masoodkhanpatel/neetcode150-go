package graph

// NumIslands returns the number of islands in the grid.
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visit := make(map[[2]int]bool)
	islands := 0

	var bfs func(int, int)
	bfs = func(r, c int) {
		queue := [][2]int{{r, c}}
		visit[[2]int{r, c}] = true

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

			for _, d := range directions {
				nr, nc := curr[0]+d[0], curr[1]+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
					grid[nr][nc] == '1' && !visit[[2]int{nr, nc}] {
					queue = append(queue, [2]int{nr, nc})
					visit[[2]int{nr, nc}] = true
				}
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' && !visit[[2]int{r, c}] {
				bfs(r, c)
				islands++
			}
		}
	}

	return islands
}
