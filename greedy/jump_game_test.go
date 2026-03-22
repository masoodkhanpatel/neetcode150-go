package greedy

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"Can reach", []int{2, 3, 1, 1, 4}, true},
		{"Cannot reach", []int{3, 2, 1, 0, 4}, false},
		{"Single element", []int{0}, true},
		{"Two elements", []int{1, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.nums); got != tt.expected {
				t.Errorf("CanJump(%v) = %v; want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
