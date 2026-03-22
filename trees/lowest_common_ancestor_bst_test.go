package trees

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root []int
		p    int
		q    int
		want int
	}{
		{[]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, 2, 8, 6},
		{[]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, 2, 4, 2},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		p := &TreeNode{Val: tt.p}
		q := &TreeNode{Val: tt.q}
		got := lowestCommonAncestor(root, p, q)
		if got == nil || got.Val != tt.want {
			t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
		}
	}
}
