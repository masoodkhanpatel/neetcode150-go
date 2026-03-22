package trees

import (
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2}, 1},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := diameterOfBinaryTree(root); got != tt.want {
			t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
		}
	}
}
