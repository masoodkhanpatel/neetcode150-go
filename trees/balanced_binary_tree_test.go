package trees

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		root []int
		want bool
	}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, true},
		{[]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, false},
		{[]int{}, true},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := isBalanced(root); got != tt.want {
			t.Errorf("isBalanced() = %v, want %v", got, tt.want)
		}
	}
}
