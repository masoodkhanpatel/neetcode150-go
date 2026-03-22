package dynamic_programming_2d

import "testing"

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
		{[]int{1}, 1},
	}
	for _, tt := range tests {
		if got := MaxCoins(tt.nums); got != tt.expected {
			t.Errorf("MaxCoins(%v) = %d; want %d", tt.nums, got, tt.expected)
		}
	}
}
