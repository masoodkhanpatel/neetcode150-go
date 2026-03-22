package trees

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res = max(res, left+right)
		return 1 + max(left, right)
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
