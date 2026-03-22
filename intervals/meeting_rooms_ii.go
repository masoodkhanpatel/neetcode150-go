package intervals

import (
	"container/heap"
	"sort"
)

// MinMeetingRooms returns the minimum number of conference rooms required.
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func MinMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &minHeap{}
	heap.Init(h)

	for _, interval := range intervals {
		if h.Len() > 0 && (*h)[0] <= interval[0] {
			heap.Pop(h)
		}
		heap.Push(h, interval[1])
	}
	return h.Len()
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
