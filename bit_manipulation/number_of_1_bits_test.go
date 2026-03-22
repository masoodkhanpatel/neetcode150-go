package bit_manipulation

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name     string
		n        uint32
		expected int
	}{
		{"Example 1", 0b00000000000000000000000000001011, 3},
		{"Example 2", 0b00000000000000000000000010000000, 1},
		{"Example 3", 0b11111111111111111111111111111101, 31},
		{"Zero", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingWeight(tt.n); got != tt.expected {
				t.Errorf("HammingWeight(%b) = %d; want %d", tt.n, got, tt.expected)
			}
		})
	}
}
