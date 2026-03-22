package advanced_graphs

import "sort"

// FindItinerary returns the itinerary starting from "JFK" using all tickets exactly once.
// Uses Hierholzer's algorithm for an Eulerian path.
// Time Complexity: O(E log E) for sorting
// Space Complexity: O(E)
func FindItinerary(tickets [][]string) []string {
	adj := make(map[string][]string)
	for _, t := range tickets {
		adj[t[0]] = append(adj[t[0]], t[1])
	}
	for k := range adj {
		sort.Strings(adj[k])
	}

	result := []string{}
	var dfs func(airport string)
	dfs = func(airport string) {
		for len(adj[airport]) > 0 {
			next := adj[airport][0]
			adj[airport] = adj[airport][1:]
			dfs(next)
		}
		result = append([]string{airport}, result...)
	}
	dfs("JFK")
	return result
}
