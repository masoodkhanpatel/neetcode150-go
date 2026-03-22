package heap

import (
	"container/heap"
)

type Point struct {
	x, y int
	dist int
}

type PointHeap []Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].dist < h[j].dist } // Min-heap
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}
	for _, p := range points {
		dist := p[0]*p[0] + p[1]*p[1]
		heap.Push(h, Point{x: p[0], y: p[1], dist: dist})
	}

	res := make([][]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(Point)
		res[i] = []int{p.x, p.y}
	}
	return res
}
