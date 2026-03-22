package intervals

import (
	"container/heap"
	"sort"
)

// MinInterval returns for each query the size of the smallest interval containing it, or -1 if none.
// Time Complexity: O((n + q) log n)
// Space Complexity: O(n + q)
func MinInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Sort queries but keep original indices
	sortedQueries := make([][2]int, len(queries))
	for i, q := range queries {
		sortedQueries[i] = [2]int{q, i}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][0] < sortedQueries[j][0]
	})

	result := make([]int, len(queries))
	h := &intervalHeap{}
	heap.Init(h)
	i := 0

	for _, sq := range sortedQueries {
		q, idx := sq[0], sq[1]

		// Add all intervals whose start <= q
		for i < len(intervals) && intervals[i][0] <= q {
			size := intervals[i][1] - intervals[i][0] + 1
			heap.Push(h, [2]int{size, intervals[i][1]})
			i++
		}

		// Remove intervals that have ended before q
		for h.Len() > 0 && (*h)[0][1] < q {
			heap.Pop(h)
		}

		if h.Len() == 0 {
			result[idx] = -1
		} else {
			result[idx] = (*h)[0][0]
		}
	}

	return result
}

type intervalHeap [][2]int // [size, end]

func (h intervalHeap) Len() int           { return len(h) }
func (h intervalHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h intervalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intervalHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *intervalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
