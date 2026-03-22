package dynamic_programming_1d

// MaxProduct returns the maximum product subarray.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxProduct(nums []int) int {
	maxProd, minProd, result := nums[0], nums[0], nums[0]
	for _, n := range nums[1:] {
		candidates := [3]int{n, maxProd * n, minProd * n}
		maxProd = candidates[0]
		minProd = candidates[0]
		for _, c := range candidates[1:] {
			if c > maxProd {
				maxProd = c
			}
			if c < minProd {
				minProd = c
			}
		}
		if maxProd > result {
			result = maxProd
		}
	}
	return result
}
