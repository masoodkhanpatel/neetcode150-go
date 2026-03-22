package dynamic_programming_1d

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{1}, 1},
		{[]int{2, 1}, 2},
	}
	for _, tt := range tests {
		if got := Rob(tt.nums); got != tt.expected {
			t.Errorf("Rob(%v) = %d; want %d", tt.nums, got, tt.expected)
		}
	}
}
