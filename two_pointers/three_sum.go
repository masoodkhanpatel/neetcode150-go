package two_pointers

import (
	"sort"
)

// ThreeSum returns all unique triplets in the array which gives the sum of zero.
// Time Complexity: O(n^2)
// Space Complexity: O(1) or O(n) depending on sorting implementation
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := n + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{n, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
