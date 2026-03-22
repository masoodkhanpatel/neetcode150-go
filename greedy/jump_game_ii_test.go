package greedy

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{2, 3, 1, 1, 4}, 2},
		{"Example 2", []int{2, 3, 0, 1, 4}, 2},
		{"Single element", []int{0}, 0},
		{"Two elements", []int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jump(tt.nums); got != tt.expected {
				t.Errorf("Jump(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
