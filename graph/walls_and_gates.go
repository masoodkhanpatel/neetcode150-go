package graph

// WallsAndGates fills each empty room with the distance to its nearest gate.
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func WallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}

	rows, cols := len(rooms), len(rooms[0])
	var queue [][2]int

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	dist := 0
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			r, c := curr[0], curr[1]

			directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && rooms[nr][nc] == 2147483647 {
					rooms[nr][nc] = rooms[r][c] + 1
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		dist++
	}
}
