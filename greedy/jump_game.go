package greedy

// CanJump returns true if you can reach the last index.
// Time Complexity: O(n)
// Space Complexity: O(1)
func CanJump(nums []int) bool {
	maxReach := 0
	for i, v := range nums {
		if i > maxReach {
			return false
		}
		if i+v > maxReach {
			maxReach = i + v
		}
	}
	return true
}
