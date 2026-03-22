package dynamic_programming_2d

import "testing"

func TestMaxProfitCooldown(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
	}
	for _, tt := range tests {
		if got := MaxProfitCooldown(tt.prices); got != tt.expected {
			t.Errorf("MaxProfitCooldown(%v) = %d; want %d", tt.prices, got, tt.expected)
		}
	}
}
