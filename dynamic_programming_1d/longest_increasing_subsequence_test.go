package dynamic_programming_1d

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}
	for _, tt := range tests {
		if got := LengthOfLIS(tt.nums); got != tt.expected {
			t.Errorf("LengthOfLIS(%v) = %d; want %d", tt.nums, got, tt.expected)
		}
	}
}
