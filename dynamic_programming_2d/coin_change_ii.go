package dynamic_programming_2d

// Change returns the number of combinations to make amount using coins (unbounded).
// Time Complexity: O(n * amount)
// Space Complexity: O(amount)
func Change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		for j := c; j <= amount; j++ {
			dp[j] += dp[j-c]
		}
	}
	return dp[amount]
}
