package heap

import (
	"container/heap"
)

func FindKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
