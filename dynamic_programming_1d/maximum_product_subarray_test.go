package dynamic_programming_1d

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2}, -2},
		{[]int{-2, 3, -4}, 24},
	}
	for _, tt := range tests {
		if got := MaxProduct(tt.nums); got != tt.expected {
			t.Errorf("MaxProduct(%v) = %d; want %d", tt.nums, got, tt.expected)
		}
	}
}
