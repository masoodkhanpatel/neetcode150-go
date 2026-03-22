package trees

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		want     []int
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{3, 9, 20, -1, -1, 15, 7}},
		{[]int{-1}, []int{-1}, []int{-1}},
	}
	for _, tt := range tests {
		root := buildTree(tt.preorder, tt.inorder)
		got := TreeToSlice(root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("buildTree() = %v, want %v", got, tt.want)
		}
	}
}
