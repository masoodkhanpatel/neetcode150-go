package trees

func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) (bool, int)
	dfs = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		leftBalanced, leftHeight := dfs(node.Left)
		rightBalanced, rightHeight := dfs(node.Right)
		
		balanced := leftBalanced && rightBalanced && abs(leftHeight-rightHeight) <= 1
		height := 1 + max(leftHeight, rightHeight)
		
		return balanced, height
	}
	balanced, _ := dfs(root)
	return balanced
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
