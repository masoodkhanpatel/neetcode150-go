package backtracking

import "sort"

// SubsetsWithDup returns all possible subsets (the power set) containing duplicates.
// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)
func SubsetsWithDup(nums []int) [][]int {
	var res [][]int
	var subset []int
	sort.Ints(nums)

	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		// All subsets that include nums[i]
		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset)-1]

		// All subsets that don't include nums[i]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		dfs(i + 1)
	}

	dfs(0)
	return res
}
