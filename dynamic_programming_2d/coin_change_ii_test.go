package dynamic_programming_2d

import "testing"

func TestChange(t *testing.T) {
	tests := []struct {
		amount   int
		coins    []int
		expected int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
	}
	for _, tt := range tests {
		if got := Change(tt.amount, tt.coins); got != tt.expected {
			t.Errorf("Change(%d, %v) = %d; want %d", tt.amount, tt.coins, got, tt.expected)
		}
	}
}
