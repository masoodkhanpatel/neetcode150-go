package graph

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "Example 1",
			grid:     [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expected: 4,
		},
		{
			name:     "Example 2",
			grid:     [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			expected: -1,
		},
		{
			name:     "Example 3",
			grid:     [][]int{{0, 2}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OrangesRotting(tt.grid)
			if result != tt.expected {
				t.Errorf("OrangesRotting() = %d; want %d", result, tt.expected)
			}
		})
	}
}
