package greedy

// MaxSubarray returns the maximum subarray sum (Kadane's algorithm).
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSubarray(nums []int) int {
	maxSum := nums[0]
	cur := nums[0]
	for _, n := range nums[1:] {
		if cur < 0 {
			cur = 0
		}
		cur += n
		if cur > maxSum {
			maxSum = cur
		}
	}
	return maxSum
}
