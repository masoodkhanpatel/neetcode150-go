package trees

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root []int
		want bool
	}{
		{[]int{2, 1, 3}, true},
		{[]int{5, 1, 4, -1, -1, 3, 6}, false},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := isValidBST(root); got != tt.want {
			t.Errorf("isValidBST() = %v, want %v", got, tt.want)
		}
	}
}
