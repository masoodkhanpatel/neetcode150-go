package trees

import "container/list"

// BFS implementation for tree creation from slice (nulls as -1 for simplicity)
func SliceToTree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := list.New()
	queue.PushBack(root)
	i := 1
	for queue.Len() > 0 && i < len(nums) {
		node := queue.Remove(queue.Front()).(*TreeNode)

		if i < len(nums) && nums[i] != -1 {
			node.Left = &TreeNode{Val: nums[i]}
			queue.PushBack(node.Left)
		}
		i++

		if i < len(nums) && nums[i] != -1 {
			node.Right = &TreeNode{Val: nums[i]}
			queue.PushBack(node.Right)
		}
		i++
	}
	return root
}

func TreeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		element := queue.Front()
		node := queue.Remove(element)
		if node == nil {
			res = append(res, -1)
		} else {
			n := node.(*TreeNode)
			res = append(res, n.Val)
			if n.Left != nil {
				queue.PushBack(n.Left)
			} else {
				queue.PushBack(nil)
			}
			if n.Right != nil {
				queue.PushBack(n.Right)
			} else {
				queue.PushBack(nil)
			}
		}
	}
	// Trim trailing -1s
	for len(res) > 0 && res[len(res)-1] == -1 {
		res = res[:len(res)-1]
	}
	return res
}
