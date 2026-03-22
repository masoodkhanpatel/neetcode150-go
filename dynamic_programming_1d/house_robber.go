package dynamic_programming_1d

// Rob returns the maximum amount of money robbing houses without robbing two adjacent.
// Time Complexity: O(n)
// Space Complexity: O(1)
func Rob(nums []int) int {
	prev2, prev1 := 0, 0
	for _, n := range nums {
		cur := prev1
		if prev2+n > cur {
			cur = prev2 + n
		}
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}
