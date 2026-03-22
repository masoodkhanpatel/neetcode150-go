package trees

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		root []int
		k    int
		want int
	}{
		{[]int{3, 1, 4, -1, 2}, 1, 1},
		{[]int{5, 3, 6, 2, 4, -1, -1, 1}, 3, 3},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := kthSmallest(root, tt.k); got != tt.want {
			t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
		}
	}
}
