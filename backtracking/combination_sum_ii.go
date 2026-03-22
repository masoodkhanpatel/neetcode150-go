package backtracking

import "sort"

// CombinationSum2 returns a list of all unique combinations where the chosen numbers sum to target.
// Time Complexity: O(2^n)
// Space Complexity: O(n)
func CombinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var combination []int
	sort.Ints(candidates)

	var dfs func(int, int)
	dfs = func(i int, currentSum int) {
		if currentSum == target {
			temp := make([]int, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}

		if currentSum > target || i >= len(candidates) {
			return
		}

		// Include candidates[i]
		combination = append(combination, candidates[i])
		dfs(i+1, currentSum+candidates[i])
		combination = combination[:len(combination)-1]

		// Skip candidates[i] and its duplicates
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(i+1, currentSum)
	}

	dfs(0, 0)
	return res
}
