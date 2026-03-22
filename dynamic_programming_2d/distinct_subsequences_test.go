package dynamic_programming_2d

import "testing"

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		s, t     string
		expected int
	}{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
	}
	for _, tt := range tests {
		if got := NumDistinct(tt.s, tt.t); got != tt.expected {
			t.Errorf("NumDistinct(%q, %q) = %d; want %d", tt.s, tt.t, got, tt.expected)
		}
	}
}
