package linked_list

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		head := SliceToList(test.input)
		reversed := reverseList(head)
		result := ListToSlice(reversed)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
