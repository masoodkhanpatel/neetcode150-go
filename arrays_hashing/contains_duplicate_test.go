package arrays_hashing

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Example 1: Duplicate exists",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "Example 2: No duplicates",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Example 3: Multiple duplicates",
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: false,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDuplicate(tt.input)
			if result != tt.expected {
				t.Errorf("ContainsDuplicate(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
