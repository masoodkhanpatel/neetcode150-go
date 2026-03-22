package backtracking

// Permute returns all possible permutations of nums.
// Time Complexity: O(n! * n)
// Space Complexity: O(n!)
func Permute(nums []int) [][]int {
	var res [][]int

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	for i := 0; i < len(nums); i++ {
		n := nums[0]
		nums = nums[1:]
		perms := Permute(nums)

		for _, p := range perms {
			p = append(p, n)
			temp := make([]int, len(p))
			copy(temp, p)
			res = append(res, temp)
		}

		nums = append(nums, n)
	}

	return res
}
