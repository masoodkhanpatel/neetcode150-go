package dynamic_programming_2d

// FindTargetSumWays returns the number of ways to assign +/- to reach target.
// Time Complexity: O(n * sum)
// Space Complexity: O(sum)
func FindTargetSumWays(nums []int, target int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	// dp with offset to handle negative indices
	offset := total
	size := 2*total + 1
	dp := make([]int, size)
	dp[offset] = 1
	for _, n := range nums {
		next := make([]int, size)
		for j := 0; j < size; j++ {
			if dp[j] > 0 {
				if j+n < size {
					next[j+n] += dp[j]
				}
				if j-n >= 0 {
					next[j-n] += dp[j]
				}
			}
		}
		dp = next
	}
	idx := target + offset
	if idx < 0 || idx >= size {
		return 0
	}
	return dp[idx]
}
