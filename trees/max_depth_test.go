package trees

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, 3},
		{[]int{1, -1, 2}, 2},
		{[]int{}, 0},
	}

	for _, test := range tests {
		root := SliceToTree(test.input)
		result := maxDepth(root)
		if result != test.expected {
			t.Errorf("expected %d, got %d", test.expected, result)
		}
	}
}
