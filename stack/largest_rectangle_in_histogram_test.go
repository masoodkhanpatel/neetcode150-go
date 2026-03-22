package stack

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights  []int
		expected int
	}{
		{
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			heights:  []int{2, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		actual := largestRectangleArea(tt.heights)
		if actual != tt.expected {
			t.Errorf("largestRectangleArea(%v) = %d; expected %d", tt.heights, actual, tt.expected)
		}
	}
}
