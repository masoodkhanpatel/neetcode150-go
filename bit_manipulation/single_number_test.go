package bit_manipulation

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{2, 2, 1}, 1},
		{"Example 2", []int{4, 1, 2, 1, 2}, 4},
		{"Single element", []int{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.nums); got != tt.expected {
				t.Errorf("SingleNumber(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
