package dynamic_programming_1d

import "sort"

// LengthOfLIS returns the length of the longest strictly increasing subsequence.
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func LengthOfLIS(nums []int) int {
	tails := []int{}
	for _, n := range nums {
		pos := sort.SearchInts(tails, n)
		if pos == len(tails) {
			tails = append(tails, n)
		} else {
			tails[pos] = n
		}
	}
	return len(tails)
}
