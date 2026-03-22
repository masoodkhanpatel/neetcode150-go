package dynamic_programming_1d

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{1, 1}, true},
		{[]int{3, 3, 3, 4, 5}, true},
	}
	for _, tt := range tests {
		if got := CanPartition(tt.nums); got != tt.expected {
			t.Errorf("CanPartition(%v) = %v; want %v", tt.nums, got, tt.expected)
		}
	}
}
