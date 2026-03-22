package heap

import (
	"container/heap"
)

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func LastStoneWeight(stones []int) int {
	h := &MaxIntHeap{}
	for _, s := range stones {
		heap.Push(h, s)
	}

	for h.Len() > 1 {
		s1 := heap.Pop(h).(int)
		s2 := heap.Pop(h).(int)
		if s1 != s2 {
			heap.Push(h, s1-s2)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}
