package linked_list

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	}

	for _, test := range tests {
		head := SliceToList(test.input)
		reorderList(head)
		result := ListToSlice(head)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
