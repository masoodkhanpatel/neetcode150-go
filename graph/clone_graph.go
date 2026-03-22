package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph returns a deep copy of the graph.
// Time Complexity: O(n + e) where n is nodes and e is edges
// Space Complexity: O(n)
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldToNew := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if newNode, ok := oldToNew[node]; ok {
			return newNode
		}

		copyNode := &Node{Val: node.Val}
		oldToNew[node] = copyNode

		for _, neighbor := range node.Neighbors {
			copyNode.Neighbors = append(copyNode.Neighbors, dfs(neighbor))
		}

		return copyNode
	}

	return dfs(node)
}
