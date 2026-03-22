package heap

import (
	"container/heap"
)

type Task struct {
	count int
}

type MaxTaskHeap []int

func (h MaxTaskHeap) Len() int           { return len(h) }
func (h MaxTaskHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxTaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxTaskHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxTaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func LeastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int)
	for _, task := range tasks {
		counts[task]++
	}

	maxHeap := &MaxTaskHeap{}
	for _, count := range counts {
		heap.Push(maxHeap, count)
	}

	time := 0
	type item struct {
		count     int
		readyTime int
	}
	queue := []item{}

	for maxHeap.Len() > 0 || len(queue) > 0 {
		time++

		if maxHeap.Len() > 0 {
			cnt := heap.Pop(maxHeap).(int) - 1
			if cnt > 0 {
				queue = append(queue, item{cnt, time + n})
			}
		}

		if len(queue) > 0 && queue[0].readyTime == time {
			heap.Push(maxHeap, queue[0].count)
			queue = queue[1:]
		}
	}

	return time
}
