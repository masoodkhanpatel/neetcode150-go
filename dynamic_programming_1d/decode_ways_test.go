package dynamic_programming_1d

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"10", 1},
		{"2101", 1},
	}
	for _, tt := range tests {
		if got := NumDecodings(tt.s); got != tt.expected {
			t.Errorf("NumDecodings(%q) = %d; want %d", tt.s, got, tt.expected)
		}
	}
}
