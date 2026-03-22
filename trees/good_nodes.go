package trees

func goodNodes(root *TreeNode) int {
	return dfsGoodNodes(root, -10001) // Assuming node values are >= -10000
}

func dfsGoodNodes(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val >= maxVal {
		count = 1
		maxVal = node.Val
	}
	count += dfsGoodNodes(node.Left, maxVal)
	count += dfsGoodNodes(node.Right, maxVal)
	return count
}
