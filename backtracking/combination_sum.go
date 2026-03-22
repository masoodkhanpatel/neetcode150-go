package backtracking

// CombinationSum returns a list of all unique combinations of candidates where the chosen numbers sum to target.
// Time Complexity: O(2^target)
// Space Complexity: O(target)
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var combination []int

	var dfs func(int, int)
	dfs = func(i int, currentSum int) {
		if currentSum == target {
			temp := make([]int, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}

		if i >= len(candidates) || currentSum > target {
			return
		}

		// Decision to include candidates[i]
		combination = append(combination, candidates[i])
		dfs(i, currentSum+candidates[i])

		// Decision NOT to include candidates[i]
		combination = combination[:len(combination)-1]
		dfs(i + 1, currentSum)
	}

	dfs(0, 0)
	return res
}
