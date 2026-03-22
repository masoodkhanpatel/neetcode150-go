package dynamic_programming_1d

// RobII returns the maximum money from circular houses (first and last are adjacent).
// Time Complexity: O(n)
// Space Complexity: O(1)
func RobII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a := robRange(nums, 0, len(nums)-2)
	b := robRange(nums, 1, len(nums)-1)
	if a > b {
		return a
	}
	return b
}

func robRange(nums []int, start, end int) int {
	prev2, prev1 := 0, 0
	for i := start; i <= end; i++ {
		cur := prev1
		if prev2+nums[i] > cur {
			cur = prev2 + nums[i]
		}
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}
