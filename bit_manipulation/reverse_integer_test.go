package bit_manipulation

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{"Example 1", 123, 321},
		{"Negative", -123, -321},
		{"Trailing zero", 120, 21},
		{"Overflow", 1534236469, 0},
		{"Zero", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.x); got != tt.expected {
				t.Errorf("Reverse(%d) = %d; want %d", tt.x, got, tt.expected)
			}
		})
	}
}
