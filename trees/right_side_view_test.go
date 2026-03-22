package trees

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		root []int
		want []int
	}{
		{[]int{1, 2, 3, -1, 5, -1, 4}, []int{1, 3, 4}},
		{[]int{1, -1, 3}, []int{1, 3}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		root := SliceToTree(tt.root)
		if got := rightSideView(root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("rightSideView() = %v, want %v", got, tt.want)
		}
	}
}
