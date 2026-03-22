package linked_list

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		inputs   [][]int
		expected []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{{}}, []int{}},
	}

	for _, test := range tests {
		var lists []*ListNode
		for _, input := range test.inputs {
			lists = append(lists, SliceToList(input))
		}
		resultHead := mergeKLists(lists)
		result := ListToSlice(resultHead)
		if !reflect.DeepEqual(result, test.expected) {
			if len(result) == 0 && len(test.expected) == 0 {
				continue
			}
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
