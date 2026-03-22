package linked_list

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	}

	for _, test := range tests {
		head := SliceToList(test.input)
		resultHead := removeNthFromEnd(head, test.n)
		result := ListToSlice(resultHead)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
