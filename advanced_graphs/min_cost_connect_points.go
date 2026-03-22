package advanced_graphs

import "container/heap"

// MinCostConnectPoints returns the minimum cost to connect all points using Manhattan distance (Prim's).
// Time Complexity: O(n^2 log n)
// Space Complexity: O(n)
func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	inMST := make([]bool, n)
	h := &costHeap{{0, 0}} // {cost, index}
	heap.Init(h)
	total, visited := 0, 0

	for h.Len() > 0 && visited < n {
		item := heap.Pop(h).([2]int)
		cost, u := item[0], item[1]
		if inMST[u] {
			continue
		}
		inMST[u] = true
		total += cost
		visited++
		for v := 0; v < n; v++ {
			if !inMST[v] {
				d := abs(points[u][0]-points[v][0]) + abs(points[u][1]-points[v][1])
				heap.Push(h, [2]int{d, v})
			}
		}
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type costHeap [][2]int

func (h costHeap) Len() int           { return len(h) }
func (h costHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h costHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *costHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *costHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
