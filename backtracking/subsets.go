package backtracking

// Subsets returns all possible subsets (the power set).
// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)
func Subsets(nums []int) [][]int {
	var res [][]int
	var subset []int

	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		// Decision to include nums[i]
		subset = append(subset, nums[i])
		dfs(i + 1)

		// Decision NOT to include nums[i]
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return res
}
