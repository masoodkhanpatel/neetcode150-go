package arrays_hashing

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Zeroes",
			nums:     []int{0, 0},
			expected: []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProductExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ProductExceptSelf(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
