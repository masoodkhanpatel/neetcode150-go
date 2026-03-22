package arrays_hashing

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4, // [1, 2, 3, 4]
		},
		{
			name:     "Example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9, // [0, 1, 2, 3, 4, 5, 6, 7, 8]
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Duplicates",
			nums:     []int{1, 2, 0, 1},
			expected: 3, // [0, 1, 2]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestConsecutive(tt.nums)
			if result != tt.expected {
				t.Errorf("LongestConsecutive(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
