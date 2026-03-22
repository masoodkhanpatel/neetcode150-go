package graph

// FindRedundantConnection returns an edge that can be removed so that the resulting graph is a tree.
// Time Complexity: O(n * alpha(n)) where alpha is inverse Ackermann function
// Space Complexity: O(n)
func FindRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	rank := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(n int) int {
		p := parent[n]
		for p != parent[p] {
			parent[p] = parent[parent[p]]
			p = parent[p]
		}
		return p
	}

	union := func(n1, n2 int) bool {
		p1, p2 := find(n1), find(n2)
		if p1 == p2 {
			return false
		}

		if rank[p1] > rank[p2] {
			parent[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parent[p1] = p2
			rank[p2] += rank[p1]
		}
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return []int{}
}
