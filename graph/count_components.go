package graph

// CountComponents returns the number of connected components in an undirected graph.
// Time Complexity: O(n + e * alpha(n))
// Space Complexity: O(n)
func CountComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(n1 int) int {
		res := n1
		for res != parent[res] {
			parent[res] = parent[parent[res]]
			res = parent[res]
		}
		return res
	}

	union := func(n1, n2 int) int {
		p1, p2 := find(n1), find(n2)
		if p1 == p2 {
			return 0
		}
		if rank[p1] > rank[p2] {
			parent[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parent[p1] = p2
			rank[p2] += rank[p1]
		}
		return 1
	}

	res := n
	for _, e := range edges {
		res -= union(e[0], e[1])
	}
	return res
}
