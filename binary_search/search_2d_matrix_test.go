package binary_search

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
	}

	for _, tt := range tests {
		actual := searchMatrix(tt.matrix, tt.target)
		if actual != tt.expected {
			t.Errorf("searchMatrix(%v, %d) = %t; expected %t", tt.matrix, tt.target, actual, tt.expected)
		}
	}
}
