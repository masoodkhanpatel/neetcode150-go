package trees

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	
	mid := 0
	for i, val := range inorder {
		if val == rootVal {
			mid = i
			break
		}
	}
	
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	
	return root
}
