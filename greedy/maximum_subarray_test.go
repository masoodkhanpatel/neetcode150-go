package greedy

import "testing"

func TestMaxSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"Single positive", []int{1}, 1},
		{"All negative", []int{-1}, -1},
		{"Example 2", []int{5, 4, -1, 7, 8}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubarray(tt.nums); got != tt.expected {
				t.Errorf("MaxSubarray(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
