package advanced_graphs

// FindCheapestPrice returns cheapest price from src to dst with at most k stops using Bellman-Ford.
// Time Complexity: O(k * E)
// Space Complexity: O(n)
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 1<<31 - 1
	prices := make([]int, n)
	for i := range prices {
		prices[i] = inf
	}
	prices[src] = 0

	for i := 0; i <= k; i++ {
		tmp := make([]int, n)
		copy(tmp, prices)
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]
			if prices[u] != inf && prices[u]+w < tmp[v] {
				tmp[v] = prices[u] + w
			}
		}
		prices = tmp
	}

	if prices[dst] == inf {
		return -1
	}
	return prices[dst]
}
