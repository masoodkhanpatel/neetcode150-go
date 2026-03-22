package advanced_graphs

// AlienOrder returns the order of characters in an alien language from sorted word list,
// or "" if no valid order exists (cycle detected).
// Time Complexity: O(C) where C is total length of all words
// Space Complexity: O(1) since at most 26 chars
func AlienOrder(words []string) string {
	// Build adjacency and in-degree
	adj := make(map[byte][]byte)
	inDegree := make(map[byte]int)
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if _, ok := inDegree[w[i]]; !ok {
				inDegree[w[i]] = 0
				adj[w[i]] = []byte{}
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2) < minLen {
			minLen = len(w2)
		}
		// If w1 is longer and is prefix of w2, invalid
		if len(w1) > len(w2) {
			found := true
			for k := 0; k < minLen; k++ {
				if w1[k] != w2[k] {
					found = false
					break
				}
			}
			if found {
				return ""
			}
		}
		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				adj[w1[j]] = append(adj[w1[j]], w2[j])
				inDegree[w2[j]]++
				break
			}
		}
	}

	// Topological sort (BFS / Kahn's)
	queue := []byte{}
	for c, d := range inDegree {
		if d == 0 {
			queue = append(queue, c)
		}
	}
	// Sort for determinism
	sortBytes(queue)

	result := []byte{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		result = append(result, c)
		neighbors := adj[c]
		sortBytes(neighbors)
		for _, nb := range neighbors {
			inDegree[nb]--
			if inDegree[nb] == 0 {
				queue = append(queue, nb)
				sortBytes(queue)
			}
		}
	}

	if len(result) != len(inDegree) {
		return ""
	}
	return string(result)
}

func sortBytes(b []byte) {
	for i := 1; i < len(b); i++ {
		for j := i; j > 0 && b[j] < b[j-1]; j-- {
			b[j], b[j-1] = b[j-1], b[j]
		}
	}
}
