package dynamic_programming_2d

// MaxCoins returns the maximum coins obtainable by bursting all balloons.
// Time Complexity: O(n^3)
// Space Complexity: O(n^2)
func MaxCoins(nums []int) int {
	// Add boundary balloons
	balloons := make([]int, len(nums)+2)
	balloons[0] = 1
	balloons[len(balloons)-1] = 1
	copy(balloons[1:], nums)
	n := len(balloons)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length < n; length++ {
		for left := 0; left < n-length; left++ {
			right := left + length
			for k := left + 1; k < right; k++ {
				coins := balloons[left]*balloons[k]*balloons[right] + dp[left][k] + dp[k][right]
				if coins > dp[left][right] {
					dp[left][right] = coins
				}
			}
		}
	}
	return dp[0][n-1]
}
