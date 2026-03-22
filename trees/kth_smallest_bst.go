package trees

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}
