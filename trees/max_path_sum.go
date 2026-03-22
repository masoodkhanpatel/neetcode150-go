package trees

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		
		leftMax := max(dfs(node.Left), 0)
		rightMax := max(dfs(node.Right), 0)
		
		res = max(res, node.Val + leftMax + rightMax)
		
		return node.Val + max(leftMax, rightMax)
	}
	
	dfs(root)
	return res
}
