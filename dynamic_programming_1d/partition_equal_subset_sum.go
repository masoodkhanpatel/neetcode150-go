package dynamic_programming_1d

// CanPartition returns true if the array can be partitioned into two equal-sum subsets.
// Time Complexity: O(n * sum)
// Space Complexity: O(sum)
func CanPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, n := range nums {
		for j := target; j >= n; j-- {
			dp[j] = dp[j] || dp[j-n]
		}
	}
	return dp[target]
}
