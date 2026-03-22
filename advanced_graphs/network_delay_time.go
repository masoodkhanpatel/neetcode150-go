package advanced_graphs

import "container/heap"

// NetworkDelayTime returns the time for signal to reach all n nodes from k, or -1 if unreachable.
// Uses Dijkstra's algorithm.
// Time Complexity: O((V + E) log V)
// Space Complexity: O(V + E)
func NetworkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][][2]int)
	for _, t := range times {
		adj[t[0]] = append(adj[t[0]], [2]int{t[2], t[1]})
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = 1<<31 - 1
	}
	dist[k] = 0

	h := &ndHeap{{0, k}}
	heap.Init(h)

	for h.Len() > 0 {
		item := heap.Pop(h).([2]int)
		d, u := item[0], item[1]
		if d > dist[u] {
			continue
		}
		for _, neighbor := range adj[u] {
			w, v := neighbor[0], neighbor[1]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(h, [2]int{dist[v], v})
			}
		}
	}

	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == 1<<31-1 {
			return -1
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	return maxDist
}

type ndHeap [][2]int

func (h ndHeap) Len() int           { return len(h) }
func (h ndHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h ndHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ndHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *ndHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
