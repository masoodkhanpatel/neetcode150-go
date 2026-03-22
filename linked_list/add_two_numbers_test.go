package linked_list

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, test := range tests {
		head1 := SliceToList(test.l1)
		head2 := SliceToList(test.l2)
		resultHead := addTwoNumbers(head1, head2)
		result := ListToSlice(resultHead)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
