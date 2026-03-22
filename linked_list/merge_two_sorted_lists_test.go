package linked_list

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
	}

	for _, test := range tests {
		head1 := SliceToList(test.l1)
		head2 := SliceToList(test.l2)
		merged := mergeTwoLists(head1, head2)
		result := ListToSlice(merged)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
