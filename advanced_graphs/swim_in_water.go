package advanced_graphs

import "container/heap"

// SwimInWater returns the minimum time to swim from (0,0) to (n-1,n-1).
// Uses modified Dijkstra: the cost to reach a cell is max elevation seen so far.
// Time Complexity: O(n^2 log n)
// Space Complexity: O(n^2)
func SwimInWater(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	h := &swimHeap{{grid[0][0], 0, 0}}
	heap.Init(h)
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for h.Len() > 0 {
		item := heap.Pop(h).([3]int)
		t, r, c := item[0], item[1], item[2]
		if r == n-1 && c == n-1 {
			return t
		}
		if visited[r][c] {
			continue
		}
		visited[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
				newT := t
				if grid[nr][nc] > newT {
					newT = grid[nr][nc]
				}
				heap.Push(h, [3]int{newT, nr, nc})
			}
		}
	}
	return -1
}

type swimHeap [][3]int

func (h swimHeap) Len() int           { return len(h) }
func (h swimHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h swimHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *swimHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}
func (h *swimHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
