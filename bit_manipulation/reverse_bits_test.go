package bit_manipulation

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name     string
		n        uint32
		expected uint32
	}{
		{"Example 1", 0b00000010100101000001111010011100, 0b00111001011110000010100101000000},
		{"Example 2", 0b11111111111111111111111111111101, 0b10111111111111111111111111111111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBits(tt.n); got != tt.expected {
				t.Errorf("ReverseBits(%b) = %b; want %b", tt.n, got, tt.expected)
			}
		})
	}
}
