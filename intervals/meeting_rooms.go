package intervals

import "sort"

// CanAttendMeetings returns true if a person can attend all meetings without overlap.
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func CanAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
