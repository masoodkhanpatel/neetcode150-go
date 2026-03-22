package trees

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		root []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-10, 9, 20, -1, -1, 15, 7}, 42},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := maxPathSum(root); got != tt.want {
			t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
		}
	}
}
