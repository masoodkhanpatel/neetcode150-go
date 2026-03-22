package linked_list

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
	}

	for _, test := range tests {
		head := SliceToList(test.input)
		resultHead := reverseKGroup(head, test.k)
		result := ListToSlice(resultHead)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
