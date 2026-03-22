package trees

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		root []int
		want [][]int
	}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{[]int{1}, [][]int{{1}}},
		{[]int{}, [][]int{}},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := levelOrder(root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("levelOrder() = %v, want %v", got, tt.want)
		}
	}
}
