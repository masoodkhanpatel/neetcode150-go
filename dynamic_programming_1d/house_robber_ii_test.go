package dynamic_programming_1d

import "testing"

func TestRobII(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
		{[]int{1}, 1},
	}
	for _, tt := range tests {
		if got := RobII(tt.nums); got != tt.expected {
			t.Errorf("RobII(%v) = %d; want %d", tt.nums, got, tt.expected)
		}
	}
}
