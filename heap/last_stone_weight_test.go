package heap

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name     string
		stones   []int
		expected int
	}{
		{
			name:     "Example 1",
			stones:   []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			name:     "Empty",
			stones:   []int{},
			expected: 0,
		},
		{
			name:     "Single stone",
			stones:   []int{1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastStoneWeight(tt.stones)
			if result != tt.expected {
				t.Errorf("LastStoneWeight(%v) = %d; want %d", tt.stones, result, tt.expected)
			}
		})
	}
}
