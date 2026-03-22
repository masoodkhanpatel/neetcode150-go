package dynamic_programming_2d

// UniquePaths returns the number of unique paths in an m×n grid (only right/down moves).
// Time Complexity: O(m*n)
// Space Complexity: O(n)
func UniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
