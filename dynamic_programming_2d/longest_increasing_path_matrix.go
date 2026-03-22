package dynamic_programming_2d

// LongestIncreasingPath returns the length of the longest increasing path in a matrix.
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func LongestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		best := 1
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j] {
				if v := 1 + dfs(ni, nj); v > best {
					best = v
				}
			}
		}
		memo[i][j] = best
		return best
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if v := dfs(i, j); v > result {
				result = v
			}
		}
	}
	return result
}
