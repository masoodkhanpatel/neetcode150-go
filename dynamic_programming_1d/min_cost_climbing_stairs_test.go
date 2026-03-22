package dynamic_programming_1d

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost     []int
		expected int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		if got := MinCostClimbingStairs(tt.cost); got != tt.expected {
			t.Errorf("MinCostClimbingStairs(%v) = %d; want %d", tt.cost, got, tt.expected)
		}
	}
}
