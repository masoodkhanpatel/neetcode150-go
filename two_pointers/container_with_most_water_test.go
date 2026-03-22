package two_pointers

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxArea(tt.height)
			if result != tt.expected {
				t.Errorf("MaxArea(%v) = %d; want %d", tt.height, result, tt.expected)
			}
		})
	}
}
