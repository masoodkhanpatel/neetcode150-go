package dynamic_programming_2d

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
		{[]int{1}, -1, 1},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1, 256},
	}
	for _, tt := range tests {
		if got := FindTargetSumWays(tt.nums, tt.target); got != tt.expected {
			t.Errorf("FindTargetSumWays(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.expected)
		}
	}
}
