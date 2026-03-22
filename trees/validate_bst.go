package trees

import "math"

func isValidBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(node *TreeNode, minVal, maxVal int64) bool {
	if node == nil {
		return true
	}
	if int64(node.Val) <= minVal || int64(node.Val) >= maxVal {
		return false
	}
	return validateBST(node.Left, minVal, int64(node.Val)) && 
		   validateBST(node.Right, int64(node.Val), maxVal)
}
