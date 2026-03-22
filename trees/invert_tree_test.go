package trees

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		root := SliceToTree(test.input)
		inverted := invertTree(root)
		result := TreeToSlice(inverted)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
