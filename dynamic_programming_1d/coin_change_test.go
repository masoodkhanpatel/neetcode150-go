package dynamic_programming_1d

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 5, 10, 25}, 41, 4},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 11, 3},
	}
	for _, tt := range tests {
		if got := CoinChange(tt.coins, tt.amount); got != tt.expected {
			t.Errorf("CoinChange(%v, %d) = %d; want %d", tt.coins, tt.amount, got, tt.expected)
		}
	}
}
