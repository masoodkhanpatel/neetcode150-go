package bit_manipulation

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{3, 0, 1}, 2},
		{"Example 2", []int{0, 1}, 2},
		{"Example 3", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{"Single zero", []int{0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.nums); got != tt.expected {
				t.Errorf("MissingNumber(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
