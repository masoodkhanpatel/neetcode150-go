package dynamic_programming_2d

// NumDistinct returns the number of distinct subsequences of s that equal t.
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func NumDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
