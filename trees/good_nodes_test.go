package trees

import (
	"testing"
)

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		root []int
		want int
	}{
		{[]int{3, 1, 4, 3, -1, 1, 5}, 4},
		{[]int{3, 3, -1, 4, 2}, 3},
		{[]int{1}, 1},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := goodNodes(root); got != tt.want {
			t.Errorf("goodNodes() = %v, want %v", got, tt.want)
		}
	}
}
