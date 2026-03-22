package intervals

import "sort"

// EraseOverlapIntervals returns the minimum number of intervals to remove to make the rest non-overlapping.
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	end := intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] < end {
			count++
		} else {
			end = interval[1]
		}
	}
	return count
}
