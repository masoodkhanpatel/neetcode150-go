package bit_manipulation

// CountBits returns an array where ans[i] is the number of 1 bits in i, for 0 <= i <= n.
// Time Complexity: O(n)
// Space Complexity: O(n)
func CountBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}
