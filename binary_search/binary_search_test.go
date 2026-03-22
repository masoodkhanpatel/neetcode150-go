package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}

	for _, tt := range tests {
		actual := binarySearch(tt.nums, tt.target)
		if actual != tt.expected {
			t.Errorf("binarySearch(%v, %d) = %d; expected %d", tt.nums, tt.target, actual, tt.expected)
		}
	}
}
