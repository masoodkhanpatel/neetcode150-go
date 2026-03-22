package heap

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("FindKthLargest(%v, %d) = %d; want %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
