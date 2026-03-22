package trees

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		root    []int
		subRoot []int
		want    bool
	}{
		{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
		{[]int{3, 4, 5, 1, 2, -1, -1, -1, -1, 0}, []int{4, 1, 2}, false},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		subRoot := SliceToTree(tt.subRoot)
		if got := isSubtree(root, subRoot); got != tt.want {
			t.Errorf("isSubtree() = %v, want %v", got, tt.want)
		}
	}
}
