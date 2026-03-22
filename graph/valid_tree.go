package graph

// ValidTree returns true if the edges form a valid tree.
// Time Complexity: O(n + e)
// Space Complexity: O(n + e)
func ValidTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}

	adj := make(map[int][]int)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	visit := make(map[int]bool)

	var dfs func(int, int) bool
	dfs = func(i, prev int) bool {
		if visit[i] {
			return false
		}

		visit[i] = true
		for _, j := range adj[i] {
			if j == prev {
				continue
			}
			if !dfs(j, i) {
				return false
			}
		}
		return true
	}

	return dfs(0, -1) && len(visit) == n
}
